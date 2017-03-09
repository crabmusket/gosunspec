// sunspec is an API that allows navigation around SunSpec device maps.
//
// The key abstractions are: Array, Device, Model, Block and Point with
// Arrays being collection of Devices, Devices collections of Models and so
// on. At the lowest level are Points which provide for strongly typed
// access to blocks of 16 bit registers.
//
// The API is intended to be backed by different physical implementations, including:
//
//  * Modbus RTU/TCP (for purposes of connecting to remote devices)
//  * XML (for purposes of data interchange)
//  * memory (for purposes of testing)
//
package sunspec

import (
	"errors"
)

var (
	ErrNoSuchBlock   = errors.New("no such block")
	ErrNoSuchModel   = errors.New("no such model")
	ErrTooManyModels = errors.New("too many models")
	ErrNoSuchPoint   = errors.New("no such point")
)

// ModelId is the type of model identifiers used with the Device.Model and Device.MustModel
// calls.
//
// To guarantee that a model's SMDX document is loaded into the SMDX registry you can use
// the value of modelXXXX.ModelID from the github.com/crabmusket/gosunspec/models/modelXXXX
// package to refer to the model, a side effect of which is to register the related SMDX model.
type ModelId uint16

// These data types are defined in the SunSpec Information Model Specification.
// http://sunspec.org/wp-content/uploads/2015/06/SunSpec-Information-Models-12041.pdf
// In Version 1.8 of the document, the data type definitions can be found on
// page 13.

// A 16 bit accumulator in the range 0-65535. The master must detect rollovers.
type Acc16 uint16

// A 16 bit accumulator in the range 0-4294967295. The master must detect rollovers.
type Acc32 uint32

// A 16 bit accumulator in the range 0-9223372036854775807. The master must detect rollovers.
type Acc64 uint64

// A 16 bit register with values from 0x0000-0x7fff. If the high bit 0x8000 is set, all
// other bits should be ignored.
type Bitfield16 uint16

// A 32 bit register with values from 0x0000-0x7fffffff. If the high bit 0x80000000 is set, all
// other bits should be ignored.
type Bitfield32 uint32

// A 16 bit register used to indicate the number of repeat blocks.
//
// NB: meaning inferred from use by some SMDX documents since SunSpec Information Model does not explicitly specify the semantics of this type.
type Count uint16

// A 16 bit register used to represent enumerated values. 65535 is reserved.
type Enum16 uint16

// A 32 bit register used to represent enumerated values. 4294967295 is reserved.
type Enum32 uint32

// An hardware address (like a MAC address) - see https://standards.ieee.org/develop/regauth/tut/eui48.pdf
type Eui48 [6]byte

// A 32bit IPv4 address (binary)
type Ipaddr [4]byte

// A 128 bit IPv6 address (binary)
type Ipv6addr [16]byte

// A 16-bit pad register. Always 0x8000.
type Pad uint16

// A 16-bit scaling factor used to scale the value of some other integer register.
type ScaleFactor int16

// An array is a collection of devices.
//
// This type is intended to be used when the underlying physical
// implementation is an XML document that may contain multiple
// 'd' elements but may also be used in cases where a Modbus address space
// has several device instances, in sequence, after the "Suns"
// marker bytes.
type Array interface {
	Do(func(d Device)) // Iterate over all the devices.
	Collect(func(d Device) bool) []Device
}

var AllDevices = func(d Device) bool { return true }

// A Device is a collection of Models that provides an uniform abstraction of
// physical devices of various kinds (Modbus, memory, XML documents)
type Device interface {
	// MustModel answers the first and only model specified by id, or panics otherwise
	// This method should not be used unless it is known for sure that the device
	// does not have multiple models with the specified id.
	MustModel(id ModelId) Model

	// Do iterates over all the models supported by the device
	Do(func(m Model))

	// Collects the subset of models that match the filter specified.
	//
	// for example:
	//     if m, err := ExactlyOneModel(d.Collect(SameModelId(model101.ModelID))); err == nil {
	//          // operate on the one and only model m
	//     }
	Collect(func(m Model) bool) []Model
}

var AllModels = func(m Model) bool { return true }

// A filter that can be used with Device.Collect to select a subset
// of the models that have the model id specified.
func SameModelId(id ModelId) func(m Model) bool {
	return func(m Model) bool {
		return id == m.Id()
	}
}

// Returns a function that matches a number of possible ModelIds
func OneOfSeveralModelIds(candidates []ModelId) func(m Model) bool {
	return func(m Model) bool {
		id := m.Id()
		for _, c := range candidates {
			if c == id {
				return true
			}
		}
		return false
	}
}

// ExactlyOneModel returns a model iff the slice contains
// exactly one model or an error otherwise.
func ExactlyOneModel(models []Model) (Model, error) {
	if len(models) > 1 {
		return nil, ErrTooManyModels
	} else if len(models) < 1 {
		return nil, ErrNoSuchModel
	} else {
		return models[0], nil
	}
}

// A Model is a collection of Blocks which represents a relocatable
// region of the physical device's address space and whose contents is specified
// by an SMDX document.
//
// By relocatable, we mean that the physical address of model is determined by
// device manufacturer, not some standard.
//
// Block 0 is the fixed block of the model, if any, otherwise it is the first repeating
// block. The remaining blocks, if any, are repeating blocks.
type Model interface {
	Id() ModelId                // Answer the model identifier of the receiver
	Blocks() int                // Answer the number of blocks
	Block(i int) (Block, error) // Answer the nth Block if it exists, or error otherwise.
	MustBlock(i int) Block      // Answer the block specified by i, or panics otherwise.
	Do(func(b Block))           // Iterate across all the blocks, in order.
}

