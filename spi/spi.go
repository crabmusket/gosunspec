// spi is a package used by physical implementations. it provides
// additional interfaces that allow the physical implementations
// to efficiently map the abstract application entities onto a physical
// substrate such as a Modbus link or an XML document.
package spi

import (
	"github.com/crabmusket/gosunspec"
)

// An anchor is a physical-implementation specific anchor to
// information that pertains to the physical implementation
// of a data element. So, for example, physical implementations
// the anchor might be the offset of a block or a model from the
// start of the physical address space. In an XML implementation
// the anchor might be a reference to correponding element
// of the XML representation.
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
	DoWithSPI(f func(p PointSPI))

	// Plan takes a set of pointIds to be read and returns a slice of points to
	// be read in the order they should be applied to the model
	// The algorithm ensures that:
	//    - if no points are specified, all are retrieved
	//    - if a point is read, then its scale factor (if any) is also read.
	//    - scale factors are applied to the model before any related points
	Plan(pointIds ...string) ([]PointSPI, error)

	// Invalidate any related point that depends on the specified point.
	Invalidate(p PointSPI)
}

// ModelSPI provides additional interfaces that the physical implementation
// needs to support the public interface of a model.
type ModelSPI interface {
	Anchored
	sunspec.Model
	Length() uint16
	DoWithSPI(f func(b BlockSPI))
	AddRepeat() error // Add one repeat to the model
}

// DeviceSPI provides additional interfaces that the physical implementation
// needs to support the public interface of a device.
type DeviceSPI interface {
	Anchored
	sunspec.Device
	DoWithSPI(f func(b ModelSPI))
	AddModel(m ModelSPI) error // Add a new model to the device
}

// ArraySPI provides additional interfaces that the physical implementation
// needs to support the public interface of an array.
type ArraySPI interface {
	Anchored
	sunspec.Array
	DoWithSPI(f func(b DeviceSPI))
	AddDevice(m DeviceSPI) error // Add a new model to the device
}

// Physical can read and write from the implementation model
// into the
type Physical interface {
	Write(block BlockSPI, pointIds ...string) error
	Read(block BlockSPI, pointIds ...string) error
}
