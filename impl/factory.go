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

// Answer a new model that can be mapped to a contiguous range of bytes. The size
// (or bound) of that range determines the number of repeats.
func NewContiguousModel(me *smdx.ModelElement, bound uint16, phys spi.Physical) spi.ModelSPI {
	reps := 0
	trunc := false
	if len(me.Blocks) > 1 {
		reps = (int(bound) - int(me.Blocks[0].Length)) / int(me.Blocks[1].Length)
		if reps < 0 {
			reps = 0
		}
	} else {
		if me.Blocks[0].Length > bound {
			// see memory.go for explanation of this condition
			trunc = true
		}
	}
	m := NewModel(me, reps, phys)
	if trunc {
		if b, err := m.Block(0); err != nil {
			b.(spi.BlockSPI).SetLength(bound)
		}
	}
	return m
}

// Answer a new model with nr repeats of the repating block, if any. The caller
// determines the number of repeats.
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
