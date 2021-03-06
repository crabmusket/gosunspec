// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model220

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block220 - Secure AC Meter Selected Readings - Include this model for secure metering

const (
	ModelID = 220
)

const (
	A            = "A"
	A_SF         = "A_SF"
	Alg          = "Alg"
	DS           = "DS"
	Evt          = "Evt"
	Hz           = "Hz"
	Hz_SF        = "Hz_SF"
	Ms           = "Ms"
	N            = "N"
	PF           = "PF"
	PF_SF        = "PF_SF"
	PhV          = "PhV"
	Rsrvd        = "Rsrvd"
	Seq          = "Seq"
	TotVAhExp    = "TotVAhExp"
	TotVAhImp    = "TotVAhImp"
	TotVAh_SF    = "TotVAh_SF"
	TotVArhExpQ3 = "TotVArhExpQ3"
	TotVArhExpQ4 = "TotVArhExpQ4"
	TotVArhImpQ1 = "TotVArhImpQ1"
	TotVArhImpQ2 = "TotVArhImpQ2"
	TotVArh_SF   = "TotVArh_SF"
	TotWhExp     = "TotWhExp"
	TotWhImp     = "TotWhImp"
	TotWh_SF     = "TotWh_SF"
	Ts           = "Ts"
	VA           = "VA"
	VAR          = "VAR"
	VAR_SF       = "VAR_SF"
	VA_SF        = "VA_SF"
	V_SF         = "V_SF"
	W            = "W"
	W_SF         = "W_SF"
)

type Block220Repeat struct {
	DS uint16 `sunspec:"offset=0,access=r"`
}

type Block220 struct {
	A            int16               `sunspec:"offset=0,sf=A_SF"`
	A_SF         sunspec.ScaleFactor `sunspec:"offset=1"`
	PhV          int16               `sunspec:"offset=2,sf=V_SF"`
	V_SF         sunspec.ScaleFactor `sunspec:"offset=3"`
	Hz           int16               `sunspec:"offset=4,sf=Hz_SF"`
	Hz_SF        sunspec.ScaleFactor `sunspec:"offset=5"`
	W            int16               `sunspec:"offset=6,sf=W_SF"`
	W_SF         sunspec.ScaleFactor `sunspec:"offset=7"`
	VA           int16               `sunspec:"offset=8,sf=VA_SF"`
	VA_SF        sunspec.ScaleFactor `sunspec:"offset=9"`
	VAR          int16               `sunspec:"offset=10,sf=VAR_SF"`
	VAR_SF       sunspec.ScaleFactor `sunspec:"offset=11"`
	PF           int16               `sunspec:"offset=12,sf=PF_SF"`
	PF_SF        sunspec.ScaleFactor `sunspec:"offset=13"`
	TotWhExp     sunspec.Acc32       `sunspec:"offset=14,sf=TotWh_SF"`
	TotWhImp     sunspec.Acc32       `sunspec:"offset=16,sf=TotWh_SF"`
	TotWh_SF     sunspec.ScaleFactor `sunspec:"offset=18"`
	TotVAhExp    sunspec.Acc32       `sunspec:"offset=19,sf=TotVAh_SF"`
	TotVAhImp    sunspec.Acc32       `sunspec:"offset=21,sf=TotVAh_SF"`
	TotVAh_SF    sunspec.ScaleFactor `sunspec:"offset=23"`
	TotVArhImpQ1 sunspec.Acc32       `sunspec:"offset=24,sf=TotVArh_SF"`
	TotVArhImpQ2 sunspec.Acc32       `sunspec:"offset=26,sf=TotVArh_SF"`
	TotVArhExpQ3 sunspec.Acc32       `sunspec:"offset=28,sf=TotVArh_SF"`
	TotVArhExpQ4 sunspec.Acc32       `sunspec:"offset=30,sf=TotVArh_SF"`
	TotVArh_SF   sunspec.ScaleFactor `sunspec:"offset=32"`
	Evt          sunspec.Bitfield32  `sunspec:"offset=33"`
	Rsrvd        sunspec.Pad         `sunspec:"offset=35,access=r"`
	Ts           uint32              `sunspec:"offset=36,access=r"`
	Ms           uint16              `sunspec:"offset=38,access=r"`
	Seq          uint16              `sunspec:"offset=39,access=r"`
	Alg          sunspec.Enum16      `sunspec:"offset=40,access=r"`
	N            uint16              `sunspec:"offset=41,access=r"`

	Repeats []Block220Repeat
}

func (self *Block220) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "ac_meter",
		Length: 43,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 42,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: A, Offset: 0, Type: typelabel.Int16, ScaleFactor: "A_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: A_SF, Offset: 1, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: PhV, Offset: 2, Type: typelabel.Int16, ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: V_SF, Offset: 3, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: Hz, Offset: 4, Type: typelabel.Int16, ScaleFactor: "Hz_SF", Units: "Hz", Mandatory: true},
					smdx.PointElement{Id: Hz_SF, Offset: 5, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: W, Offset: 6, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "W", Mandatory: true},
					smdx.PointElement{Id: W_SF, Offset: 7, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: VA, Offset: 8, Type: typelabel.Int16, ScaleFactor: "VA_SF", Units: "VA"},
					smdx.PointElement{Id: VA_SF, Offset: 9, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: VAR, Offset: 10, Type: typelabel.Int16, ScaleFactor: "VAR_SF", Units: "var"},
					smdx.PointElement{Id: VAR_SF, Offset: 11, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: PF, Offset: 12, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "Pct"},
					smdx.PointElement{Id: PF_SF, Offset: 13, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: TotWhExp, Offset: 14, Type: typelabel.Acc32, ScaleFactor: "TotWh_SF", Units: "Wh", Mandatory: true},
					smdx.PointElement{Id: TotWhImp, Offset: 16, Type: typelabel.Acc32, ScaleFactor: "TotWh_SF", Units: "Wh", Mandatory: true},
					smdx.PointElement{Id: TotWh_SF, Offset: 18, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: TotVAhExp, Offset: 19, Type: typelabel.Acc32, ScaleFactor: "TotVAh_SF", Units: "VAh"},
					smdx.PointElement{Id: TotVAhImp, Offset: 21, Type: typelabel.Acc32, ScaleFactor: "TotVAh_SF", Units: "VAh"},
					smdx.PointElement{Id: TotVAh_SF, Offset: 23, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: TotVArhImpQ1, Offset: 24, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhImpQ2, Offset: 26, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhExpQ3, Offset: 28, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhExpQ4, Offset: 30, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArh_SF, Offset: 32, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: Evt, Offset: 33, Type: typelabel.Bitfield32, Mandatory: true},
					smdx.PointElement{Id: Rsrvd, Offset: 35, Type: typelabel.Pad, Access: "r", Mandatory: true},
					smdx.PointElement{Id: Ts, Offset: 36, Type: typelabel.Uint32, Access: "r", Mandatory: true},
					smdx.PointElement{Id: Ms, Offset: 38, Type: typelabel.Uint16, Access: "r", Mandatory: true},
					smdx.PointElement{Id: Seq, Offset: 39, Type: typelabel.Uint16, Access: "r", Mandatory: true},
					smdx.PointElement{Id: Alg, Offset: 40, Type: typelabel.Enum16, Access: "r", Mandatory: true},
					smdx.PointElement{Id: N, Offset: 41, Type: typelabel.Uint16, Access: "r", Mandatory: true},
				},
			},
			smdx.BlockElement{
				Length: 1,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: DS, Offset: 0, Type: typelabel.Uint16, Access: "r", Mandatory: true},
				},
			},
		}})
}
