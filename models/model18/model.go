// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model18

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block18 - Cellular Link - Include this model to support a cellular interface link

const (
	ModelID          = 18
	ModelLabel       = "Cellular Link"
	ModelDescription = "Include this model to support a cellular interface link"
)

const (
	APN  = "APN"
	IMEI = "IMEI"
	Nam  = "Nam"
	Num  = "Num"
	Pin  = "Pin"
)

type Block18 struct {
	Nam  string `sunspec:"offset=0,len=4,access=rw"`
	IMEI uint32 `sunspec:"offset=4,access=rw"`
	APN  string `sunspec:"offset=6,len=4,access=rw"`
	Num  string `sunspec:"offset=10,len=6,access=rw"`
	Pin  string `sunspec:"offset=16,len=6,access=rw"`
}

func (self *Block18) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 22,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 22,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: Nam, Offset: 0, Type: typelabel.String, Access: "rw", Length: 4, Label: "Name", Description: "Interface name"},
					smdx.PointElement{Id: IMEI, Offset: 4, Type: typelabel.Uint32, Access: "rw", Label: "IMEI", Description: "International Mobile Equipment Identifier for the interface"},
					smdx.PointElement{Id: APN, Offset: 6, Type: typelabel.String, Access: "rw", Length: 4, Label: "APN", Description: "Access Point Name for the interface"},
					smdx.PointElement{Id: Num, Offset: 10, Type: typelabel.String, Access: "rw", Length: 6, Label: "Number", Description: "Phone number for the interface"},
					smdx.PointElement{Id: Pin, Offset: 16, Type: typelabel.String, Access: "rw", Length: 6, Label: "PIN", Description: "Personal Identification Number for the interface"},
				},
			},
		}})
}
