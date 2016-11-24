// NOTICE
// This file was automatically generated by ../../../generators/core.go. Do not edit it!
// You can regenerate it by running 'go generate ./core' from the directory above.

package model601

import (
	"github.com/crabmusket/gosunspec/core"
	"github.com/crabmusket/gosunspec/smdx"
)

// Block601 - Tracker Controller DRAFT 2 - Monitors and controls multiple trackers

const (
	ModelID = 601
)

const (
	Alm       = "Alm"
	AzCtl     = "AzCtl"
	AzPos     = "AzPos"
	AzTrgt    = "AzTrgt"
	Ctl       = "Ctl"
	Day       = "Day"
	Dgr_SF    = "Dgr_SF"
	DtLoc     = "DtLoc"
	ElCtl     = "ElCtl"
	ElPos     = "ElPos"
	ElTrgt    = "ElTrgt"
	GlblAlm   = "GlblAlm"
	GlblAzCtl = "GlblAzCtl"
	GlblCtl   = "GlblCtl"
	GlblElCtl = "GlblElCtl"
	Id        = "Id"
	N         = "N"
	Nam       = "Nam"
	TmLoc     = "TmLoc"
	Typ       = "Typ"
)

type Block601Repeat struct {
	Id     core.String     `sunspec:"offset=0,len=8"`
	ElTrgt int32           `sunspec:"offset=8,sf=Dgr_SF"`
	AzTrgt int32           `sunspec:"offset=10,sf=SF"`
	ElPos  int32           `sunspec:"offset=12,sf=Dgr_SF"`
	AzPos  int32           `sunspec:"offset=14,sf=Dgr_SF"`
	ElCtl  int32           `sunspec:"offset=16,sf=Dgr_SF,access=rw"`
	AzCtl  int32           `sunspec:"offset=18,sf=Dgr_SF,access=rw"`
	Ctl    core.Enum16     `sunspec:"offset=20,access=rw"`
	Alm    core.Bitfield16 `sunspec:"offset=21"`
}

type Block601 struct {
	Nam       core.String      `sunspec:"offset=0,len=8"`
	Typ       core.Enum16      `sunspec:"offset=8"`
	DtLoc     core.String      `sunspec:"offset=9,len=5"`
	TmLoc     core.String      `sunspec:"offset=14,len=3"`
	Day       uint16           `sunspec:"offset=17"`
	GlblElCtl int32            `sunspec:"offset=18,sf=Dgr_SF,access=rw"`
	GlblAzCtl int32            `sunspec:"offset=20,sf=Dgr_SF,access=rw"`
	GlblCtl   core.Enum16      `sunspec:"offset=22,access=rw"`
	GlblAlm   core.Bitfield16  `sunspec:"offset=23"`
	Dgr_SF    core.ScaleFactor `sunspec:"offset=24"`
	N         uint16           `sunspec:"offset=25"`

	Repeats []Block601Repeat
}

func (self *Block601) GetId() core.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "tracker_controller",
		Length: 48,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 26,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: Nam, Offset: 0, Type: "string", Length: 8},
					smdx.PointElement{Id: Typ, Offset: 8, Type: "enum16", Mandatory: true},
					smdx.PointElement{Id: DtLoc, Offset: 9, Type: "string", Units: "YYYYMMDD", Length: 5},
					smdx.PointElement{Id: TmLoc, Offset: 14, Type: "string", Units: "hhmmss", Length: 3},
					smdx.PointElement{Id: Day, Offset: 17, Type: "uint16"},
					smdx.PointElement{Id: GlblElCtl, Offset: 18, Type: "int32", ScaleFactor: "Dgr_SF", Units: "Degrees", Access: "rw"},
					smdx.PointElement{Id: GlblAzCtl, Offset: 20, Type: "int32", ScaleFactor: "Dgr_SF", Units: "Degrees", Access: "rw"},
					smdx.PointElement{Id: GlblCtl, Offset: 22, Type: "enum16", Access: "rw"},
					smdx.PointElement{Id: GlblAlm, Offset: 23, Type: "bitfield16"},
					smdx.PointElement{Id: Dgr_SF, Offset: 24, Type: "sunssf", Mandatory: true},
					smdx.PointElement{Id: N, Offset: 25, Type: "uint16", Mandatory: true},
				},
			},
			smdx.BlockElement{Name: "tracker",
				Length: 22,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: Id, Offset: 0, Type: "string", Length: 8},
					smdx.PointElement{Id: ElTrgt, Offset: 8, Type: "int32", ScaleFactor: "Dgr_SF", Units: "Degrees"},
					smdx.PointElement{Id: AzTrgt, Offset: 10, Type: "int32", ScaleFactor: "SF", Units: "Degrees"},
					smdx.PointElement{Id: ElPos, Offset: 12, Type: "int32", ScaleFactor: "Dgr_SF", Units: "Degrees"},
					smdx.PointElement{Id: AzPos, Offset: 14, Type: "int32", ScaleFactor: "Dgr_SF", Units: "Degrees"},
					smdx.PointElement{Id: ElCtl, Offset: 16, Type: "int32", ScaleFactor: "Dgr_SF", Units: "Degrees", Access: "rw"},
					smdx.PointElement{Id: AzCtl, Offset: 18, Type: "int32", ScaleFactor: "Dgr_SF", Units: "Degrees", Access: "rw"},
					smdx.PointElement{Id: Ctl, Offset: 20, Type: "enum16", Access: "rw"},
					smdx.PointElement{Id: Alm, Offset: 21, Type: "bitfield16"},
				},
			},
		}})
}
