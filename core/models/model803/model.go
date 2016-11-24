// NOTICE
// This file was automatically generated by ../../../generators/core.go. Do not edit it!
// You can regenerate it by running 'go generate ./core' from the directory above.

package model803

import (
	"github.com/crabmusket/gosunspec/core"
	"github.com/crabmusket/gosunspec/smdx"
)

// Block803 - Lithium-Ion Battery Bank Model -

const (
	ModelID = 803
)

const (
	A_SF            = "A_SF"
	CellV_SF        = "CellV_SF"
	ModTmpAvg       = "ModTmpAvg"
	ModTmpMax       = "ModTmpMax"
	ModTmpMaxMod    = "ModTmpMaxMod"
	ModTmpMaxStr    = "ModTmpMaxStr"
	ModTmpMin       = "ModTmpMin"
	ModTmpMinMod    = "ModTmpMinMod"
	ModTmpMinStr    = "ModTmpMinStr"
	ModTmp_SF       = "ModTmp_SF"
	NCellBal        = "NCellBal"
	NStr            = "NStr"
	NStrCon         = "NStrCon"
	Pad1            = "Pad1"
	Pad2            = "Pad2"
	Pad3            = "Pad3"
	Pad4            = "Pad4"
	Pad5            = "Pad5"
	SoH_SF          = "SoH_SF"
	StrA            = "StrA"
	StrAAvg         = "StrAAvg"
	StrAMax         = "StrAMax"
	StrAMaxStr      = "StrAMaxStr"
	StrAMin         = "StrAMin"
	StrAMinStr      = "StrAMinStr"
	StrCellVAvg     = "StrCellVAvg"
	StrCellVMax     = "StrCellVMax"
	StrCellVMaxMod  = "StrCellVMaxMod"
	StrCellVMin     = "StrCellVMin"
	StrCellVMinMod  = "StrCellVMinMod"
	StrConFail      = "StrConFail"
	StrConSt        = "StrConSt"
	StrEvt1         = "StrEvt1"
	StrEvt2         = "StrEvt2"
	StrEvtVnd1      = "StrEvtVnd1"
	StrEvtVnd2      = "StrEvtVnd2"
	StrModTmpAvg    = "StrModTmpAvg"
	StrModTmpMax    = "StrModTmpMax"
	StrModTmpMaxMod = "StrModTmpMaxMod"
	StrModTmpMin    = "StrModTmpMin"
	StrModTmpMinMod = "StrModTmpMinMod"
	StrNMod         = "StrNMod"
	StrSetCon       = "StrSetCon"
	StrSetEna       = "StrSetEna"
	StrSoC          = "StrSoC"
	StrSoH          = "StrSoH"
	StrSt           = "StrSt"
	StrVAvg         = "StrVAvg"
	StrVMax         = "StrVMax"
	StrVMaxStr      = "StrVMaxStr"
	StrVMin         = "StrVMin"
	StrVMinStr      = "StrVMinStr"
)

type Block803Repeat struct {
	StrNMod         uint16          `sunspec:"offset=0"`
	StrSt           core.Bitfield32 `sunspec:"offset=1"`
	StrConFail      core.Enum16     `sunspec:"offset=3"`
	StrSoC          uint16          `sunspec:"offset=4"`
	StrSoH          uint16          `sunspec:"offset=5,sf=SoH_SF"`
	StrA            int16           `sunspec:"offset=6,sf=A_SF"`
	StrCellVMax     uint16          `sunspec:"offset=7,sf=CellV_SF"`
	StrCellVMaxMod  uint16          `sunspec:"offset=8"`
	StrCellVMin     uint16          `sunspec:"offset=9,sf=CellV_SF"`
	StrCellVMinMod  uint16          `sunspec:"offset=10"`
	StrCellVAvg     uint16          `sunspec:"offset=11,sf=CellV_SF"`
	StrModTmpMax    int16           `sunspec:"offset=12,sf=ModTmp_SF"`
	StrModTmpMaxMod uint16          `sunspec:"offset=13"`
	StrModTmpMin    int16           `sunspec:"offset=14,sf=ModTmp_SF"`
	StrModTmpMinMod uint16          `sunspec:"offset=15"`
	StrModTmpAvg    int16           `sunspec:"offset=16,sf=ModTmp_SF"`
	Pad3            core.Pad        `sunspec:"offset=17"`
	StrConSt        core.Bitfield32 `sunspec:"offset=18"`
	StrEvt1         core.Bitfield32 `sunspec:"offset=20"`
	StrEvt2         core.Bitfield32 `sunspec:"offset=22"`
	StrEvtVnd1      core.Bitfield32 `sunspec:"offset=24"`
	StrEvtVnd2      core.Bitfield32 `sunspec:"offset=26"`
	StrSetEna       core.Enum16     `sunspec:"offset=28,access=rw"`
	StrSetCon       core.Enum16     `sunspec:"offset=29,access=rw"`
	Pad4            core.Pad        `sunspec:"offset=30"`
	Pad5            core.Pad        `sunspec:"offset=31"`
}

