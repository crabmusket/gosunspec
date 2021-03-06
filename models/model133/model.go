// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model133

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block133 - Basic Scheduling - Basic Scheduling

const (
	ModelID = 133
)

const (
	ActIndx = "ActIndx"
	ActPts  = "ActPts"
	ActSchd = "ActSchd"
	IntvTyp = "IntvTyp"
	ModEna  = "ModEna"
	NPts    = "NPts"
	NSchd   = "NSchd"
	Nam     = "Nam"
	Pad     = "Pad"
	RepPer  = "RepPer"
	RmpTms  = "RmpTms"
	StrTms  = "StrTms"
	WinTms  = "WinTms"
	X1      = "X1"
	X10     = "X10"
	X2      = "X2"
	X3      = "X3"
	X4      = "X4"
	X5      = "X5"
	X6      = "X6"
	X7      = "X7"
	X8      = "X8"
	X9      = "X9"
	XTyp    = "XTyp"
	X_SF    = "X_SF"
	Y1      = "Y1"
	Y10     = "Y10"
	Y2      = "Y2"
	Y3      = "Y3"
	Y4      = "Y4"
	Y5      = "Y5"
	Y6      = "Y6"
	Y7      = "Y7"
	Y8      = "Y8"
	Y9      = "Y9"
	YTyp    = "YTyp"
	Y_SF    = "Y_SF"
)

type Block133Repeat struct {
	ActPts  uint16              `sunspec:"offset=0,len=1,access=rw"`
	StrTms  uint32              `sunspec:"offset=1,len=2,access=rw"`
	RepPer  uint16              `sunspec:"offset=3,len=1,access=rw"`
	IntvTyp sunspec.Enum16      `sunspec:"offset=4,len=1,access=rw"`
	XTyp    sunspec.Enum16      `sunspec:"offset=5,len=1,access=rw"`
	X_SF    sunspec.ScaleFactor `sunspec:"offset=6,len=1,access=rw"`
	YTyp    sunspec.Enum16      `sunspec:"offset=7,len=1,access=rw"`
	Y_SF    sunspec.ScaleFactor `sunspec:"offset=8,len=1,access=rw"`
	X1      int32               `sunspec:"offset=9,len=2,sf=X_SF,access=rw"`
	Y1      int32               `sunspec:"offset=11,len=2,sf=Y_SF,access=rw"`
	X2      int32               `sunspec:"offset=13,len=2,sf=X_SF,access=rw"`
	Y2      int32               `sunspec:"offset=15,len=2,sf=Y_SF,access=rw"`
	X3      int32               `sunspec:"offset=17,len=2,sf=X_SF,access=rw"`
	Y3      int32               `sunspec:"offset=19,len=2,sf=Y_SF,access=rw"`
	X4      int32               `sunspec:"offset=21,len=2,sf=X_SF,access=rw"`
	Y4      int32               `sunspec:"offset=23,len=2,sf=Y_SF,access=rw"`
	X5      int32               `sunspec:"offset=25,len=2,sf=X_SF,access=rw"`
	Y5      int32               `sunspec:"offset=27,len=2,sf=Y_SF,access=rw"`
	X6      int32               `sunspec:"offset=29,len=2,sf=X_SF,access=rw"`
	Y6      int32               `sunspec:"offset=31,len=2,sf=Y_SF,access=rw"`
	X7      int32               `sunspec:"offset=33,len=2,sf=X_SF,access=rw"`
	Y7      int32               `sunspec:"offset=35,len=2,sf=Y_SF,access=rw"`
	X8      int32               `sunspec:"offset=37,len=2,sf=X_SF,access=rw"`
	Y8      int32               `sunspec:"offset=39,len=2,sf=Y_SF,access=rw"`
	X9      int32               `sunspec:"offset=41,len=2,sf=X_SF,access=rw"`
	Y9      int32               `sunspec:"offset=43,len=2,sf=Y_SF,access=rw"`
	X10     int32               `sunspec:"offset=45,len=2,sf=X_SF,access=rw"`
	Y10     int32               `sunspec:"offset=47,len=2,sf=Y_SF,access=rw"`
	Nam     string              `sunspec:"offset=49,len=8,access=rw"`
	WinTms  uint16              `sunspec:"offset=57,len=1,access=rw"`
	RmpTms  uint16              `sunspec:"offset=58,len=1,access=rw"`
	ActIndx uint16              `sunspec:"offset=59,len=1,access=r"`
}

type Block133 struct {
	ActSchd sunspec.Bitfield32 `sunspec:"offset=0,len=2,access=rw"`
	ModEna  sunspec.Bitfield16 `sunspec:"offset=2,len=1,access=rw"`
	NSchd   uint16             `sunspec:"offset=3,len=1,access=r"`
	NPts    uint16             `sunspec:"offset=4,len=1,access=r"`
	Pad     sunspec.Pad        `sunspec:"offset=5,len=1,access=r"`

	Repeats []Block133Repeat
}

func (self *Block133) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "schedule",
		Length: 66,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 6,
				Type:   "fixed",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: ActSchd, Offset: 0, Type: typelabel.Bitfield32, Access: "rw", Length: 2, Mandatory: true},
					smdx.PointElement{Id: ModEna, Offset: 2, Type: typelabel.Bitfield16, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: NSchd, Offset: 3, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true},
					smdx.PointElement{Id: NPts, Offset: 4, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true},
					smdx.PointElement{Id: Pad, Offset: 5, Type: typelabel.Pad, Access: "r", Length: 1},
				},
			},
			smdx.BlockElement{
				Length: 60,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: ActPts, Offset: 0, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: StrTms, Offset: 1, Type: typelabel.Uint32, Units: "Secs", Access: "rw", Length: 2, Mandatory: true},
					smdx.PointElement{Id: RepPer, Offset: 3, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: IntvTyp, Offset: 4, Type: typelabel.Enum16, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: XTyp, Offset: 5, Type: typelabel.Enum16, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: X_SF, Offset: 6, Type: typelabel.ScaleFactor, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: YTyp, Offset: 7, Type: typelabel.Enum16, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: Y_SF, Offset: 8, Type: typelabel.ScaleFactor, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: X1, Offset: 9, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2, Mandatory: true},
					smdx.PointElement{Id: Y1, Offset: 11, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2, Mandatory: true},
					smdx.PointElement{Id: X2, Offset: 13, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: Y2, Offset: 15, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: X3, Offset: 17, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: Y3, Offset: 19, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: X4, Offset: 21, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: Y4, Offset: 23, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: X5, Offset: 25, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: Y5, Offset: 27, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: X6, Offset: 29, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: Y6, Offset: 31, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: X7, Offset: 33, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: Y7, Offset: 35, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: X8, Offset: 37, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: Y8, Offset: 39, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: X9, Offset: 41, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: Y9, Offset: 43, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: X10, Offset: 45, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: Y10, Offset: 47, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2},
					smdx.PointElement{Id: Nam, Offset: 49, Type: typelabel.String, Access: "rw", Length: 8},
					smdx.PointElement{Id: WinTms, Offset: 57, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: RmpTms, Offset: 58, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: ActIndx, Offset: 59, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true},
				},
			},
		}})
}
