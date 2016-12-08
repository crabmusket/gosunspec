package xml

import (
	"bytes"
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/memory"
	"github.com/crabmusket/gosunspec/models/model101"
	"testing"
)

func TestCopyDevice(t *testing.T) {
	arr, _ := memory.Open(memory.TwoDeviceSlab)
	arr.Do(func(d sunspec.Device) {
		d.Do(func(m sunspec.Model) {
			m.Do(func(b sunspec.Block) {
				_ = b.Read()
			})
		})
	})
	_, cloneX := CopyArray(arr)
	// bytes, _ := xml.MarshalIndent(cloneX, "", "    ")
	// fmt.Printf("%s\n", string(bytes))
	arr2, err := Open(cloneX)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	arr2.Do(func(d sunspec.Device) {
		d.Do(func(m sunspec.Model) {
			m.Do(func(b sunspec.Block) {
				if err := b.Read(); err != nil {
					t.Fatal(err)
				}
				b.Do(func(p sunspec.Point) {
					if err := p.Error(); err != nil {
						t.Fatalf("p has error. model=%d, point=%s\n", m.Id(), p.Id())
					}

					if v := p.Value(); v != memory.ExpectedValues[p.Type()] {
						t.Fatalf("unexpected value. model=%d, point=%s. actual=%#v, expected=%#v. type=%s", m.Id(), p.Id(), v, memory.ExpectedValues[p.Type()], p.Type())
					}
					count++
				})
			})
		})
	})

	expected := 426
	if count != expected {
		t.Fatalf("unexpected number of points. actual: %d, expected: %d", count, expected)
	}
}

const inconsistentXML = `
<SunSpecData v="1">
	<d>
		<m id="101">
			<p id="A" sf="2">31</p>
		</m>
		<m id="101">
			<p id="A" sf="2">31</p>
			<p id="A_SF">-1</p>
		</m>
	</d>
</SunSpecData>
`

func TestInconsitentXML(t *testing.T) {
	data, err := parseXML(bytes.NewBuffer([]byte(inconsistentXML)))
	if err != nil {
		t.Fatal(err)
	}
	arr, err := Open(data)
	if err != nil {
		t.Fatal(err)
	}
	arr.Do(func(d sunspec.Device) {
		d.Do(func(m sunspec.Model) {
			b := m.MustBlock(0)
			if err := b.Read(model101.A); err != nil {
				// t.Fatal(err)
			}
			p := b.MustPoint("A")
			if p.Error() != nil {
				t.Fatal(p.Error())
			}
			v := p.ScaledValue()
			expected := 3100.0
			if v != expected {
				t.Fatalf("actual: %0f, expected: %0f", v, expected)
			}
		})
	})
}
