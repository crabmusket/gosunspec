package xml

import (
	"bytes"
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/memory"
	"github.com/crabmusket/gosunspec/models/model1"
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

	expected := 432
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

func TestInconsistentXML(t *testing.T) {
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
			if m.Id() != model101.ModelID {
				return
			}
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

const implicitModel1XML = `
<SunSpecData v="1">
	<d man="ACME" mod="Model X" sn="000001">
	</d>
</SunSpecData>
`

func TestImplicitModel1(t *testing.T) {
	data, err := parseXML(bytes.NewBuffer([]byte(implicitModel1XML)))
	if err != nil {
		t.Fatal(err)
	}
	arr, err := Open(data)
	if err != nil {
		t.Fatal(err)
	}
	arr.Do(func(d sunspec.Device) {
		d.MustModel(model1.ModelID)
		d.Do(func(m sunspec.Model) {
			b := m.MustBlock(0)
			if err := b.Read(model1.SN, model1.Md, model1.Mn); err != nil {
				t.Fatal(err)
			}

			// read from implicit model 1 element and see device element contents

			expected := "ACME"
			actual := b.MustPoint(model1.Mn).StringValue()
			if expected != actual {
				t.Fatalf("Mn. actual=%s, expected=%s\n", actual, expected)
			}
			expected = "Model X"
			actual = b.MustPoint(model1.Md).StringValue()
			if expected != actual {
				t.Fatalf("Md. actual=%s, expected=%s\n", actual, expected)
			}
			expected = "000001"
			actual = b.MustPoint(model1.SN).StringValue()
			if expected != actual {
				t.Fatalf("SN. actual=%s, expected=%s\n", actual, expected)
			}

			// write into implicit model 1 element and see results in device element

			b.MustPoint(model1.Mn).SetStringValue("Road Runner")
			b.Write(model1.Mn)
			expected = "Road Runner"
			actual = data.Devices[0].Manufacturer
			if expected != actual {
				t.Fatalf("Man. actual=%s, expected=%s\n", actual, expected)
			}

			b.MustPoint(model1.Md).SetStringValue("Model XX")
			b.Write(model1.Md)
			expected = "Model XX"
			actual = data.Devices[0].Model
			if expected != actual {
				t.Fatalf("Model. actual=%s, expected=%s\n", actual, expected)
			}

			b.MustPoint(model1.SN).SetStringValue("000002")
			b.Write(model1.SN)
			expected = "000002"
			actual = data.Devices[0].Serial
			if expected != actual {
				t.Fatalf("Model. actual=%s, expected=%s\n", actual, expected)
			}

			{
				actual := len(data.Devices[0].Models)
				expected := 0
				if expected != actual {
					t.Fatalf("#models. actual=%d, expected=%d\n", actual, expected)
				}
			}

			b.MustPoint(model1.Opt).SetStringValue("OPT")
			b.Write(model1.Opt)

			{
				actual := len(data.Devices[0].Models)
				expected := 1
				if expected != actual {
					t.Fatalf("#models (after write to opt). actual=%d, expected=%d\n", actual, expected)
				}
			}

		})
	})

}

func TestXmlOpen(t *testing.T) {
	buffer := bytes.NewBuffer([]byte(example))
	data, err := parseXML(buffer)
	if err != nil {
		t.Fatal("could not parse example", err.Error())
	}
	if x, err := Open(data); err != nil {
		t.Fatal(err)
	} else {
		x.Do(func(d sunspec.Device) {
			d.Do(func(m sunspec.Model) {
				m.Do(func(b sunspec.Block) {
					_ = b.Read()
				})
			})
		})
		c, cx := CopyArray(x)
		c.Do(func(d sunspec.Device) {
			_ = d.MustModel(101)
		})
		if cx.Devices[0].Models[1].Id != 101 {
			t.Fatalf("unexpected id found")
		}
		if len(cx.Devices[0].Models[1].Points) != 15 {
			t.Fatalf("not enough points: %d", len(cx.Devices[0].Models[1].Points))
		}
	}
}
