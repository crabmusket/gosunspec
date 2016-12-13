package impl

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/spi"
)

type array struct {
	anchored
	devices []*device
}

func (a *array) Do(f func(d sunspec.Device)) {
	for _, d := range a.devices {
		f(d)
	}
}

func (a *array) AddDevice(d spi.DeviceSPI) error {
	a.devices = append(a.devices, d.(*device))
	return nil
}
