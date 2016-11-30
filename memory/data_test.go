package memory

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/models/model101"
	"github.com/crabmusket/gosunspec/models/model11"
	"github.com/crabmusket/gosunspec/models/model304"
	"github.com/crabmusket/gosunspec/models/model401"
	"github.com/crabmusket/gosunspec/models/model501"
	"github.com/crabmusket/gosunspec/models/model63001"
	"github.com/crabmusket/gosunspec/typelabel"
)

var ComplexEmptySlab []byte
var ComplexNonZeroSlab []byte

const (
	CONST_ACC16  = sunspec.Acc16(106)
	CONST_ACC32  = sunspec.Acc32(100106)
	CONST_ACC64  = sunspec.Acc64(10000000106)
	CONST_BIT16  = sunspec.Bitfield16(0x0103)
	CONST_BIT32  = sunspec.Bitfield32(0x04050103)
	CONST_COUNT  = sunspec.Count(2)
	CONST_ENUM16 = sunspec.Enum16(104)
	CONST_ENUM32 = sunspec.Enum32(100104)
	CONST_PAD    = sunspec.Pad(0x8000)
	CONST_SF     = sunspec.ScaleFactor(-1)
	CONST_STRING = "A"
	CONST_UINT16 = uint16(102)
	CONST_UINT32 = uint32(100102)
)

var (
	CONST_INT16    = int16(-102)
	CONST_INT32    = int32(-100102)
	CONST_INT64    = int64(-10000000000000102)
	CONST_FLOAT32  = float32(100.02)
	CONST_IP       = sunspec.Ipaddr{127, 0, 0, 1}
	CONST_IPV6     = sunspec.Ipv6addr{0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00}
	CONST_EUI48    = sunspec.Eui48{0, 1, 2, 3, 4, 5}
	ExpectedValues = map[string]interface{}{
		typelabel.Acc16:       CONST_ACC16,
		typelabel.Acc32:       CONST_ACC32,
		typelabel.Acc64:       CONST_ACC64,
		typelabel.Bitfield16:  CONST_BIT16,
		typelabel.Bitfield32:  CONST_BIT32,
		typelabel.Count:       CONST_COUNT,
		typelabel.Enum16:      CONST_ENUM16,
		typelabel.Enum32:      CONST_ENUM32,
		typelabel.Float32:     CONST_FLOAT32,
		typelabel.Int16:       CONST_INT16,
		typelabel.Int32:       CONST_INT32,
		typelabel.Int64:       CONST_INT64,
		typelabel.Pad:         CONST_PAD,
		typelabel.ScaleFactor: CONST_SF,
		typelabel.String:      CONST_STRING,
		typelabel.Uint16:      CONST_UINT16,
		typelabel.Uint32:      CONST_UINT32,
		typelabel.Ipaddr:      CONST_IP,
		typelabel.Ipv6addr:    CONST_IPV6,
		typelabel.Eui48:       CONST_EUI48,
	}
)

func init() {
	bytes, _ := NewSlabBuilder().
		AddModel(model11.ModelID).
		AddModel(model101.ModelID).
		AddModel(model304.ModelID).
		AddRepeat(model304.ModelID).
		AddRepeat(model304.ModelID).
		AddModel(model401.ModelID).
		AddRepeat(model401.ModelID).
		AddRepeat(model401.ModelID).
		AddModel(model501.ModelID).
		AddModel(model63001.ModelID).
		AddRepeat(model63001.ModelID).
		Build()

	ComplexEmptySlab = bytes
	ComplexNonZeroSlab = make([]byte, len(bytes), len(bytes))
	copy(ComplexNonZeroSlab, ComplexEmptySlab)

	d, _ := Open(ComplexNonZeroSlab)
	d.Do(func(m sunspec.Model) {
		m.Do(func(b sunspec.Block) {
			// set scale factors first
			b.Do(func(p sunspec.Point) {
				switch p.Type() {
				case typelabel.ScaleFactor:
					p.SetScaleFactor(CONST_SF)
				default:
				}
			})

			// then the other types ...
			b.Do(func(p sunspec.Point) {
				switch p.Type() {
				case typelabel.Count:
					p.SetCount(CONST_COUNT)
				case typelabel.Uint32:
					p.SetUint32(CONST_UINT32)
				case typelabel.Float32:
					p.SetFloat32(CONST_FLOAT32)
				case typelabel.Uint16:
					p.SetUint16(CONST_UINT16)
				case typelabel.Acc16:
					p.SetAcc16(CONST_ACC16)
				case typelabel.Acc32:
					p.SetAcc32(CONST_ACC32)
				case typelabel.Acc64:
					p.SetAcc64(CONST_ACC64)
				case typelabel.Int16:
					p.SetInt16(CONST_INT16)
				case typelabel.Int32:
					p.SetInt32(CONST_INT32)
				case typelabel.Int64:
					p.SetInt64(CONST_INT64)
				case typelabel.Bitfield16:
					p.SetBitfield16(CONST_BIT16)
				case typelabel.Bitfield32:
					p.SetBitfield32(CONST_BIT32)
				case typelabel.Enum32:
					p.SetEnum32(CONST_ENUM32)
				case typelabel.Enum16:
					p.SetEnum16(CONST_ENUM16)
				case typelabel.Pad:
					p.SetPad(CONST_PAD)
				case typelabel.String:
					p.SetStringValue(CONST_STRING)
				case typelabel.Ipaddr:
					p.SetIpaddr(CONST_IP)
				case typelabel.Ipv6addr:
					p.SetIpv6addr(CONST_IPV6)
				case typelabel.Eui48:
					p.SetEui48(CONST_EUI48)
				case typelabel.ScaleFactor:
					// do nothing, to avoid resetting related points
				default:
					panic("unhandled type: " + p.Type())
				}
			})
			b.Write()
		})
	})
}
