// spi is a package used by physical implementations or drivers. It provides
// additional interfaces to each implementation object that allow
// drivers to efficiently map the abstract API entities onto the underlying
// physical substrate such as a byte slice, a  Modbus link or an XML document.
package spi

import (
	"github.com/crabmusket/gosunspec"
)

// Driver is the interface implemented by components that can map
// the abstract API implementations onto a physical substrate. Each
// Driver implementation knows about a different kind of physical
// substrate, for example: byte slices, Modbus links or XML documents.
type Driver interface {
	Write(block BlockSPI, pointIds ...string) error
	Read(block BlockSPI, pointIds ...string) error
}

// An anchor is a driver specific anchor to information that pertains to the
// physical implementation of an API object. So, for example, in an address
// space based driver the anchor is the offset of a block or a model from
// the start of the physical address space. In an XML implementation the anchor
// is a reference to correponding element of the XML representation.
//
// The purpose it to allow navigation back to the physical
// implementation from the "canonical" representation.
type Anchor interface{}

type Anchored interface {
	Anchor() Anchor
	SetAnchor(a Anchor)
}

// PointSPI provides additional interfaces that the physical implementation
// needs to support the public interface of the point.
type PointSPI interface {
	Anchored
	sunspec.Point
	Length() uint16
	Offset() uint16
	Unmarshal([]byte) error
	Marshal([]byte) error
	SetError(err error) // SetError clears the value and sets the error to the specified error
	ScaleFactorPoint() PointSPI
	MarshalXML() string
	UnmarshalXML(s string) error
}

// BlockSPI provides additional interfaces that the physical implementation
// needs to support the public interface of a block.
type BlockSPI interface {
	Anchored
	sunspec.Block
	Length() uint16
	SetLength(l uint16) // in cases where actual length differs from the spec length

	// Plan takes a set of pointIds to be read and returns a slice of points to
	// be read in the order they should be applied to the model
	//
	// The algorithm ensures that:
	//    - if no points are specified, then all are read
	//    - if a point is read, then the related scale factor point (if any) is also read.
	//    - if a scale factor point is read, then any other point dependent on the scale factor
	//      is also read.
	//    - scale factors are applied to the model before any related points
	Plan(pointIds ...string) ([]PointSPI, error)
}

// ModelSPI provides additional interfaces that the physical implementation
// needs to support the public interface of a model.
type ModelSPI interface {
	Anchored
	sunspec.Model
	Length() uint16
	AddRepeat() error // Add one repeat to the model
}

// DeviceSPI provides additional interfaces that the physical implementation
// needs to support the public interface of a device.
type DeviceSPI interface {
	Anchored
	sunspec.Device
	AddModel(m ModelSPI) error // Add a new model to the device
}

// ArraySPI provides additional interfaces that the physical implementation
// needs to support the public interface of an array.
type ArraySPI interface {
	Anchored
	sunspec.Array
	AddDevice(m DeviceSPI) error // Add a new model to the device
}

// WithDeviceSPI answers a function that will apply the specified function, f, to the function's
// argument if, and only if, the argument is a Device which also implements DeviceSPI.
func WithDeviceSPI(f func(DeviceSPI)) func(sunspec.Device) {
	return func(d sunspec.Device) {
		if ds, ok := d.(DeviceSPI); ok {
			f(ds)
		}
	}
}

// WithModelSPI answers a function that will apply the specified function, f, to the function's
// argument if, and only if, the argument is a Model which also implements ModelSPI.
func WithModelSPI(f func(ModelSPI)) func(sunspec.Model) {
	return func(m sunspec.Model) {
		if ms, ok := m.(ModelSPI); ok {
			f(ms)
		}
	}
}

// WithBlockSPI answers a function that will apply the specified function, f, to the function's
// argument if, and only if, the argument is a Block which also implements BlockSPI.
func WithBlockSPI(f func(BlockSPI)) func(sunspec.Block) {
	return func(b sunspec.Block) {
		if bs, ok := b.(BlockSPI); ok {
			f(bs)
		}
	}
}

// WithPointSPI answers a function that will apply the specified function, f, to the function's
// argument if, and only if, the argument is a Point which also implements PointSPI.
func WithPointSPI(f func(PointSPI)) func(sunspec.Point) {
	return func(p sunspec.Point) {
		if ps, ok := p.(PointSPI); ok {
			f(ps)
		}
	}
}
