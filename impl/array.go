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

func (a *array) DoWithSPI(f func(d spi.DeviceSPI)) {
	for _, d := range a.devices {
		f(d)
	}
}

func (a *array) AddDevice(m spi.DeviceSPI) error {
	return nil
}
