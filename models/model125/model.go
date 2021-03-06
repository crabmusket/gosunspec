// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model125

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block125 - Pricing - Pricing Signal

const (
	ModelID = 125
)

const (
	ModEna  = "ModEna"
	Pad     = "Pad"
	RmpTms  = "RmpTms"
	RvtTms  = "RvtTms"
	Sig     = "Sig"
	SigType = "SigType"
	Sig_SF  = "Sig_SF"
	WinTms  = "WinTms"
)

type Block125 struct {
	ModEna  sunspec.Bitfield16  `sunspec:"offset=0,len=1,access=rw"`
	SigType sunspec.Enum16      `sunspec:"offset=1,len=1,sf= ,access=rw"`
	Sig     int16               `sunspec:"offset=2,len=1,sf=Sig_SF,access=rw"`
	WinTms  uint16              `sunspec:"offset=3,len=1,access=rw"`
	RvtTms  uint16              `sunspec:"offset=4,len=1,access=rw"`
	RmpTms  uint16              `sunspec:"offset=5,len=1,access=rw"`
	Sig_SF  sunspec.ScaleFactor `sunspec:"offset=6,len=1,sf= ,access=r"`
	Pad     sunspec.Pad         `sunspec:"offset=7,len=1,access=r"`
}

func (self *Block125) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "pricing",
		Length: 8,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 8,
				Type:   "fixed",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: ModEna, Offset: 0, Type: typelabel.Bitfield16, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: SigType, Offset: 1, Type: typelabel.Enum16, ScaleFactor: " ", Access: "rw", Length: 1},
					smdx.PointElement{Id: Sig, Offset: 2, Type: typelabel.Int16, ScaleFactor: "Sig_SF", Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: WinTms, Offset: 3, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: RvtTms, Offset: 4, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: RmpTms, Offset: 5, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Sig_SF, Offset: 6, Type: typelabel.ScaleFactor, ScaleFactor: " ", Access: "r", Length: 1, Mandatory: true},
					smdx.PointElement{Id: Pad, Offset: 7, Type: typelabel.Pad, Access: "r", Length: 1},
				},
			},
		}})
}
