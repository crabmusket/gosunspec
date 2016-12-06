package modbus

import (
	"github.com/crabmusket/gosunspec/spi"
)

// runBuilder breaks a sequence of points into a sequence of runs, where each run
// of points is strictly adjacent
type runBuilder struct {
	runs [][]spi.PointSPI
	run  []spi.PointSPI
}

func newRunBuilder() *runBuilder {
	run := []spi.PointSPI{}
	return &runBuilder{
		run:  run,
		runs: [][]spi.PointSPI{run},
	}
}

func (r *runBuilder) spawn(p spi.PointSPI) {
	r.run = []spi.PointSPI{p}
	r.runs = append(r.runs, r.run)

}

func (r *runBuilder) adjacent(p spi.PointSPI) bool {
	if len(r.run) == 0 {
		return true
	} else {
		last := r.run[len(r.run)-1]
		return last.Offset()+last.Length() == p.Offset()
	}
}

func (r *runBuilder) extend(p spi.PointSPI) {
	r.run = append(r.run, p)
	r.runs[len(r.runs)-1] = r.run
}

func (r *runBuilder) add(p spi.PointSPI) {
	if r.adjacent(p) {
		r.extend(p)
	} else {
		r.spawn(p)
	}
}
