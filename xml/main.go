// xml implements data structures and functions for parsing, working with and
// exporting SunSpec data in XML exchange format. The XML schema in use is
// described in the SunSpec Data Exchange Specification version 1.2,
// particularly on page 11. Note that the parser in this package is case-sensitive
// even though some examples in the specification use `sunSpecData` instead of
// `SunSpecData` as the top-level tag name (this package assumes the latter).
// http://sunspec.org/wp-content/uploads/2015/06/SunSpec-Model-Data-Exchange-12021.pdf
package xml

import (
	"encoding/xml"
	sunspec "github.com/crabmusket/gosunspec/core"
	"io"
	"math"
	"strconv"
	"time"
)

type DataElement struct {
	XMLName xml.Name        `xml:"SunSpecData"`
	Version string          `xml:"v,attr"`
	Devices []DeviceElement `xml:"d"`
}

type DeviceElement struct {
	XMLName       xml.Name       `xml:"d"`
	CorrelationId uint32         `xml:"cid,attr,omitempty"`
	Id            string         `xml:"id,attr,omitempty"`
	Namespace     string         `xml:"ns,attr"`
	LoggerId      string         `xml:"lid,attr,omitempty"`
	Manufacturer  string         `xml:"man,attr"`
	Model         string         `xml:"mod,attr"`
	Serial        string         `xml:"sn,attr"`
	Timestamp     time.Time      `xml:"t,attr"`
	Models        []ModelElement `xml:"m"`
}

// Note that we can use omitempty on Index because indices in SunSpec XML start
// at 1. Therefore an index of 0 (unused) will not be serialised.
type ModelElement struct {
	XMLName   xml.Name        `xml:"m"`
	Id        sunspec.ModelId `xml:"id,attr"`
	Namespace string          `xml:"ns,attr,omitempty"`
	Index     uint32          `xml:"x,attr,omitempty"`
	Points    []PointElement  `xml:"p"`
}

// Note that we can use omitempty on ScaleFactor because a scale factor of 0
// means no scaling. Therefore an sf of 0 is meaningless anyway.
type PointElement struct {
	XMLName     xml.Name  `xml:"p"`
	Description string    `xml:"d,attr,omitempty"`
	Id          string    `xml:"id,attr"`
	ScaleFactor int16     `xml:"sf,attr,omitempty"`
	Timestamp   time.Time `xml:"t,attr,omitempty"`
	Unit        string    `xml:"u,attr,omitempty"`
	Index       uint32    `xml:"x,attr,omitempty"`
	Value       string    `xml:",chardata"`
}

func parseXML(reader io.Reader) (elements []DataElement, err error) {
	decoder := xml.NewDecoder(reader)
	err = decoder.Decode(&elements)
	return
}

func LoadDevices(reader io.Reader) (devices []*sunspec.Device, err error) {
	dataElements, err := parseXML(reader)
	if err != nil {
		return
	}

	for _, dataElement := range dataElements {
		for _, deviceElement := range dataElement.Devices {
			if loaded := deviceFromElement(deviceElement); loaded != nil {
				devices = append(devices, loaded)
			}
		}
	}

	return
}

func deviceFromElement(deviceElement DeviceElement) *sunspec.Device {
	device := sunspec.Device{
		Models: []sunspec.Model{},
	}

	for _, modelElement := range deviceElement.Models {
		if model := modelFromElement(modelElement); model != nil {
			device.Models = append(device.Models, model)
		}
	}

	if len(device.Models) > 0 {
		return &device
	} else {
		return nil
	}
}

func (self *ModelElement) GetPointValueString(id string) string {
	rawVal := ""
	for _, point := range self.Points {
		if point.Id == id {
			rawVal = point.Value
			break
		}
	}
	return rawVal
}

func (self *ModelElement) GetPointValueUint16(id string) uint16 {
	rawVal := ""
	for _, point := range self.Points {
		if point.Id == id {
			rawVal = point.Value
			break
		}
	}
	if rawVal == "" {
		return 0
	}

	val, err := strconv.Atoi(rawVal)
	if err != nil {
		return 0
	} else {
		return uint16(val)
	}
}

func (self *ModelElement) GetPointScaleFactor(id string, sfid string) float64 {
	generalScaleFactor := int16(0)
	pointScaleFactor := int16(0)
	foundPointScaleFactor := false
	for _, point := range self.Points {
		if point.Id == sfid {
			generalScaleFactor = point.ScaleFactor
		}
		if point.Id == id {
			pointScaleFactor = point.ScaleFactor
			foundPointScaleFactor = true
			break
		}
	}

	if foundPointScaleFactor {
		return float64(math.Pow(10, float64(pointScaleFactor)))
	} else {
		return float64(math.Pow(10, float64(generalScaleFactor)))
	}
}
