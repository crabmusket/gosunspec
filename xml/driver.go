package xml

import (
	"github.com/crabmusket/gosunspec"
)

// Open open a new sunspec.Array which is populated from the specified
// DataElement.
func Open(e *DataElement) (sunspec.Array, error) {
	panic("not implemented")
}

// OpenDevice opens a new sunspec.Device which is populated from the specified
// DeviceElement.
func OpenDevice(d *DeviceElement) (sunspec.Device, error) {
	panic("not implemented")
}

// CopyArray copies an existing SunSpec Array into a new SunSpec Array and an
// XML DataElement. Operations on the returned SunSpec Array edit the
// returned DataElement.
func CopyArray(a sunspec.Array) (sunspec.Array, *DataElement) {
	panic("not implemented")
}

// CopyDevice copies an existing SunSpec Device into a new SunSpec Device and an
// XML DeviceElement. Operations on the returned SunSpec Device edit the
// returned DeviceElement.
func CopyDevice(d sunspec.Device) (sunspec.Device, *DeviceElement) {
	panic("not implemented")
}
