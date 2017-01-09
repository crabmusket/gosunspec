package layout

import (
	"errors"
	"github.com/crabmusket/gosunspec/spi"
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
