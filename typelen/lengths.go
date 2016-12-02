package typelen

import (
	"github.com/crabmusket/gosunspec/typelabel"
)

const (
	Acc16       = 1
	Acc32       = 2
	Acc64       = 4
	Bitfield16  = 1
	Bitfield32  = 2
	Count       = 1
	Enum16      = 1
	Enum32      = 2
	Eui48       = 3
	Float32     = 2
	Int16       = 1
	Int32       = 2
	Int64       = 4
	Ipaddr      = 2
	Ipv6addr    = 8
	Pad         = 1
	String      = 0
	ScaleFactor = 1
	Uint16      = 1
	Uint32      = 2
	Uint64      = 4
)

type descriptor struct {
	Name   string
	Length uint16
}

var descriptors = []descriptor{
	descriptor{typelabel.Acc16, Acc16},
	descriptor{typelabel.Acc32, Acc32},
	descriptor{typelabel.Acc64, Acc64},
	descriptor{typelabel.Bitfield16, Bitfield16},
	descriptor{typelabel.Bitfield32, Bitfield32},
	descriptor{typelabel.Count, Count},
	descriptor{typelabel.Enum16, Enum16},
	descriptor{typelabel.Enum32, Enum32},
	descriptor{typelabel.Eui48, Eui48},
	descriptor{typelabel.Float32, Float32},
	descriptor{typelabel.Int16, Int16},
	descriptor{typelabel.Int32, Int32},
	descriptor{typelabel.Int64, Int64},
	descriptor{typelabel.Ipaddr, Ipaddr},
	descriptor{typelabel.Ipv6addr, Ipv6addr},
	descriptor{typelabel.Pad, Pad},
	descriptor{typelabel.String, String},
	descriptor{typelabel.ScaleFactor, ScaleFactor},
	descriptor{typelabel.Uint16, Uint16},
	descriptor{typelabel.Uint32, Uint32},
	descriptor{typelabel.Uint64, Uint64},
}

var descriptorMap = map[string]descriptor{}

func init() {
	for _, v := range descriptors {
		descriptorMap[v.Name] = v
	}
}

func Length(n string) uint16 {
	return descriptorMap[n].Length
}
