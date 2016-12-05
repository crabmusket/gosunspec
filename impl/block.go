package impl

import (
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

// Plan expands and re-orders the specified set of points, if any, then returns a
// slice that is ordered in the order that they should be applied to the
// model.
//
// If no points are specified, then the set is expanded to all points in the
// block. If any points with scale factors are specified, then the scale
// factors are included if they are not already included. The set of points
// to be read is also extended to re-read any currently valid point
// that references any scale factor that is included in the list of points
// to be read (so that these points do not "unexpectedly" become invalid.)
//
// Then the points are sorted so scale factors appear first in the list,
// then non-scale factors and then within each section, the points
// are sorted in offset order.
func (b *block) Plan(pointIds ...string) ([]spi.PointSPI, error) {
	points := []spi.PointSPI{}

	if len(pointIds) == 0 {
		// if there are no specified points, include all points

		for _, p := range b.points {
			points = append(points, p)
		}
	} else {
		included := map[string]bool{}
		included_sf := map[string]bool{}

		// include all specified points
		for _, id := range pointIds {
			if p, ok := b.points[id]; !ok {
				return nil, sunspec.ErrNoSuchPoint
			} else {
				if !included[id] {
					points = append(points, p)
					included[id] = true
				}
				if p.Type() == typelabel.ScaleFactor {
					included_sf[id] = true
				}
			}
		}

		// include their scale factors too...
		//
		// we do this for several reasons:
		//    - to interpret a point that uses a scale factor, we need the scale factor too
		//    - if we don't there we may read a value point after its scale factor point has changed
		//      By forcing contemporaneous reads of a scale factor and its related points we help to ensure
		//      that the two values are consistent.
		//    - we want to avoid app programmers having to encode knowedlege in their programs
		//      about these depednencies - the knowledge is in the SMDX documents, so lets use it
		for _, p := range points {
			sfp := p.(*point).scaleFactor
			if sfp != nil {
				if !included[sfp.Id()] {
					points = append(points, sfp.(spi.PointSPI))
					included[sfp.Id()] = true
					included_sf[sfp.Id()] = true
				}
			}
		}

		// We also include all the currently valid points that reference any scale
		// factor points we are going to read since we don't want such points to
		// unexpectedly enter an error state when they are invalidated by the
		// read of the scale factor point. This allows twp separate reads each
		// of which have a point that reference a shared scale factor point to
		// be equivalent to a single read of all points or to two reads in which
		// all points related to a single scale factor are read in the same read
		// as the scale factor itself.
		//
		// One consequence of this behaviour is that any local changes (via a
		// setter) to a point dependent on a scale factor point may be lost by a
		// read of any point that is dependent on the same scale factor which
		// itself means that local changes to points should be written to the
		// physical device with Block.Write before the next Block.Read or else
		// they may be lost under some circumstances even if the point concerned
		// is not directly referened by the Read call.
		//
		// Part of the reason we do this is to maximise the consistency of data
		// exposed by the API while minimising both the effort for the programmer
		// to maintain the consistency and also surprising behaviour.
		for _, p := range b.points {
			if sfp := p.scaleFactor; sfp == nil || p.Error() != nil || !included_sf[sfp.Id()] {
				continue
			} else {
				if !included[p.Id()] {
					points = append(points, p)
					included[p.Id()] = true
				}
			}
		}
	}

	// sort so scale factors come first, then other points in offset order
	sort.Sort(scaleFactorFirstOrder(points))
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
