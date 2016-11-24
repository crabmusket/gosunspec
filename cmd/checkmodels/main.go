package main

import (
	"fmt"
	_ "github.com/crabmusket/gosunspec/core"
	_ "github.com/crabmusket/gosunspec/core/models"
	"github.com/crabmusket/gosunspec/smdx"
	"log"
	"os"
)

type Type struct {
	Name   string
	Length int
}

var types = []Type{
	Type{"acc16", 1},
	Type{"acc32", 2},
	Type{"acc64", 4},
	Type{"bitfield16", 1},
	Type{"bitfield32", 2},
	Type{"count", 1},
	Type{"enum16", 1},
	Type{"enum32", 2},
	Type{"eui48", 3},
	Type{"float32", 2},
	Type{"int16", 1},
	Type{"int32", 2},
	Type{"int64", 4},
	Type{"ipaddr", 2},
	Type{"ipv6addr", 8},
	Type{"pad", 1},
	Type{"string", 0},
	Type{"sunssf", 1},
	Type{"uint16", 1},
	Type{"uint32", 2},
	Type{"uint64", 4},
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
