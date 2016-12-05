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
	r := []sunspec.Model{}
	d.Do(func(m sunspec.Model) {
		if m.Id() == id {
			r = append(r, m)
		}
	})
	if len(r) < 1 {
		panic(sunspec.ErrNoSuchModel)
	} else if len(r) > 1 {
		panic(sunspec.ErrTooManyModels)
	}
	return r[0]
}

func (d *device) Do(f func(b sunspec.Model)) {
	for _, b := range d.models {
		f(b)
	}
}

func (d *device) DoWithSPI(f func(b spi.ModelSPI)) {
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
