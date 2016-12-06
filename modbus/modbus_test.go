package modbus

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/memory"
	//"github.com/crabmusket/gosunspec/typelabel"
	"testing"
)

func TestModBusSimulator(t *testing.T) {
	if sim, err := OpenSimulator(memory.ComplexNonZeroSlab, 40000); err != nil {
		t.Fatal(err)
	} else {
		if arr, err := Open(sim); err != nil {
			t.Fatal(err)
		} else {
			arr.Do(func(d sunspec.Device) {})
		}

	}
}

// TestComplexSlab iterate over all points and check that they have the expected values.
func TestComplexSlab(t *testing.T) {
	if sim, err := OpenSimulator(memory.ComplexNonZeroSlab, 40000); err != nil {
		t.Fatal(err)
	} else {
		if arr, err := Open(sim); err != nil {
			t.Fatal(err)
		} else {
			arr.Do(func(d sunspec.Device) {
				count := 0
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

				expected := 213
				if count != expected {
					t.Fatalf("unexpected number of points. actual: %d, expected: %d", count, expected)
				}
			})
		}

	}
}

// TestComplexSlab iterate over all points and check that they have the expected values.
func TestComplexSlabStaggered(t *testing.T) {
	if sim, err := OpenSimulator(memory.ComplexNonZeroSlab, 40000); err != nil {
		t.Fatal(err)
	} else {
		if arr, err := Open(sim); err != nil {
			t.Fatal(err)
		} else {
			count := 0
			arr.Do(func(d sunspec.Device) {
				d.Do(func(m sunspec.Model) {
					m.Do(func(b sunspec.Block) {

						// read pairs of adjacaent points

						p := []string{}
						q := []string{}
						c := 0
						b.Do(func(pt sunspec.Point) {
							if c%4 < 2 {
								p = append(p, pt.Id())
							} else {
								q = append(q, pt.Id())
							}
							c++
						})

						if err := b.Read(p...); err != nil {
							t.Fatal(err)
						}

						if err := b.Read(q...); err != nil {
							t.Fatal(err)
						}

						// check all the values

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
			expected := 213
			if count != expected {
				t.Fatalf("unexpected number of points. actual: %d, expected: %d", count, expected)
			}
		}

	}
}
