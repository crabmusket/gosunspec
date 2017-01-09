package layout

import (
	"github.com/crabmusket/gosunspec"
)

// RawModel layout specifies the model ID used to describe a block of memory at an Address
type RawModelLayout struct {
	Address uint16          `json:"address"`
	ModelId sunspec.ModelId `json:"modelId"`
}

// RawDeviceLayout is a slice of RawModelLayouts
type RawDeviceLayout struct {
	Models []RawModelLayout `json:"models"`
}

// RawLayout is the type of layout used for non-SunSpec address spaces. This means
// arbitrary Modbus address spaces where blocks are located at arbitrary addresses
// in an address space and neither the model ID nor the block length are encoded
// in the address space itself.
//
// The intent of RawLayout is to allow the sunspec API to be used effectively
// with non-SunSpec address spaces, assuming the work has been done to
// markup the address space with SMDX models and a JSON document that maps
// addresses to models.
type RawLayout struct {
	Devices []RawDeviceLayout `json:"devices"`
}
