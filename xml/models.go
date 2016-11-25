package xml

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/models/model1"
	"github.com/crabmusket/gosunspec/models/model101"
)

func modelFromElement(el ModelElement) sunspec.Model {
	switch el.Id {
	case 1:
		return &model1.Block1{
			Mn:  el.GetPointValueString("Mn"),
			Md:  el.GetPointValueString("Md"),
			Opt: el.GetPointValueString("Opt"),
			Vr:  el.GetPointValueString("Vr"),
			SN:  el.GetPointValueString("SN"),
			DA:  el.GetPointValueUint16("DA"),
		}

	case 101:
		return &model101.Block101{
			A: el.GetPointValueUint16("A"),
			/*
				AphA uint16
				AphB uint16
				AphC uint16
				A_SF ScaleFactor
				PPVphAB uint16
				PPVphBC uint16
				PPVphCA uint16
			*/
			PhVphA: el.GetPointValueUint16("PhVphA"),
			/*
				PhVphB uint16
				PhVphC uint16
				V_SF ScaleFactor
				W int16
				W_SF ScaleFactor
				Hz uint16
				Hz_SF ScaleFactor
				VA int16
				VA_SF ScaleFactor
				VAr int16
				VAr_SF ScaleFactor
				PF int16
				PF_SF ScaleFactor
				WH Acc32
				WH_SF ScaleFactor
				DCA uint16
				DCA_SF ScaleFactor
				DCV uint16
				DCV_SF ScaleFactor
				DCW int16
				DCW_SF ScaleFactor
				TmpCab int16
				TmpSnk int16
				TmpTrns int16
				TmpOt int16
				Tmp_SF ScaleFactor
				St Enum16
				StVnd Enum16
				Evt1 Bitfield32
				Evt2 Bitfield32
				EvtVnd1 Bitfield32
				EvtVnd2 Bitfield32
				EvtVnd3 Bitfield32
				EvtVnd4 Bitfield32
			*/
		}
	}

	return nil
}

// Common: All SunSpec compliant devices must include this as the first model
type Model1 struct {
}
