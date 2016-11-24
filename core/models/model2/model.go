// NOTICE
// This file was automatically generated by ../../../generators/core.go. Do not edit it!
// You can regenerate it by running 'go generate ./core' from the directory above.

package model2

import (
	"github.com/crabmusket/gosunspec/core"
	"github.com/crabmusket/gosunspec/smdx"
)

// Block2 - Basic Aggregator - Aggregates a collection of models for a given model id

const (
	ModelID = 2
)

const (
	AID    = "AID"
	Ctl    = "Ctl"
	CtlVl  = "CtlVl"
	CtlVnd = "CtlVnd"
	Evt    = "Evt"
	EvtVnd = "EvtVnd"
	N      = "N"
	St     = "St"
	StVnd  = "StVnd"
	UN     = "UN"
)

type Block2 struct {
	AID    uint16          `sunspec:"offset=0"`
	N      uint16          `sunspec:"offset=1"`
	UN     uint16          `sunspec:"offset=2"`
	St     core.Enum16     `sunspec:"offset=3"`
	StVnd  core.Enum16     `sunspec:"offset=4"`
	Evt    core.Bitfield32 `sunspec:"offset=5"`
	EvtVnd core.Bitfield32 `sunspec:"offset=7"`
	Ctl    core.Enum16     `sunspec:"offset=9"`
	CtlVnd core.Enum32     `sunspec:"offset=10"`
	CtlVl  core.Enum32     `sunspec:"offset=12"`
}

func (self *Block2) GetId() core.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "aggregator",
		Length: 14,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 14,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: AID, Offset: 0, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: N, Offset: 1, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: UN, Offset: 2, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: St, Offset: 3, Type: "enum16", Mandatory: true},
					smdx.PointElement{Id: StVnd, Offset: 4, Type: "enum16"},
					smdx.PointElement{Id: Evt, Offset: 5, Type: "bitfield32", Mandatory: true},
					smdx.PointElement{Id: EvtVnd, Offset: 7, Type: "bitfield32"},
					smdx.PointElement{Id: Ctl, Offset: 9, Type: "enum16"},
					smdx.PointElement{Id: CtlVnd, Offset: 10, Type: "enum32"},
					smdx.PointElement{Id: CtlVl, Offset: 12, Type: "enum32"},
				},
			},
		}})
}
