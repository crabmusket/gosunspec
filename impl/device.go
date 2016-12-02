package impl

import (
	"errors"
	"fmt"
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/spi"
)

var (
	errWrongType      = errors.New("supplied model has wrong implementation type")
	errDuplicateModel = errors.New("duplicate model")
)

type device struct {
	anchored
	models []*model
}

func (d *device) Model(id sunspec.ModelId) (sunspec.Model, error) {
	for _, m := range d.models {
		if m.Id() == id {
			return m, nil
		}
	}
	return nil, fmt.Errorf("invalid model id: ", id)
}

func (d *device) MustModel(id sunspec.ModelId) sunspec.Model {
	if m, err := d.Model(id); err != nil {
		panic(err)
	} else {
		return m
	}
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
		for _, e := range d.models {
			if e.Id() == m.Id() {
				return errDuplicateModel
			}
		}
		d.models = append(d.models, r)
	} else if !ok {
		return errWrongType
	}
	return nil
}
