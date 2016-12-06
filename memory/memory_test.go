package memory

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/models/model1"
	"github.com/crabmusket/gosunspec/models/model101"
	"github.com/crabmusket/gosunspec/models/model304"
	"github.com/crabmusket/gosunspec/models/model502"
	"testing"
)

// TestSlab tests that we can build a slab, re-open it, write to it, re-open it
// and read the data that we wrote to it back. Also check that implicit reads of
// scaling factors occur as expected and that points are invalidted if their
// related scaling factor is updated.
func TestSlab(t *testing.T) {

	// build a slab...

	bytes, err := NewSlabBuilder().
		AddModel(model101.ModelID).
		Build()
	if err != nil {
		t.Fatal(err)
	}

	// open it for writing

	{
		d, err := Open(bytes)
		if err != nil {
			t.Fatal(err)
		}
		// write a scale factor and a value point

		m := d.MustModel(model101.ModelID)
		k := m.MustBlock(0)
		k.MustPoint(model101.A_SF).SetScaleFactor(1)
		k.MustPoint(model101.A).SetUint16(10)

		err = k.Write(model101.A, model101.A_SF)
		if err != nil {
			t.Fatal(err)
		}

		{
			// write a string point

			m := d.MustModel(model1.ModelID)
			b := m.MustBlock(0)
			b.MustPoint(model1.SN).SetStringValue("abcde")
			b.Write(model1.SN)
		}

	}

	// reopen the slab for reading

	{
		d, err := Open(bytes)
		if err != nil {
			t.Fatal(err)
		}

		m := d.MustModel(model101.ModelID)
		k := m.MustBlock(0)

		// read the value point with an implied read of the relataed scale factor

		err = k.Read(model101.A)
		if err != nil {
			t.Fatal(err)
		}

		p := k.MustPoint(model101.A)
		sf := k.MustPoint(model101.A_SF)

		// check that the sf was actually read

		if err := sf.Error(); err != nil {
			t.Fatal(sf.Error())
		}

		// check that the value is what we wrote.

		v := p.Uint16()
		if v != 10 {
			t.Fatalf("actual=%d, expected=%d", v, 10)
		}

		// check that scaling works as expected

		v2 := p.ScaledValue()
		if v2 != 100.0 {
			t.Fatalf("actual=%f, expected=%f", v2, 100.0)
		}

		// check that adjusting the scale factor
		// invalidates the related value point

		sf.SetScaleFactor(sunspec.ScaleFactor(2))
		if p.Error() == nil {
			t.Fatalf("actual: no error, expected: error")
		}

		// restore the value point (and scaling factor)
		// from the underlying model

		if err := k.Read(model101.A); err != nil {
			t.Fatal(err)
		}

		// check that scaling works as expected with
		// the original values

		v3 := p.ScaledValue()
		if v3 != v2 {
			t.Fatalf("actual: %f, expected: %f", v3, v2)
		}

		{
			// check that the string we wrote into
			// model1 survived the round trip.

			m := d.MustModel(model1.ModelID)
			b := m.MustBlock(0)
			if err := b.Read(model1.SN); err != nil {
				t.Fatal(err)
			}
			s := b.MustPoint(model1.SN).StringValue()
			if "abcde" != s {
				t.Fatalf("actual: %v (%d), expected: %s (%d)", s, len(s), "abcde", len("abcde"))
			}
		}
	}
}

// TestComplexSlab iterate over all points and check that they have the expected values.
func TestComplexSlab(t *testing.T) {
	d, _ := Open(ComplexNonZeroSlab)

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

				if v := p.Value(); v != ExpectedValues[p.Type()] {
					t.Fatalf("unexpected value. model=%d, point=%s. actual=%#v, expected=%#v. type=%s", m.Id(), p.Id(), v, ExpectedValues[p.Type()], p.Type())
				}
				count++
			})
		})
	})

	{
		actual := d.MustModel(model304.ModelID).Blocks()
		expected := 2
		if actual != expected {
			t.Fatalf("wrong number of blocks. actual: %d, expected: %d", actual, expected)
		}
	}

	expected := 213
	if count != expected {
		t.Fatalf("unexpected number of points. actual: %d, expected: %d", count, expected)
	}
}

func TestMustModelFailsWithManyModels(t *testing.T) {

	d, _ := Open(ComplexNonZeroSlab)

	err := func() (err error) {
		defer func() {
			r := recover()
			if e, ok := r.(error); ok {
				err = e
			}
			panic(r)
		}()
		d.MustModel(model502.ModelID)
		return nil
	}
	if err == nil {
		t.Fatalf("error expected")
	}
}