// A Block is a contiguous collection of Points within a Model. The Points that
// comprise the Block are specified by a block element of a model's SMDX
// document.
//
// Subsets of a Block's Points can be read from or written to the underlying
// physical device using the Read and Write methods.
//
// A Block's Points exist in an error state prior to reading from underlying
// physical device. Upon a successful Read, the Points move from the error state
// to a populated state at which point the values can be accessed using the
// Point getter methods.
type Block interface {
	// Point answers the specified point if it exists or returns error otherwise.
	Point(id string) (Point, error)

	// MustPoint must answer the specified point or else it panics.
	MustPoint(id string) Point

	// Do iterates over all the points in the block in the specification order.
	Do(func(p Point))

	// Do iterates over all the scale factor points, then the non-scale factor
	// points. Both groups are processed in specification order. This method
	// is useful when multiple points in the block need to be updated and the
	// set of points to be updated includes both scale factor points and
	// value points related to the scale factor points.
	DoScaleFactorsFirst(func(p Point))

	// Read the registers implied by the specified pointIds from the physical
	// device. If no pointIds are specified, then all the points in the block
	// are read.
	//
	// If any specified or implied point has a related scaling factor, then the
	// related scaling factor point is also read.
	//
	// A successful read implies that subsequent access to the specified (and
	// implied) points will be error free.
	Read(pointIds ...string) error

	// Write the specified pointIds (or all, if none specified) into the
	// physical device. Like the Read case, if not pointIds are specified, then
	// all Blocks are written into the physical device. Unlike the Read case,
	// related scaling factors are not written unless explicitly specified.
	Write(pointIds ...string) error
}

// A Point is block of one or more 16 bit registers that can be mapped
// onto a value of a golang type of the correct shape.
//
// A Point is initially in an error state meaning its Error() method yields a
// non-nil value. While in the error state, any attempt to access the value of
// the Point will result in a panic. The error state is cleared by successfully
// reading the value of the Point from the physical device (using Block.Read) or
// by successfully setting its value with one of the Point setter methods.
//
// A program can assume that, immediately after a successful call to Block.Read(),
// then all Points that were either implicitly or explicitly specified by
// the arguments to Block.Read() were successfully read and hence will
// not be in an error state.
//
// If a program cannot prove that a Point is not in an error state, then
// the error state should be checked by calling Error() method prior
// to attempting to access the value of the Point.
//
// A Point can return to an error state if a related scaling factor is read
// (without reading the Point itself) or if the value of a related scaling
// factor is changed by calling one of its setter methods. One implication of
// this behaviour is that if you need to modify a value and its related scaling
// factor, you should always modify the related scaling factor first.
//
// Points also know how to apply scaling factors. If the ScaledValue() method is
// called for a Point with a 16, 32 or 64 bit numeric register type, then the
// register value will be multiplied by 10^scaling factor to produce a float64
// result containing scaled value of the Point's register value. Note that
// ScaledValue() will panic if either the Point or its related scaling factor
// Point are in an error state.
//
// A note about use of panics in the definition of this interface
//
// The intent of extensive use of panics in the definition of this interface is
// to permit detection of unsafe access to data without requiring each and every
// call to be guarded with an error check which, in a correctly written program,
// will not be required.
type Point interface {
	// The id of the point - see the models/modelXXX and the
	// related SMDX documents for a list of these.
	Id() string

	// The type, as defined by SunSpec. See the typelabel package for a list of these.
	Type() string

	// Answers the error associated with the point. This will be non nil, if:
	// - the parent block has not been read with a range of points that includes the receiver
	// - an attempt was made to read the parent block, but an error occurred during the read
	// - the related scaling factor point (if any) has been updated by a read or by a set
	Error() error

	// getter methods - calls will panic if Error() is not nil or if incorrectly typed method is called
	Acc16() Acc16
	Acc32() Acc32
	Acc64() Acc64
	Bitfield16() Bitfield16
	Bitfield32() Bitfield32
	Count() Count
	Enum16() Enum16
	Enum32() Enum32
	Eui48() Eui48
	Float32() float32
	Int16() int16
	Int32() int32
	Int64() int64
	Ipaddr() Ipaddr
	Ipv6addr() Ipv6addr
	Pad() Pad
	StringValue() string
	ScaleFactor() ScaleFactor
	Uint16() uint16
	Uint32() uint32
	Uint64() uint64

	// setter methods - calls will panic if incorrectly typed method is called.
	SetAcc16(v Acc16)
	SetAcc32(v Acc32)
	SetAcc64(v Acc64)
	SetBitfield16(v Bitfield16)
	SetBitfield32(v Bitfield32)
	SetCount(v Count)
	SetEnum16(v Enum16)
	SetEnum32(v Enum32)
	SetEui48(v Eui48)
	SetFloat32(v float32)
	SetInt16(v int16)
	SetInt32(v int32)
	SetInt64(v int64)
	SetIpaddr(v Ipaddr)
	SetIpv6addr(v Ipv6addr)
	SetPad(v Pad)
	SetStringValue(v string)
	SetScaleFactor(v ScaleFactor)
	SetUint16(v uint16)
	SetUint32(v uint32)
	SetUint64(v uint64)

	// Answer the scaled value of the point as a float64. This method
	// will panic if the Error() method of either the point or the
	// related scaling factor is not nil.
	ScaledValue() float64

	// Answer the value of the point. This method will panic if Error() is not nil.
	Value() interface{}

	// Set the value associated with the point. This method will panic if the value
	// type is not consistent with the type described by Type().
	SetValue(v interface{}) error
}
