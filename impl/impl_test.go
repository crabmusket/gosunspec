package impl

import (
	"encoding/binary"
	"testing"

	sunspec "github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/spi"
	"github.com/crabmusket/gosunspec/typelabel"
)

func TestCompletePointInterface(t *testing.T) {
	_ = spi.PointSPI((&point{}))
}

func TestCompleteDevice(t *testing.T) {
	_ = spi.DeviceSPI((&device{}))
}

func TestCompleteArray(t *testing.T) {
	_ = spi.ArraySPI((&array{}))
}

func TestMarshalEui48(t *testing.T) {
	p := point{
		err: nil,
		smdx: &smdx.PointElement{
			Type: typelabel.Eui48,
		},
	}

	p.Unmarshal([]byte{1, 2, 3, 4, 5, 6, 7, 8})

	v := p.Value().(sunspec.Eui48)
	if binary.BigEndian.Uint64(v[:]) != 0x0102030405060708 {
		t.Errorf("unexpected value result, got %v", v)
	}

	if v := p.MarshalXML(); v != "03:04:05:06:07:08" {
		t.Errorf("unexpected xml result, got %v", v)
	}
}
