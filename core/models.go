package core

// Common: All SunSpec compliant devices must include this as the first model
type Model1 struct {
	// Well known value registered with SunSpec for compliance
	Mn String
	// Manufacturer specific value (32 chars)
	Md String
	// Manufacturer specific value (16 chars)
	Opt String
	// Manufacturer specific value (16 chars)
	Vr String
	// Manufacturer specific value (32 chars)
	SN String
	// Modbus device address
	DA uint16
	// Force even alignment
	Pad Pad
}

func (self Model1) GetId() ModelId {
	return 1
}

// Inverter (Single Phase): Include this model for single phase inverter monitoring
type Model101 struct {
	// AC Current
	A uint16
	// Phase A Current
	AphA uint16
	// Phase B Current
	AphB uint16
	// Phase C Current
	AphC uint16
	//
	A_SF ScaleFactor
	// Phase Voltage AB
	PPVphAB uint16
	// Phase Voltage BC
	PPVphBC uint16
	// Phase Voltage CA
	PPVphCA uint16
	// Phase Voltage AN
	PhVphA uint16
	// Phase Voltage BN
	PhVphB uint16
	// Phase Voltage CN
	PhVphC uint16
	//
	V_SF ScaleFactor
	// AC Power
	W int16
	//
	W_SF ScaleFactor
	// Line Frequency
	Hz uint16
	//
	Hz_SF ScaleFactor
	// AC Apparent Power
	VA int16
	//
	VA_SF ScaleFactor
	// AC Reactive Power
	VAr int16
	//
	VAr_SF ScaleFactor
	// AC Power Factor
	PF int16
	//
	PF_SF ScaleFactor
	// AC Energy
	WH Acc32
	//
	WH_SF ScaleFactor
	// DC Current
	DCA uint16
	//
	DCA_SF ScaleFactor
	// DC Voltage
	DCV uint16
	//
	DCV_SF ScaleFactor
	// DC Power
	DCW int16
	//
	DCW_SF ScaleFactor
	// Cabinet Temperature
	TmpCab int16
	// Heat Sink Temperature
	TmpSnk int16
	// Transformer Temperature
	TmpTrns int16
	// Other Temperature
	TmpOt int16
	//
	Tmp_SF ScaleFactor
	// Enumerated value.  Operating state
	St Enum16
	// Vendor specific operating state code
	StVnd Enum16
	// Bitmask value. Event fields
	Evt1 Bitfield32
	// Reserved for future use
	Evt2 Bitfield32
	// Vendor defined events
	EvtVnd1 Bitfield32
	// Vendor defined events
	EvtVnd2 Bitfield32
	// Vendor defined events
	EvtVnd3 Bitfield32
	// Vendor defined events
	EvtVnd4 Bitfield32
}

func (self Model101) GetId() ModelId {
	return 101
}
