package impl

import (
	"github.com/crabmusket/gosunspec/spi"
)

// A structure which can be anchored to the physical implementation
type anchored struct {
	anchor spi.Anchor
}

func (a *anchored) SetAnchor(p spi.Anchor) {
	a.anchor = p
}

func (a *anchored) Anchor() spi.Anchor {
	return a.anchor
}
