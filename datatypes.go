package gosunspec

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

type String []byte

type Float float32

type ScaleFactor int16
