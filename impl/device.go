package impl

import (
	"errors"
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/spi"
)

var (
	errWrongType = errors.New("supplied model has wrong implementation type")
)

type device struct {
	anchored
	models []*model
}

func (d *device) MustModel(id sunspec.ModelId) sunspec.Model {
	if m, err := sunspec.ExactlyOneModel(d.Collect(sunspec.SameModelId(id))); err == nil {
		return m
	} else {
		panic(err)
	}
}

func (d *device) Do(f func(b sunspec.Model)) {
	for _, b := range d.models {
		f(b)
	}
}

func (d *device) AddModel(m spi.ModelSPI) error {
	if r, ok := m.(*model); ok {
		d.models = append(d.models, r)
	} else if !ok {
		return errWrongType
	}
	return nil
}

func (d *device) Collect(f func(m sunspec.Model) bool) []sunspec.Model {
	r := []sunspec.Model{}
	d.Do(func(m sunspec.Model) {
		if f(m) {
			r = append(r, m)
		}
	})
	return r
}
