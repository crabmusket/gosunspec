package sunspec

import (
	"errors"
)

var (
	ErrNoSuchBlock error = errors.New("no such block")
	ErrNoSuchPoint       = errors.New("no such point")
)

type ModelId uint16

// These data types are defined in the SunSpec Information Model Specification.
// http://sunspec.org/wp-content/uploads/2015/06/SunSpec-Information-Models-12041.pdf
// In Version 1.8 of the document, the data type definitions can be found on
// page 13.

type Acc16 uint16
type Enum16 uint16
type Bitfield16 uint16
type Pad uint16

type Acc32 uint32
type Enum32 uint32
type Bitfield32 uint32

type Acc64 uint64

type Ipaddr [4]byte
type Ipv6addr [16]byte
type Eui48 [6]byte

type Float float32

type ScaleFactor int16

type Count uint16

// An array is a collection of devices.
//
// This type is used when the underlying representation
// is an XML document.
type Array interface {
	Do(func(d Device)) // Iterate over all the devices.
}

// A Device is a collection of Models.
type Device interface {
	// Model answers the model specified by id, or returns an error otherwise
	Model(id ModelId) (Model, error) // Answer the specified model instance or nil

	// MustModel answers the model specified by id, or panics otherwise
	MustModel(id ModelId) Model

	// Do iterates over all the models supported by the device
	Do(func(m Model))
}

// A Model is a collection of Blocks.
type Model interface {
	Id() ModelId // Answer the model identifier of the receiver
	Blocks() int // Answer the number of blocks

	Block(i int) (Block, error) // Answer the nth Block if it exists, or error otherwise.
	MustBlock(i int) Block      // Answer the block specified by i, or panics otherwise.
	Do(func(b Block))           // Iterate across all the blocks, in order.
}

// A Block is a collection of Points
//
// A Block maintains a buffer containing a copy of parts of the device address space,
//
// The Read operation populates the buffer from the physical device (using modbus or some other protocol)
// while a Write operation writes the buffer into the the physical device.
//
type Block interface {
	// Point answers the specified point if it exists or an error otherwise
	Point(id string) (Point, error)

	// MustPoint must answer the specified point or else it errors.
	MustPoint(id string) Point

	// Do iterates over all the points in the block in specification order.
	Do(func(p Point))

	// Read the specified pointIds or all if none is specified. A successful read implies
	// that subsequent access to the specified points will be error free.
	Read(pointIds ...string) error

	// Write the specified pointIds (or all, if none specified) into the physical device.
	Write(pointIds ...string) error
}

// A Point is collection of one or more 16 bit registers that is mapped
// onto a strongly typed golang type of the correct shape.
//
// The value associated with a Point must be accessed by the properly typed
// accessor or a panic will result. An attempt to obtain the value of a point
// which has a current error or which has not yet been read from the physical device
// will also result in a panic.
//
// Points also know how to apply scaling factors if an appropriate accessor
// method is used, so that if you use the ScaledValue() method with a 16 or 32 bit
// register for which a scaling factor is defined, then the register will be
// multiplied by the scaling factor to produce the expected 64 bit float result.
// Note that both the point and the related scaling factor point (if any) must
// be error free at the point the scaling factor is applied or a panic will result.
//
// The intent of the use of panics by this type is permit detection of unsafe
// access to data without requiring each and every call to be guarded with an
// error check which, in a correctly written program, will not actually be
// required.
//
type Point interface {
	Id() string   // the id of the point
	Type() string // the type, as defined by sunspec

	Error() error // any error preventing successful use of a getter or setter method.

	// getter methods

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

	// setter methods

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

	ScaledValue() float64

	Value() interface{}           // return a value of the correct type
	SetValue(v interface{}) error // set the value if it is the right type, return an error otherwise
}
