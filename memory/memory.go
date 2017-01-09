// memory is a package which allows sunspec devices to be simulated in memory
// so that such devices canbe accessed with the SunSpec API implemented
// by http://github.com/crabmusket/gosunspec
package memory

import (
	"encoding/binary"
	"errors"
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/impl"
	"github.com/crabmusket/gosunspec/layout"
	"github.com/crabmusket/gosunspec/models/model1"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/spi"
)

var (
	errNoModel        = errors.New("no model")
	errBadEyeCatcher  = errors.New("bad eyecatcher")
	errBufferTooShort = errors.New("buffer too short")
	eyeCatcher        = []byte{0x53, 0x75, 0x6e, 0x53} // "SunS"
)

type memoryDriver struct {
	buffer []byte
}

func (d *memoryDriver) iterator(b spi.BlockSPI, pointIds ...string) func(f func(buffer []byte, p spi.PointSPI) error) error {
	return func(f func(buffer []byte, p spi.PointSPI) error) error {
		var firstErr error

		points := []spi.PointSPI{}

		setError := func(e error) {
			if firstErr == nil && e != nil {
				firstErr = e
			}
		}

		for _, v := range pointIds {

			if p, err := b.Point(v); err != nil {
				setError(err)
			} else {
				points = append(points, p.(spi.PointSPI))
			}
		}

		if firstErr != nil {
			return firstErr
		}

		if len(points) == 0 {
			b.Do(spi.WithPointSPI(func(p spi.PointSPI) {
				points = append(points, p.(spi.PointSPI))
			}))
		}

		buffer := d.buffer[b.Anchor().(uint16)*2:]

		for _, p := range points {
			if p.Offset()+p.Length() > b.Length() {
				err := errBufferTooShort
				p.SetError(err)
				setError(err)
				continue
			}
			err := f(buffer[p.Offset()*2:(p.Offset()+p.Length())*2], p)
			if err != nil {
				p.SetError(err)
				setError(err)
			}
		}

		return firstErr
	}

}

func (d *memoryDriver) Read(b spi.BlockSPI, pointIds ...string) error {
	if points, err := b.Plan(pointIds...); err != nil {
		return err
	} else {
		var firstErr error
		buffer := d.buffer[b.Anchor().(uint16)*2:]
		for _, p := range points {
			if p.Offset()+p.Length() > b.Length() {
				err := errBufferTooShort
				p.SetError(err)
				if firstErr == nil {
					firstErr = err
				}
				continue
			}
			err := p.Unmarshal(buffer[p.Offset()*2 : (p.Offset()+p.Length())*2])
			if err != nil {
				p.SetError(err)
				if firstErr == nil {
					firstErr = err
				}
			}
		}
		return firstErr
	}
}

func (d *memoryDriver) Write(b spi.BlockSPI, pointIds ...string) error {
	return d.iterator(b, pointIds...)(func(buffer []byte, p spi.PointSPI) error {
		if p.Error() == nil {
			return p.Marshal(buffer)
		} else {
			return nil
		}
	})
}

func (d *memoryDriver) ReadWords(a uint16, l uint16) ([]byte, error) {
	return d.buffer[a*2 : (a+l)*2], nil
}

func (d *memoryDriver) BaseOffsets() []uint16 {
	return []uint16{0}
}

// Open a memory mapped Sunspec device from the specified
// byte slice or return an error if this cannot be done.
func Open(bytes []byte) (sunspec.Array, error) {
	return OpenWithLayout(bytes, &layout.SunSpecLayout{})
}

// Open a memory mapped Sunspec device from the specified
// byte slice, using the layout specified to match regions
// of the slice to particular devices and models.
func OpenWithLayout(bytes []byte, l layout.AddressSpaceLayout) (sunspec.Array, error) {
	return l.Read(&memoryDriver{buffer: bytes})
}

// SlabBuilder creates a slab of memory (actually a byte slice)
// to which a Sunspec device can be mapped.
//
// The main purpose of this type is to enable unit testing of the API
// without an actual Modbus device.
type SlabBuilder interface {
	AddDevice() SlabBuilder
	AddModel(id sunspec.ModelId) SlabBuilder
	AddRepeat(id sunspec.ModelId) SlabBuilder
	Build() ([]byte, error) // generates a byte slice containing the memory mapped device.
}

// Create a new device map builder.
func NewSlabBuilder() SlabBuilder {
	b := &builder{
		array: impl.NewArray(),
	}
	b.device = impl.NewDevice()
	b.AddModel(model1.ModelID) // all maps include the common model
	b.array.AddDevice(b.device)
	return b
}

type builder struct {
	array  spi.ArraySPI
	device spi.DeviceSPI
	err    error
}

func (b *builder) record(err error) {
	if err != nil && b.err == nil {
		b.err = err
	}
}

// Add the specified model to the device
func (b *builder) AddModel(id sunspec.ModelId) SlabBuilder {
	me := smdx.GetModel(uint16(id))
	if me != nil {
		m := impl.NewModel(me, 0, nil)
		b.record(b.device.AddModel(m))
	} else {
		b.record(errNoModel)
	}
	return b
}

func (b *builder) AddDevice() SlabBuilder {
	b.device = impl.NewDevice()
	b.AddModel(model1.ModelID) // all maps include the common model
	b.array.AddDevice(b.device)
	return b
}

// Add a repeat to the specified model
func (b *builder) AddRepeat(id sunspec.ModelId) SlabBuilder {
	// Note: this assumes all models of the same identifier
	// in the same address space have the same number of repeats
	// In principle, this might not be true. In practice,
	// it is likely to be true. We can live with the simplification
	// for now.
	b.device.Do(func(m sunspec.Model) {
		if m.Id() == id {
			b.record(m.(spi.ModelSPI).AddRepeat())
		}
	})
	return b
}

// Build a fresh byte slice in which the Sunspec device definition
// has been encoded. This byte slice can be accessed with the API
// by passing it to the Open function.
func (b *builder) Build() ([]byte, error) {

	if b.err != nil {
		return nil, b.err
	}

	// calculate the total size

	total := uint16(4) // eyecatcher + endmarker
	b.array.Do(spi.WithDeviceSPI(func(d spi.DeviceSPI) {
		d.Do(spi.WithModelSPI(func(m spi.ModelSPI) {
			total += 2 + m.Length() // header + model
		}))
	}))
	output := make([]byte, total*2)

	// render the eyecatcher and model/length markers

	copy(output, eyeCatcher)
	offset := 4
	b.array.Do(spi.WithDeviceSPI(func(d spi.DeviceSPI) {
		d.Do(spi.WithModelSPI(func(m spi.ModelSPI) {
			binary.BigEndian.PutUint16(output[offset:], uint16(m.Id()))
			binary.BigEndian.PutUint16(output[offset+2:], m.Length())
			offset += 4 + int(m.Length())*2
		}))
	}))

	// render the end marker

	binary.BigEndian.PutUint16(output[offset:], 0xffff)
	binary.BigEndian.PutUint16(output[offset+2:], 0)
	return output, nil
}
