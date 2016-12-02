package impl

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/spi"
	"github.com/crabmusket/gosunspec/typelabel"
)

func NewArray() spi.ArraySPI {
	return &array{
		devices: []*device{},
	}
}

func NewDevice() spi.DeviceSPI {
	return &device{
		models: []*model{},
	}
}

// Answer a new model instance.
func NewModel(me *smdx.ModelElement, nr int, physical spi.Physical) spi.ModelSPI {

	blocks := []*block{}
	for x, _ := range me.Blocks {
		blocks = append(blocks, newBlock(&me.Blocks[x], physical))
		if nr == 0 {
			break // don't put in the first repeat
		}
	}

	m := &model{
		smdx:     me,
		physical: physical,
		blocks:   blocks,
	}

	for nr > 1 {
		nr--
		m.AddRepeat()
	}

	return m
}

// Answer a new block.
func newBlock(blockSmdx *smdx.BlockElement, physical spi.Physical) *block {
	b := &block{
		smdx:     blockSmdx,
		points:   map[string]*point{},
		physical: physical,
	}

	p := []*smdx.PointElement{}
	q := []*smdx.PointElement{}

	for x, pe := range blockSmdx.Points {
		if pe.Type == typelabel.ScaleFactor {
			p = append(p, &blockSmdx.Points[x])
		} else {
			q = append(q, &blockSmdx.Points[x])
		}
	}

	p = append(p, q...)

	for _, pe := range p {
		var sfp sunspec.Point
		var sp *point
		var ok bool
		if sfp, ok = b.points[pe.ScaleFactor]; ok {
			sp = newPoint(pe, sfp, b)
		} else {
			sp = newPoint(pe, nil, b)
		}
		b.points[pe.Id] = sp
	}

	return b
}

// Answer a new ppint.
func newPoint(smdx *smdx.PointElement, scaleFactor sunspec.Point, b spi.BlockSPI) *point {
	return &point{
		smdx:        smdx,
		scaleFactor: scaleFactor,
		block:       b,
		err:         errNotInitialized,
	}
}
