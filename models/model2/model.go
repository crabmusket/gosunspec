// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model2

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block2 - Basic Aggregator - Aggregates a collection of models for a given model id

const (
	ModelID          = 2
	ModelLabel       = "Basic Aggregator"
	ModelDescription = "Aggregates a collection of models for a given model id"
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
	AID    uint16             `sunspec:"offset=0"`
	N      uint16             `sunspec:"offset=1"`
	UN     uint16             `sunspec:"offset=2"`
	St     sunspec.Enum16     `sunspec:"offset=3"`
	StVnd  sunspec.Enum16     `sunspec:"offset=4"`
	Evt    sunspec.Bitfield32 `sunspec:"offset=5"`
	EvtVnd sunspec.Bitfield32 `sunspec:"offset=7"`
	Ctl    sunspec.Enum16     `sunspec:"offset=9"`
	CtlVnd sunspec.Enum32     `sunspec:"offset=10"`
	CtlVl  sunspec.Enum32     `sunspec:"offset=12"`
}

func (self *Block2) GetId() sunspec.ModelId {
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
					smdx.PointElement{Id: AID, Offset: 0, Type: typelabel.Uint16, Mandatory: true, Label: "AID", Description: "Aggregated model id"},
					smdx.PointElement{Id: N, Offset: 1, Type: typelabel.Uint16, Mandatory: true, Label: "N", Description: "Number of aggregated models"},
					smdx.PointElement{Id: UN, Offset: 2, Type: typelabel.Uint16, Mandatory: true, Label: "UN", Description: "Update Number.  Incrementing nunber each time the mappping is changed.  If the number is not changed from thelast reading the direct access to a specific offset will result in reading the same logical model as before.  Otherwise the entire model must be read to refresh the changes"},
					smdx.PointElement{Id: St, Offset: 3, Type: typelabel.Enum16, Mandatory: true, Label: "Status", Description: "Enumerated status code"},
					smdx.PointElement{Id: StVnd, Offset: 4, Type: typelabel.Enum16, Label: "Vendor Status", Description: "Vendor specific status code"},
					smdx.PointElement{Id: Evt, Offset: 5, Type: typelabel.Bitfield32, Mandatory: true, Label: "Event Code", Description: "Bitmask event code"},
					smdx.PointElement{Id: EvtVnd, Offset: 7, Type: typelabel.Bitfield32, Label: "Vendor Event Code", Description: "Vendor specific event code"},
					smdx.PointElement{Id: Ctl, Offset: 9, Type: typelabel.Enum16, Label: "Control", Description: "Control register for all aggregated devices"},
					smdx.PointElement{Id: CtlVnd, Offset: 10, Type: typelabel.Enum32, Label: "Vendor Control", Description: "Vendor control register for all aggregated devices"},
					smdx.PointElement{Id: CtlVl, Offset: 12, Type: typelabel.Enum32, Label: "Control Value", Description: "Numerical value used as a parameter to the control"},
				},
			},
		}})
}
