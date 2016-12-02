package impl

import (
	_ "fmt"
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/spi"
	"github.com/crabmusket/gosunspec/typelabel"
	"sort"
)

type block struct {
	anchored
	physical spi.Physical
	smdx     *smdx.BlockElement
	points   map[string]*point
	length   uint16
}

func (b *block) Point(id string) (sunspec.Point, error) {
	if p := b.points[id]; p != nil {
		return p, nil
	} else {
		return nil, sunspec.ErrNoSuchPoint
	}
}

func (b *block) MustPoint(id string) sunspec.Point {
	if p, err := b.Point(id); err != nil {
		panic(err)
	} else {
		return p
	}
}

func (b *block) Do(f func(p sunspec.Point)) {
	for _, pe := range b.smdx.Points {
		f(b.points[pe.Id])
	}
}

func (b *block) Read(pointIds ...string) error {
	return b.physical.Read(b, pointIds...)
}

func (b *block) Write(pointIds ...string) error {
	return b.physical.Write(b, pointIds...)
}

func (b *block) DoWithSPI(f func(p spi.PointSPI)) {
	for _, pe := range b.smdx.Points {
		f(b.points[pe.Id])
	}
}

func (b *block) Length() uint16 {
	if b.length > 0 {
		return b.length
	} else {
		return b.smdx.Length
	}
}

func (b *block) SetLength(l uint16) {
	b.length = l
}

type ScaleFactorFirstOrder []spi.PointSPI

func (o ScaleFactorFirstOrder) Len() int {
	return len(o)
}

func (o ScaleFactorFirstOrder) Less(i, j int) bool {
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

func (o ScaleFactorFirstOrder) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

// Plan expands and re-orders the specified set of points, if any, then returns a
// slice that is ordered in the order that they should be applied to the
// model.
//
// If no points are specified, then the set is expanded to all points in the
// block. If any points with scale factors are specified, then the scale
// factors are included if they are not already inncluded.
//
// Then the points are sorted so scale factors appear first in the list,
// then non-scale factors and then within each section, the points
// are sorted in offset order.
func (b *block) Plan(pointIds ...string) ([]spi.PointSPI, error) {
	points := []spi.PointSPI{}
	included := map[string]bool{}

	// include all specified points
	for _, id := range pointIds {
		if p, ok := b.points[id]; !ok {
			return nil, sunspec.ErrNoSuchPoint
		} else {
			if !included[id] {
				points = append(points, p)
				included[id] = true
			}
		}
	}

	// include their scale factors too
	for _, p := range points {
		sfp := p.(*point).scaleFactor
		if sfp != nil {
			if !included[sfp.Id()] {
				points = append(points, sfp.(spi.PointSPI))
				included[sfp.Id()] = true
			}
		}
	}

	// if there are no points, include all points
	if len(pointIds) == 0 {
		for _, p := range b.points {
			points = append(points, p)
		}
	}

	// sort so scale factors come first, then other points in offset order
	sort.Sort(ScaleFactorFirstOrder(points))
	return points, nil
}

// Invalidate checks if the specified point is a scale factor. If it is, then
// any other point that specifies this point as its scale factor is invalidated.
func (b *block) Invalidate(p spi.PointSPI) {
	for _, d := range b.points {
		if sfp := d.scaleFactor; sfp != nil {
			if sfp.Id() == p.Id() {
				d.SetValue(nil)
			}
		}
	}
}