type Block803 struct {
	NStr         uint16           `sunspec:"offset=0"`
	NStrCon      uint16           `sunspec:"offset=1"`
	ModTmpMax    int16            `sunspec:"offset=2,sf=ModTmp_SF"`
	ModTmpMaxStr uint16           `sunspec:"offset=3"`
	ModTmpMaxMod uint16           `sunspec:"offset=4"`
	ModTmpMin    int16            `sunspec:"offset=5,sf=ModTmp_SF"`
	ModTmpMinStr uint16           `sunspec:"offset=6"`
	ModTmpMinMod uint16           `sunspec:"offset=7"`
	ModTmpAvg    uint16           `sunspec:"offset=8"`
	StrVMax      uint16           `sunspec:"offset=9,sf=V_SF"`
	StrVMaxStr   uint16           `sunspec:"offset=10"`
	StrVMin      uint16           `sunspec:"offset=11,sf=V_SF"`
	StrVMinStr   uint16           `sunspec:"offset=12"`
	StrVAvg      uint16           `sunspec:"offset=13,sf=V_SF"`
	StrAMax      int16            `sunspec:"offset=14,sf=A_SF"`
	StrAMaxStr   uint16           `sunspec:"offset=15"`
	StrAMin      int16            `sunspec:"offset=16,sf=A_SF"`
	StrAMinStr   uint16           `sunspec:"offset=17"`
	StrAAvg      int16            `sunspec:"offset=18,sf=A_SF"`
	NCellBal     uint16           `sunspec:"offset=19"`
	CellV_SF     core.ScaleFactor `sunspec:"offset=20"`
	ModTmp_SF    core.ScaleFactor `sunspec:"offset=21"`
	A_SF         core.ScaleFactor `sunspec:"offset=22"`
	SoH_SF       core.ScaleFactor `sunspec:"offset=23"`
	Pad1         core.Pad         `sunspec:"offset=24"`
	Pad2         core.Pad         `sunspec:"offset=25"`

	Repeats []Block803Repeat
}

