// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model502

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block502 - Solar Module - A solar module model supporing DC-DC converter

const (
	ModelID = 502
)

const (
	A_SF     = "A_SF"
	Ctl      = "Ctl"
	CtlVal   = "CtlVal"
	CtlVend  = "CtlVend"
	Evt      = "Evt"
	EvtVend  = "EvtVend"
	InA      = "InA"
	InV      = "InV"
	InW      = "InW"
	InWh     = "InWh"
	OutA     = "OutA"
	OutPw    = "OutPw"
	OutV     = "OutV"
	OutWh    = "OutWh"
	Stat     = "Stat"
	StatVend = "StatVend"
	Tmp      = "Tmp"
	Tms      = "Tms"
	V_SF     = "V_SF"
	W_SF     = "W_SF"
	Wh_SF    = "Wh_SF"
)

type Block502 struct {
	A_SF     sunspec.ScaleFactor `sunspec:"offset=0"`
	V_SF     sunspec.ScaleFactor `sunspec:"offset=1"`
	W_SF     sunspec.ScaleFactor `sunspec:"offset=2"`
	Wh_SF    sunspec.ScaleFactor `sunspec:"offset=3"`
	Stat     sunspec.Enum16      `sunspec:"offset=4"`
	StatVend sunspec.Enum16      `sunspec:"offset=5"`
	Evt      sunspec.Bitfield32  `sunspec:"offset=6"`
	EvtVend  sunspec.Bitfield32  `sunspec:"offset=8"`
	Ctl      sunspec.Enum16      `sunspec:"offset=10,access=rw"`
	CtlVend  sunspec.Enum32      `sunspec:"offset=11,access=rw"`
	CtlVal   int32               `sunspec:"offset=13,access=rw"`
	Tms      uint32              `sunspec:"offset=15"`
	OutA     int16               `sunspec:"offset=17,sf=A_SF"`
	OutV     int16               `sunspec:"offset=18,sf=V_SF"`
	OutWh    sunspec.Acc32       `sunspec:"offset=19,sf=Wh_SF"`
	OutPw    int16               `sunspec:"offset=21,sf=W_SF"`
	Tmp      int16               `sunspec:"offset=22"`
	InA      int16               `sunspec:"offset=23,sf=A_SF"`
	InV      int16               `sunspec:"offset=24,sf=V_SF"`
	InWh     sunspec.Acc32       `sunspec:"offset=25,sf=Wh_SF"`
	InW      int16               `sunspec:"offset=27,sf=W_SF"`
}

func (self *Block502) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "solar_module",
		Length: 28,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 28,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: A_SF, Offset: 0, Type: typelabel.Sunssf},
					smdx.PointElement{Id: V_SF, Offset: 1, Type: typelabel.Sunssf},
					smdx.PointElement{Id: W_SF, Offset: 2, Type: typelabel.Sunssf},
					smdx.PointElement{Id: Wh_SF, Offset: 3, Type: typelabel.Sunssf},
					smdx.PointElement{Id: Stat, Offset: 4, Type: typelabel.Enum16, Mandatory: true},
					smdx.PointElement{Id: StatVend, Offset: 5, Type: typelabel.Enum16},
					smdx.PointElement{Id: Evt, Offset: 6, Type: typelabel.Bitfield32, Mandatory: true},
					smdx.PointElement{Id: EvtVend, Offset: 8, Type: typelabel.Bitfield32},
					smdx.PointElement{Id: Ctl, Offset: 10, Type: typelabel.Enum16, Access: "rw"},
					smdx.PointElement{Id: CtlVend, Offset: 11, Type: typelabel.Enum32, Access: "rw"},
					smdx.PointElement{Id: CtlVal, Offset: 13, Type: typelabel.Int32, Access: "rw"},
					smdx.PointElement{Id: Tms, Offset: 15, Type: typelabel.Uint32, Units: "Secs"},
					smdx.PointElement{Id: OutA, Offset: 17, Type: typelabel.Int16, ScaleFactor: "A_SF", Units: "A"},
					smdx.PointElement{Id: OutV, Offset: 18, Type: typelabel.Int16, ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: OutWh, Offset: 19, Type: typelabel.Acc32, ScaleFactor: "Wh_SF", Units: "Wh"},
					smdx.PointElement{Id: OutPw, Offset: 21, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "W"},
					smdx.PointElement{Id: Tmp, Offset: 22, Type: typelabel.Int16, Units: "C"},
					smdx.PointElement{Id: InA, Offset: 23, Type: typelabel.Int16, ScaleFactor: "A_SF", Units: "A"},
					smdx.PointElement{Id: InV, Offset: 24, Type: typelabel.Int16, ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: InWh, Offset: 25, Type: typelabel.Acc32, ScaleFactor: "Wh_SF", Units: "Wh"},
					smdx.PointElement{Id: InW, Offset: 27, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "W"},
				},
			},
		}})
}