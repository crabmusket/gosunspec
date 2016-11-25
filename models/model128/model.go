// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model128

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block128 - Dynamic Reactive Current - Dynamic Reactive Current

const (
	ModelID = 128
)

const (
	ArGraMod   = "ArGraMod"
	ArGraSag   = "ArGraSag"
	ArGraSwell = "ArGraSwell"
	ArGra_SF   = "ArGra_SF"
	BlkZnTmms  = "BlkZnTmms"
	BlkZnV     = "BlkZnV"
	DbVMax     = "DbVMax"
	DbVMin     = "DbVMin"
	FilTms     = "FilTms"
	HoldTmms   = "HoldTmms"
	HysBlkZnV  = "HysBlkZnV"
	ModEna     = "ModEna"
	Pad        = "Pad"
	VRefPct_SF = "VRefPct_SF"
)

type Block128 struct {
	ArGraMod   sunspec.Enum16      `sunspec:"offset=0,len=1,access=rw"`
	ArGraSag   uint16              `sunspec:"offset=1,len=1,sf=ArGra_SF,access=rw"`
	ArGraSwell uint16              `sunspec:"offset=2,len=1,sf=ArGra_SF,access=rw"`
	ModEna     sunspec.Bitfield16  `sunspec:"offset=3,len=1,access=rw"`
	FilTms     uint16              `sunspec:"offset=4,len=1,access=rw"`
	DbVMin     uint16              `sunspec:"offset=5,len=1,sf=VRefPct_SF,access=rw"`
	DbVMax     uint16              `sunspec:"offset=6,len=1,sf=VRefPct_SF,access=rw"`
	BlkZnV     uint16              `sunspec:"offset=7,len=1,sf=VRefPct_SF,access=rw"`
	HysBlkZnV  uint16              `sunspec:"offset=8,len=1,sf=VRefPct_SF,access=rw"`
	BlkZnTmms  uint16              `sunspec:"offset=9,len=1,access=rw"`
	HoldTmms   uint16              `sunspec:"offset=10,len=1,access=rw"`
	ArGra_SF   sunspec.ScaleFactor `sunspec:"offset=11,len=1,access=r"`
	VRefPct_SF sunspec.ScaleFactor `sunspec:"offset=12,len=1,access=r"`
	Pad        sunspec.Pad         `sunspec:"offset=13,len=1,access=r"`
}

func (self *Block128) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "reactive_current",
		Length: 14,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 14,
				Type:   "fixed",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: ArGraMod, Offset: 0, Type: typelabel.Enum16, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: ArGraSag, Offset: 1, Type: typelabel.Uint16, ScaleFactor: "ArGra_SF", Units: "%ARtg/%dV", Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: ArGraSwell, Offset: 2, Type: typelabel.Uint16, ScaleFactor: "ArGra_SF", Units: "%ARtg/%dV", Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: ModEna, Offset: 3, Type: typelabel.Bitfield16, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: FilTms, Offset: 4, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: DbVMin, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "VRefPct_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: DbVMax, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "VRefPct_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: BlkZnV, Offset: 7, Type: typelabel.Uint16, ScaleFactor: "VRefPct_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: HysBlkZnV, Offset: 8, Type: typelabel.Uint16, ScaleFactor: "VRefPct_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: BlkZnTmms, Offset: 9, Type: typelabel.Uint16, Units: "mSecs", Access: "rw", Length: 1},
					smdx.PointElement{Id: HoldTmms, Offset: 10, Type: typelabel.Uint16, Units: "mSecs", Access: "rw", Length: 1},
					smdx.PointElement{Id: ArGra_SF, Offset: 11, Type: typelabel.Sunssf, Access: "r", Length: 1, Mandatory: true},
					smdx.PointElement{Id: VRefPct_SF, Offset: 12, Type: typelabel.Sunssf, Access: "r", Length: 1},
					smdx.PointElement{Id: Pad, Offset: 13, Type: typelabel.Pad, Access: "r", Length: 1},
				},
			},
		}})
}
