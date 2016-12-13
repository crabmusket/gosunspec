package modbus

import (
	"encoding/binary"
	"errors"
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/impl"
	"github.com/crabmusket/gosunspec/models/model1"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/spi"
	"github.com/goburrow/modbus"
	"log"
)

const (
	SunSpec = 0x53756e53 // "SunS" - marker bytes used to confirm that a region of Modbus address space is laid out according to SunSpec standards
)

var (
	ErrNotSunspecDevice = errors.New("not a SunSpec device") // if the Modbus address space doesn't contain the expected marker bytes
	ErrShortRead        = errors.New("short read")           // if an attempt to read from the Modbus addess space returns fewer bytes than expected
)

// Open uses the Modbus connection provided by client to connect
// to a Modbus address space. The address space is scanned for
// one or more SunSpec devices and a reference to
// a sunspec.Array that provides access to these devices is returned.
func Open(client modbus.Client) (sunspec.Array, error) {

	// Attempt to locate SunSpec register within modbus address space.

	baseRange := []uint16{40000, 50000, 0}
	base := uint16(0xffff)
	for _, b := range baseRange {
		if id, err := client.ReadHoldingRegisters(b, 2); err != nil {
			continue
		} else if binary.BigEndian.Uint32(id) != SunSpec {
			continue
		} else {
			base = b
			break
		}
	}
	if base == 0xffff {
		return nil, ErrNotSunspecDevice
	}

	phys := &modbusDriver{client: client}
	array := impl.NewArray()
	dev := impl.NewDevice()

	// Build up model

	offset := uint16(2) // number of 16 bit registers
	for {
		if bytes, err := client.ReadHoldingRegisters(base+offset, 2); err != nil {
			return nil, err
		} else if len(bytes) < 4 {
			return nil, ErrShortRead
		} else {
			modelId := binary.BigEndian.Uint16(bytes)
			modelLength := binary.BigEndian.Uint16(bytes[2:])

			if modelId == 0xffff {
				break
			}

			me := smdx.GetModel(modelId)
			if me != nil {

				if modelId == uint16(model1.ModelID) {
					dev = impl.NewDevice()
					array.(spi.ArraySPI).AddDevice(dev)
				}

				m := impl.NewContiguousModel(me, modelLength, phys)

				// set anchors on the blocks

				blockOffset := offset + 2
				m.Do(spi.WithBlockSPI(func(b spi.BlockSPI) {
					b.SetAnchor(uint16(base + blockOffset))
					blockOffset += b.Length()
				}))
				dev.AddModel(m)
			} else {
				log.Printf("unrecognised model identifier skipped @ offset: %d, %d\n", modelId, offset)
			}
			offset += 2 + modelLength
		}
	}
	return array, nil
}

type modbusDriver struct {
	client modbus.Client
}

// Write out the points in exactly the order specified, coalescing
// adjacent points if they are adjacent in the specified order.
func (p *modbusDriver) Write(block spi.BlockSPI, pointIds ...string) error {

	if len(pointIds) == 0 {
		block.Do(func(p sunspec.Point) {
			pointIds = append(pointIds, p.Id())
		})
	}

	// identify runs of adajacent points

	runs := newRunBuilder()

	// note: we preserve the programmer specified order
	// (not the specification order) because the write order
	// maybe significant in some cases especially if one
	// register is used to activate values previously
	// written into other registers.
	for _, pid := range pointIds {
		if p, err := block.Point(pid); err != nil {
			return err
		} else {
			runs.add(p.(spi.PointSPI))
		}
	}

	// marshal each group of adjacent points into byte slices and then
	// immediately write each byte slice into the modbus client.
	for _, run := range runs.runs {
		l := uint16(0)
		for _, pt := range run {
			l += pt.Length()
		}
		buffer := make([]byte, l*2, l*2)
		woff := uint16(0)
		for _, pt := range run {
			if err := pt.Marshal(buffer[woff : woff+pt.Length()*2]); err != nil {
				return err
			}
			woff += pt.Length() * 2
		}
		if _, err := p.client.WriteMultipleRegisters(block.Anchor().(uint16)+run[0].Offset(), l, buffer); err != nil {
			return err
		}
	}
	return nil
}

// Read extends the specified set of points with Block.Plan() then determines
// runs of points that can be read together. The points are read and then
// unmarshaled into the model in the order determined by slice returned by Block.Plan()
func (p *modbusDriver) Read(block spi.BlockSPI, pointIds ...string) error {
	if applicationOrder, err := block.Plan(pointIds...); err != nil {
		return err
	} else {
		runs := newRunBuilder()
		offsets := map[string]uint16{} // offsets into read buffer, by point
		off := uint16(0)               // the current offset
		toRead := map[string]bool{}    // the set of ponts to read

		// initialise the toRead set
		for _, p := range applicationOrder {
			toRead[p.Id()] = true
		}

		// break the list of points to be retrieved
		// into runs of strictly adjacent points and
		// record for each point the offset into a buffer
		// in which the marshaled point value will be read
		block.Do(spi.WithPointSPI(func(pt spi.PointSPI) {
			if !toRead[pt.Id()] {
				return
			}

			runs.add(pt)
			offsets[pt.Id()] = off
			off += pt.Length() * 2
		}))

		// allocate a buffer that can contain all the read points
		buffer := make([]byte, off, off)

		// read runs of points into the buffer
		off = 0
		for _, run := range runs.runs {
			l := uint16(0)
			for _, pt := range run {
				l += pt.Length()
			}
			if bytes, err := p.client.ReadHoldingRegisters(block.Anchor().(uint16)+run[0].Offset(), l); err != nil {
				return err
			} else {
				copy(buffer[off:off+l*2], bytes)
				off += (l * 2)
			}
		}

		// finally, unmarshal the buffer into points in the order determined by the plan
		for _, a := range applicationOrder {
			lbound := offsets[a.Id()]
			rbound := lbound + a.Length()*2
			if err := a.Unmarshal(buffer[lbound:rbound]); err != nil {
				return err
			}
		}
	}
	return nil
}
