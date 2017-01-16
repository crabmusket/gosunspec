package layout_test

import (
	"bytes"
	"encoding/hex"
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/layout"
	"github.com/crabmusket/gosunspec/memory"
	"github.com/crabmusket/gosunspec/spi"
	"log"
	"sync"
	"testing"
)

var xml = `
<layout>
	<device>
		<model id="1"     addr="4" />
		<model id="11"    addr="72" /> 
		<model id="101"   addr="87" />
		<model id="304"   addr="139" repeats="2"/>
		<model id="401"   addr="159" repeats="2"/>
		<model id="501"   addr="191" />
		<model id="502"   addr="224" />
		<model id="502"   addr="254" />
		<model id="63001" addr="284" repeats="1"/>
	</device>
</layout>
`

func TestRawLayoutReading(t *testing.T) {
	blank := make([]byte, len(memory.ComplexEmptySlab))
	copy(blank, memory.ComplexEmptySlab)

	defer func() {
		if e := recover(); e != nil {
			log.Printf("dump:\n %+v \n %+v", hex.Dump(blank), hex.Dump(memory.ComplexNonZeroSlab))
			panic(e)
		}
	}()

	mem, err := memory.Open(blank)
	if err != nil {
		t.Fatal(err)
	}
	layout, err := layout.FromLayoutXML(bytes.NewReader([]byte(xml)))
	if err != nil {
		t.Fatal(err)
	}
	raw, err := memory.OpenWithLayout(memory.ComplexNonZeroSlab, layout)
	if err != nil {
		t.Fatal(err)
	}

	chD := make(chan sunspec.Device)
	chM := make(chan sunspec.Model)
	chB := make(chan sunspec.Block)
	chP := make(chan sunspec.Point)

	wg := sync.WaitGroup{}
	wg.Add(1)

	test304 := func(n string) func(d sunspec.Device) {
		return func(d sunspec.Device) {
			m := d.MustModel(304)
			{
				expected := uint16(m.(spi.ModelSPI).Blocks() * 6)
				actual := uint16(m.(spi.ModelSPI).Length())
				if actual != expected {
					t.Fatalf("model 304 (%s) has inconsistent model lemgth. actual=%d, expected=%d\n", n, actual, expected)
				}

			}
			{
				expected := uint16(3)
				actual := uint16(m.Blocks())
				if actual != expected {
					t.Fatalf("model 304 (%s) has wrong number of blocks. actual=%d, expected=%d\n", n, actual, expected)
				}
			}
		}
	}
	test401 := func(n string) func(d sunspec.Device) {
		return func(d sunspec.Device) {
			m := d.MustModel(401)
			{
				expected := uint16(3)
				actual := uint16(m.Blocks())
				if actual != expected {
					t.Fatalf("model 401 (%s) has wrong number of blocks. actual=%d, expected=%d\n", n, actual, expected)
				}
			}
			{
				expected := uint16((m.(spi.ModelSPI).Blocks()-1)*8 + 14)
				actual := uint16(m.(spi.ModelSPI).Length())
				if actual != expected {
					t.Fatalf("model 401 (%s) has inconsistent model lemgth. actual=%d, expected=%d\n", n, actual, expected)
				}
			}
		}
	}

	mem.Do(test304("mem"))
	raw.Do(test304("raw"))
	mem.Do(test401("mem"))
	raw.Do(test401("raw"))

	go func() {
		mem.Do(func(d sunspec.Device) {
			d1 := <-chD
			if d1 != nil {
				d.Do(func(m sunspec.Model) {
					_ = <-chM
					bn := 0
					m.Do(func(b sunspec.Block) {
						b1 := <-chB
						_ = b1
						b.DoScaleFactorsFirst(func(p sunspec.Point) {
							p1 := <-chP
							if p1.Error() == nil {
								p.SetValue(p1.Value())
							}
						})
						bn++
						b.Write()
					})
				})
			}
		})
		wg.Done()
	}()

	raw.Do(func(d sunspec.Device) {
		chD <- d
		d.Do(func(m sunspec.Model) {
			chM <- m
			bn := 0
			m.Do(func(b sunspec.Block) {
				_ = b.Read()
				chB <- b
				b.DoScaleFactorsFirst(func(p sunspec.Point) {
					chP <- p
				})
				bn++
			})
		})
	})
	close(chD)

	wg.Wait()

	for i, b := range blank {
		if b != memory.ComplexNonZeroSlab[i] {
			t.Fatalf("mismatch at position 0x%04x:\n %+v \n %+v", i, hex.Dump(blank), hex.Dump(memory.ComplexNonZeroSlab))
		}
	}

}
