package impl

import (
	"github.com/crabmusket/gosunspec/spi"
	"testing"
)

func TestCompletePointInterface(t *testing.T) {
	_ = spi.PointSPI((&point{}))
}

func TestCompleteDevice(t *testing.T) {
	_ = spi.DeviceSPI((&device{}))
}

func TestCompleteArray(t *testing.T) {
	_ = spi.ArraySPI((&array{}))
}
