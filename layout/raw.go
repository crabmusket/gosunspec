package layout

import (
	"encoding/xml"
	"github.com/crabmusket/gosunspec"
	"io"
)

// RawModel layout specifies the model ID used to describe a block of memory at an Address
type RawModelLayout struct {
	XMLName xml.Name        `xml:"model"`
	ModelId sunspec.ModelId `xml:"modelId,attr"`
	Address *uint16         `xml:"address,attr,omitempty"`
}

// RawDeviceLayout is a slice of RawModelLayouts
type RawDeviceLayout struct {
	XMLName xml.Name         `xml:"device"`
	Models  []RawModelLayout `xml:"models"`
}

// RawLayout is the type of layout used for non-SunSpec address spaces. This means
// arbitrary Modbus address spaces where blocks are located at arbitrary addresses
// in an address space and neither the model ID nor the block length are encoded
// in the address space itself.
//
// The intent of RawLayout is to allow the sunspec API to be used effectively
// with non-SunSpec address spaces, assuming the work has been done to
// markup the address space with SMDX models and an XML document that maps
// addresses to model ids.
type RawLayout struct {
	XMLName xml.Name          `xml:"layout"`
	Devices []RawDeviceLayout `xml:"devices"`
}

// FromLayoutXML reads a layout description from an XML stream.
func FromLayoutXML(reader io.Reader) (*RawLayout, error) {
	decoder := xml.NewDecoder(reader)
	layout := &RawLayout{}
	return layout, decoder.Decode(layout)
}