func (self *Block803) GetId() core.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "lithium_ion_bank",
		Length: 33,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 26,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: NStr, Offset: 0, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: NStrCon, Offset: 1, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: ModTmpMax, Offset: 2, Type: "int16", ScaleFactor: "ModTmp_SF", Mandatory: true},
					smdx.PointElement{Id: ModTmpMaxStr, Offset: 3, Type: "uint16"},
					smdx.PointElement{Id: ModTmpMaxMod, Offset: 4, Type: "uint16"},
					smdx.PointElement{Id: ModTmpMin, Offset: 5, Type: "int16", ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true},
					smdx.PointElement{Id: ModTmpMinStr, Offset: 6, Type: "uint16"},
					smdx.PointElement{Id: ModTmpMinMod, Offset: 7, Type: "uint16"},
					smdx.PointElement{Id: ModTmpAvg, Offset: 8, Type: "uint16"},
					smdx.PointElement{Id: StrVMax, Offset: 9, Type: "uint16", ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: StrVMaxStr, Offset: 10, Type: "uint16"},
					smdx.PointElement{Id: StrVMin, Offset: 11, Type: "uint16", ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: StrVMinStr, Offset: 12, Type: "uint16"},
					smdx.PointElement{Id: StrVAvg, Offset: 13, Type: "uint16", ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: StrAMax, Offset: 14, Type: "int16", ScaleFactor: "A_SF", Units: "A"},
					smdx.PointElement{Id: StrAMaxStr, Offset: 15, Type: "uint16"},
					smdx.PointElement{Id: StrAMin, Offset: 16, Type: "int16", ScaleFactor: "A_SF", Units: "A"},
					smdx.PointElement{Id: StrAMinStr, Offset: 17, Type: "uint16"},
					smdx.PointElement{Id: StrAAvg, Offset: 18, Type: "int16", ScaleFactor: "A_SF", Units: "A"},
					smdx.PointElement{Id: NCellBal, Offset: 19, Type: "uint16"},
					smdx.PointElement{Id: CellV_SF, Offset: 20, Type: "sunssf", Mandatory: true},
					smdx.PointElement{Id: ModTmp_SF, Offset: 21, Type: "sunssf", Mandatory: true},
					smdx.PointElement{Id: A_SF, Offset: 22, Type: "sunssf", Mandatory: true},
					smdx.PointElement{Id: SoH_SF, Offset: 23, Type: "sunssf"},
					smdx.PointElement{Id: Pad1, Offset: 24, Type: "pad", Mandatory: true},
					smdx.PointElement{Id: Pad2, Offset: 25, Type: "pad", Mandatory: true},
				},
			},
			smdx.BlockElement{Name: "string",
				Length: 28,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: StrNMod, Offset: 0, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: StrSt, Offset: 1, Type: "bitfield32", Mandatory: true},
					smdx.PointElement{Id: StrConFail, Offset: 3, Type: "enum16"},
					smdx.PointElement{Id: StrSoC, Offset: 4, Type: "uint16", Units: "%", Mandatory: true},
					smdx.PointElement{Id: StrSoH, Offset: 5, Type: "uint16", ScaleFactor: "SoH_SF", Units: "%"},
					smdx.PointElement{Id: StrA, Offset: 6, Type: "int16", ScaleFactor: "A_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: StrCellVMax, Offset: 7, Type: "uint16", ScaleFactor: "CellV_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: StrCellVMaxMod, Offset: 8, Type: "uint16"},
					smdx.PointElement{Id: StrCellVMin, Offset: 9, Type: "uint16", ScaleFactor: "CellV_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: StrCellVMinMod, Offset: 10, Type: "uint16"},
					smdx.PointElement{Id: StrCellVAvg, Offset: 11, Type: "uint16", ScaleFactor: "CellV_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: StrModTmpMax, Offset: 12, Type: "int16", ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true},
					smdx.PointElement{Id: StrModTmpMaxMod, Offset: 13, Type: "uint16"},
					smdx.PointElement{Id: StrModTmpMin, Offset: 14, Type: "int16", ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true},
					smdx.PointElement{Id: StrModTmpMinMod, Offset: 15, Type: "uint16"},
					smdx.PointElement{Id: StrModTmpAvg, Offset: 16, Type: "int16", ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true},
					smdx.PointElement{Id: Pad3, Offset: 17, Type: "pad", Mandatory: true},
					smdx.PointElement{Id: StrConSt, Offset: 18, Type: "bitfield32"},
					smdx.PointElement{Id: StrEvt1, Offset: 20, Type: "bitfield32", Mandatory: true},
					smdx.PointElement{Id: StrEvt2, Offset: 22, Type: "bitfield32"},
					smdx.PointElement{Id: StrEvtVnd1, Offset: 24, Type: "bitfield32"},
					smdx.PointElement{Id: StrEvtVnd2, Offset: 26, Type: "bitfield32"},
					smdx.PointElement{Id: StrSetEna, Offset: 28, Type: "enum16", Access: "rw"},
					smdx.PointElement{Id: StrSetCon, Offset: 29, Type: "enum16", Access: "rw"},
					smdx.PointElement{Id: Pad4, Offset: 30, Type: "pad", Mandatory: true},
					smdx.PointElement{Id: Pad5, Offset: 31, Type: "pad", Mandatory: true},
				},
			},
		}})
}
