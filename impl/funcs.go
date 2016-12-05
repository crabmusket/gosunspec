package impl

import (
	"github.com/crabmusket/gosunspec/spi"
	"github.com/crabmusket/gosunspec/typelabel"
)

// scaleFactorFirstOrder defines the sort order for points in which scale factors appear before non-scale factors
// and everything is sorted in modbus offset order. Points need to be written into the model in the order
// defined by this order so as to guarantee that the non-scale factor points are not invalidated.
type scaleFactorFirstOrder []spi.PointSPI

func (o scaleFactorFirstOrder) Len() int {
	return len(o)
}

func (o scaleFactorFirstOrder) Less(i, j int) bool {
	p1 := o[i]
	p2 := o[j]
	if p1.Type() == typelabel.ScaleFactor && p2.Type() != typelabel.ScaleFactor {
		return true
	} else if p1.Type() != typelabel.ScaleFactor && p2.Type() == typelabel.ScaleFactor {
		return false
	} else {
		return p1.Offset() < p2.Offset()
	}
}

func (o scaleFactorFirstOrder) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}
