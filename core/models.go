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
}

func (self Model1) GetId() ModelId {
	return 1
}

// Inverter (Single Phase): Include this model for single phase inverter monitoring
type Model101 struct {
	// AC Current
	A float64
	// Phase A Current
	AphA float64
	// Phase B Current
	AphB float64
	// Phase C Current
	AphC float64
	//
	A_SF ScaleFactor
	// Phase Voltage AB
	PPVphAB float64
	// Phase Voltage BC
	PPVphBC float64
	// Phase Voltage CA
	PPVphCA float64
	// Phase Voltage AN
	PhVphA float64
	// Phase Voltage BN
	PhVphB float64
	// Phase Voltage CN
	PhVphC float64
	//
	V_SF ScaleFactor
	// AC Power
	W float64
	//
	W_SF ScaleFactor
	// Line Frequency
	Hz float64
	//
	Hz_SF ScaleFactor
	// AC Apparent Power
	VA float64
	//
	VA_SF ScaleFactor
	// AC Reactive Power
	VAr float64
	//
	VAr_SF ScaleFactor
	// AC Power Factor
	PF float64
	//
	PF_SF ScaleFactor
	// AC Energy
	WH float64
	//
	WH_SF ScaleFactor
	// DC Current
	DCA float64
	//
	DCA_SF ScaleFactor
	// DC Voltage
	DCV float64
	//
	DCV_SF ScaleFactor
	// DC Power
	DCW float64
	//
	DCW_SF ScaleFactor
	// Cabinet Temperature
	TmpCab float64
	// Heat Sink Temperature
	TmpSnk float64
	// Transformer Temperature
	TmpTrns float64
	// Other Temperature
	TmpOt float64
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
