// NOTICE
// This file was automatically generated by ../../../generators/core.go. Do not edit it!
// You can regenerate it by running 'go generate ./core' from the directory above.

package model10

import (
	"github.com/crabmusket/gosunspec/core"
	"github.com/crabmusket/gosunspec/smdx"
)

// Block10 - Communication Interface Header - To be included first for a complete interface description

const (
	ModelID = 10
)

const (
	Ctl = "Ctl"
	Pad = "Pad"
	St  = "St"
	Typ = "Typ"
)

type Block10 struct {
	St  core.Enum16 `sunspec:"offset=0"`
	Ctl uint16      `sunspec:"offset=1,access=rw"`
	Typ core.Enum16 `sunspec:"offset=2"`
	Pad core.Pad    `sunspec:"offset=3"`
}

func (self *Block10) GetId() core.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 4,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 4,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: St, Offset: 0, Type: "enum16", Mandatory: true},
					smdx.PointElement{Id: Ctl, Offset: 1, Type: "uint16", Access: "rw"},
					smdx.PointElement{Id: Typ, Offset: 2, Type: "enum16"},
					smdx.PointElement{Id: Pad, Offset: 3, Type: "pad"},
				},
			},
		}})
}
