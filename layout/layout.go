package layout

import (
	"encoding/binary"
	"errors"
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/impl"
	"github.com/crabmusket/gosunspec/models/model1"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/spi"
	"log"
)

const (
	SunSpec = 0x53756e53 // "SunS" - marker bytes used to confirm that a region of Modbus address space is laid out according to SunSpec standards
)

var (
	ErrNotSunspecDevice = errors.New("not a SunSpec device") // if the Modbus address space doesn't contain the expected marker bytes
	ErrShortRead        = errors.New("short read")           // if an attempt to read from the Modbus addess space returns fewer bytes than expected
)

// AddressSpaceDriver abstracts the behaviour of drivers that are mapped onto a linear address space,
// so that the same AddressSpaceLayout implementation can be used across different linear
// address space implementations (e.g. memory & Modbus)
type AddressSpaceDriver interface {
	spi.Driver
	BaseOffsets() []uint16
	ReadWords(address uint16, length uint16) ([]byte, error)
}

// AddressSpaceLayout encapsulates an algorithm for scanning an address space for devices. There are
// two implementations - layout.SunSpecLayout and layout.RawLayout.
type AddressSpaceLayout interface {
	Read(a AddressSpaceDriver) (spi.ArraySPI, error)
}

// SunspecLayout is the type of layout that understands the SunSpec layout conventions.
type SunSpecLayout struct {
}

// RawModel layout specifies the model ID used to describe a block of memory at an Address
type RawModelLayout struct {
	Address uint16          `json:"address"`
	ModelId sunspec.ModelId `json:"modelId"`
}

// RawDeviceLayout is a slice of RawModelLayouts
type RawDeviceLayout struct {
	Models []RawModelLayout `json:"models"`
}

// RawLayout is the type of layout used for non-SunSpec address spaces. This means
// arbitrary Modbus address spaces where blocks are located at arbitrary addresses
// in an address space and neither the model ID nor the block length are encoded
// in the address space itself.
//
// The intent of RawLayout is to allow the sunspec API to be used effectively
// with non-SunSpec address spaces, assuming the work has been done to
// markup the address space with SMDX models and a JSON document that maps
// addresses to models.
type RawLayout struct {
	Devices []RawDeviceLayout `json:"devices"`
}

// Read scans the supplied address space and returns an array of the
// devices found.
func (s *SunSpecLayout) Read(a AddressSpaceDriver) (spi.ArraySPI, error) {

	// Attempt to locate SunSpec register within modbus address space.

	baseRange := a.BaseOffsets()
	base := uint16(0xffff)
	for _, b := range baseRange {
		if id, err := a.ReadWords(b, 2); err != nil {
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

	phys := a
	array := impl.NewArray()
	dev := impl.NewDevice()

	// Build up model

	offset := uint16(2) // number of 16 bit registers
	for {
		if bytes, err := a.ReadWords(base+offset, 2); err != nil {
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
