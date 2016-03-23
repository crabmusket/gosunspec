package xml

import (
	"bytes"
	"testing"
)

// Examples from the SunSpec Data Exchange Specification version 1.2.
// http://sunspec.org/wp-content/uploads/2015/06/SunSpec-Model-Data-Exchange-12021.pdf

// Page 10
const example = `
<SunSpecData v="1">
	<d ns="mac" lid="11:22:33:44:55:66" man="gsc" mod="r300" sn="123456" t="2011-05-12T09:21:49Z" cid="2">
		<m id="101" x="1">
			<p id="A">30.43</p>
			<p id="PhVphA" sf="-1">2216</p>
			<p id="W" u="Watts">6701.3</p>
			<p id="Hz">60.01</p>
			<p id="WH">126973</p>
			<p id="DCA">14.28</p>
			<p id="DCV">469</p>
			<p id="DCW">6805</p>
			<p id="TmpOt">32.94</p>
			<p id="St">4</p>
		</m>
	</d>
</SunSpecData>
`

func TestXmlParse(t *testing.T) {
	buffer := bytes.NewBuffer([]byte(example))
	data, err := parseXML(buffer)
	if err != nil {
		t.Fatal("could not parse example", err.Error())
	}

	if len(data) != 1 {
		t.Fatal("wrong number of data packets found")
	}
	if data[0].Version != "1" {
		t.Error("wrong version found")
	}
	if len(data[0].Devices) != 1 {
		t.Fatal("wrong number of devices found")
	}
	if len(data[0].Devices[0].Models) != 1 {
		t.Fatal("wrong number of models found")
	}
	if len(data[0].Devices[0].Models[0].Points) != 10 {
		t.Fatal("wrong number of points found")
	}
	if id := data[0].Devices[0].Models[0].Points[0].Id; id != "A" {
		t.Error("wrong id in first point:", id)
	}
	if value := data[0].Devices[0].Models[0].Points[0].Value; value != "30.43" {
		t.Error("wrong value in first point:", value)
	}
	if scale := data[0].Devices[0].Models[0].Points[1].ScaleFactor; scale != -1 {
		t.Error("wrong scale factor in second point:", scale)
	}
	if units := data[0].Devices[0].Models[0].Points[2].Unit; units != "Watts" {
		t.Error("wrong units in third point:", units)
	}
}
