package main

import (
	"fmt"
	"github.com/crabmusket/gosunspec/core/typelabel"
	"github.com/crabmusket/gosunspec/core/typelen"
	_ "github.com/crabmusket/gosunspec/models"
	"github.com/crabmusket/gosunspec/smdx"
	"log"
	"os"
)

type Type struct {
	Name   string
	Length int
}

var types = []Type{
	Type{typelabel.Acc16, typelen.Acc16},
	Type{typelabel.Acc32, typelen.Acc32},
	Type{typelabel.Acc64, typelen.Acc64},
	Type{typelabel.Bitfield16, typelen.Bitfield16},
	Type{typelabel.Bitfield32, typelen.Bitfield32},
	Type{typelabel.Count, typelen.Count},
	Type{typelabel.Enum16, typelen.Enum16},
	Type{typelabel.Enum32, typelen.Enum32},
	Type{typelabel.Eui48, typelen.Eui48},
	Type{typelabel.Float32, typelen.Float32},
	Type{typelabel.Int16, typelen.Int16},
	Type{typelabel.Int32, typelen.Int32},
	Type{typelabel.Int64, typelen.Int64},
	Type{typelabel.Ipaddr, typelen.Ipaddr},
	Type{typelabel.Ipv6addr, typelen.Ipv6addr},
	Type{typelabel.Pad, typelen.Pad},
	Type{typelabel.String, typelen.String},
	Type{typelabel.Sunssf, typelen.Sunssf},
	Type{typelabel.Uint16, typelen.Uint16},
	Type{typelabel.Uint32, typelen.Uint32},
	Type{typelabel.Uint64, typelen.Uint64},
}

var typeMap = map[string]Type{}

func init() {
	for _, v := range types {
		typeMap[v.Name] = v
	}
}

func main() {
	if err := smdx.DoModels(func(m *smdx.ModelElement) error {
		if len(m.Blocks) < 1 {
			fmt.Fprintf(os.Stderr, "%d: not enough blocks\n", m.Id)
		} else if len(m.Blocks) > 2 {
			fmt.Fprintf(os.Stderr, "%d: too many blocks\n", m.Id)
		}
		totalBlockLen := 0
		for _, b := range m.Blocks {
			totalBlockLen += int(b.Length)
		}
		if totalBlockLen != int(m.Length) {
			fmt.Fprintf(os.Stderr, "%d: total block length mismatch\n", m.Id)
		}
		for bx, b := range m.Blocks {
			totalPointLen := 0
			for _, p := range b.Points {
				t := typeMap[p.Type]
				if t.Length == 0 {
					if p.Length == 0 {
						fmt.Fprintf(os.Stderr, "%d: %d: %s: zero length point\n", m.Id, bx, p.Id)
					}
					totalPointLen += int(p.Length)
				} else {
					totalPointLen += int(t.Length)
					if p.Length != 0 && int(p.Length) != t.Length {
						fmt.Fprintf(os.Stderr, "%d: %d: %d: inconsistent point length\n", m.Id, bx, p.Id)
					}
				}
			}
			if int(b.Length)-totalPointLen > 1 || int(b.Length) < totalPointLen {
				fmt.Fprintf(os.Stderr, "%d: %d: block length inconsistent with point length: %d, %d\n", m.Id, bx, b.Length, totalPointLen)

			}
		}
		return nil
	}); err != nil {
		log.Fatalf("failed: %v", err)
	}
}
