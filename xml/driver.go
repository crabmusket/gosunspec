package xml

import (
	"errors"
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/impl"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/spi"
	"strconv"
)

type modelAnchor struct {
	model *ModelElement
}

type blockAnchor struct {
	model *ModelElement
	index int
}

type pointAnchor struct {
	position int // position of the point in the model, -1 if it is not present
}

type xmlPhysical struct {
}

var ErrNoSuchElement = errors.New("no such element")
var ErrEmptyValue = errors.New("empty value")
var ErrTooManyErrors = errors.New("too many errors")

// Open open a new sunspec.Array which is populated from the specified
// DataElement.
func Open(e *DataElement) (sunspec.Array, error) {
	arr := impl.NewArray()
	for x, _ := range e.Devices {
		if dev, err := OpenDevice(e.Devices[x]); err != nil {
			return nil, err
		} else if err := arr.AddDevice(dev.(spi.DeviceSPI)); err != nil {
			return nil, err
		}
	}
	return arr, nil
}

// OpenDevice opens a new sunspec.Device which is populated from the specified
// DeviceElement.
func OpenDevice(dx *DeviceElement) (sunspec.Device, error) {
	d := impl.NewDevice()
	xp := &xmlPhysical{}
	// iterate through the model elements, creating one new model for each
	// then create a block for each index and add points for each point element
	for _, mx := range dx.Models {
		smdx := smdx.GetModel(uint16(mx.Id))
		if smdx == nil {
			continue
		}
		max := uint32(0)
		for _, px := range mx.Points {
			if px.Index > max {
				max = px.Index
			}
		}
		m := impl.NewModel(smdx, int(max), xp)
		if err := d.AddModel(m); err != nil {
			return nil, err
		} else {
			for pi, px := range mx.Points {
				bi := int(px.Index)
				if bi > 0 && len(smdx.Blocks) == 1 {
					bi = bi - 1
				}
				b := m.MustBlock(bi)
				if p, err := b.Point(px.Id); err != nil {
					return nil, err
				} else {
					p.(spi.PointSPI).SetAnchor(&pointAnchor{position: pi})
				}
			}
			bi := 0
			m.DoWithSPI(func(b spi.BlockSPI) {
				b.SetAnchor(&blockAnchor{model: mx, index: bi})
				b.DoWithSPI(func(p spi.PointSPI) {
					if p.Anchor() == nil {
						p.SetAnchor(&pointAnchor{position: -1})
					}
				})
				bi++
			})
		}
	}
	return d, nil
}

// CopyArray copies an existing SunSpec Array into a new SunSpec Array and an
// XML DataElement. Operations on the returned SunSpec Array edit the
// returned DataElement.
func CopyArray(a sunspec.Array) (sunspec.Array, *DataElement) {
	arr := impl.NewArray()
	x := &DataElement{
		Version: "1",
	}

	devices := []spi.DeviceSPI{}
	a.Do(func(d sunspec.Device) {
		da, dx := CopyDevice(d)
		x.Devices = append(x.Devices, dx)
		devices = append(devices, da.(spi.DeviceSPI))
	})
	for _, c := range devices {
		if err := arr.AddDevice(c.(spi.DeviceSPI)); err != nil {
			// we are not expecting this to happen
			panic(err)
		}
	}
	return arr, x
}

// CopyDevice copies an existing SunSpec Device into a new SunSpec Device and an
// XML DeviceElement. Operations on the returned SunSpec Device edit the
// returned DeviceElement.
func CopyDevice(d sunspec.Device) (sunspec.Device, *DeviceElement) {
	dc := impl.NewDevice()
	dx := &DeviceElement{}
	xp := &xmlPhysical{}

	modelAnchors := map[sunspec.ModelId]*modelAnchor{}

	newModelAnchor := func(id sunspec.ModelId) *modelAnchor {

		// we only use labels when there is ambiguity

		var ma *modelAnchor
		var ok bool
		if ma, ok = modelAnchors[id]; ok {
			if ma.model.Index == 0 {
				ma.model.Index = 1
			}
			ma = &modelAnchor{model: &ModelElement{Id: id, Index: ma.model.Index + 1}}
		} else {
			ma = &modelAnchor{model: &ModelElement{Id: id, Index: 0}}
		}
		modelAnchors[id] = ma
		return ma
	}

	d.Do(func(m sunspec.Model) {
		smdx := smdx.GetModel(uint16(m.Id()))

		mc := impl.NewModel(smdx, m.Blocks(), xp)
		ma := newModelAnchor(m.Id())
		mc.SetAnchor(ma)

		dx.Models = append(dx.Models, ma.model)

		repeatOnly := m.Blocks() > 1 && len(smdx.Blocks) == 1

		for i := 0; i < m.Blocks(); i++ {
			b := m.MustBlock(i).(spi.BlockSPI)
			bc := mc.MustBlock(i).(spi.BlockSPI)

			x := i
			if repeatOnly {
				x = x + 1
			}
			anchor := &blockAnchor{
				model: ma.model,
				index: x,
			}
			bc.SetAnchor(anchor)
			points, _ := bc.Plan()
			for _, pc := range points {
				pa := &pointAnchor{position: -1}
				pc.SetAnchor(pa)
				p := b.MustPoint(pc.Id())
				if p.Error() == nil {
					pc.SetValue(p.Value())
				}
			}

			// write the copied values into the copied block

			if err := bc.Write(); err != nil {
				// not expecting this to happen
				panic(err)
			}
		}
	})
	return dc, dx
}

func (phys *xmlPhysical) Read(b spi.BlockSPI, pointIds ...string) error {
	errCount := 0
	ba := b.Anchor().(*blockAnchor)
	if points, err := b.Plan(pointIds...); err != nil {
		return err
	} else {
		for _, p := range points {
			pa := p.Anchor().(*pointAnchor)
			if pa.position < 0 {
				p.SetError(ErrNoSuchElement)
				errCount++
				continue
			}
			px := ba.model.Points[pa.position]
			if len(px.Value) == 0 {
				p.SetError(ErrEmptyValue)
				errCount++
				continue
			}
			sfp := p.ScaleFactorPoint()
			v := px.Value
			if sfp != nil {
				vsf := sunspec.ScaleFactor(px.ScaleFactor)
				if sfp.Error() != nil {
					sfp.SetScaleFactor(vsf)
				}
				if vi, err := strconv.Atoi(v); err != nil {
					p.SetError(err)
					errCount++
					continue
				} else {
					for vsf > sfp.ScaleFactor() {
						vi = vi * 10
						vsf--
					}
					for vsf < sfp.ScaleFactor() {
						vi = vi / 10
						vsf++
					}
					v = strconv.Itoa(vi)
				}
			}
			if err := p.UnmarshalXML(v); err != nil {
				p.SetError(err)
				errCount++
				continue
			}

			// need to values with scaling factors match the scaling factor or else
			// adjust the value to be consistent with the current scaling factor value
		}
	}
	if errCount > 0 {
		return ErrTooManyErrors
	} else {
		return nil
	}
}

func (phys *xmlPhysical) Write(b spi.BlockSPI, pointIds ...string) error {

	write := map[string]bool{}
	if len(pointIds) == 0 {
		b.Do(func(p sunspec.Point) {
			write[p.Id()] = true
		})
	} else {
		for _, id := range pointIds {
			write[id] = true
		}
	}

	ba := b.Anchor().(*blockAnchor)
	b.DoWithSPI(func(p spi.PointSPI) {
		if !write[p.Id()] {
			return
		}
		pa := p.Anchor().(*pointAnchor)
		var px *PointElement
		if pa.position == -1 && p.Error() == nil {
			// add a new point to the model element
			px = &PointElement{Id: p.Id(), Index: uint32(ba.index)}
			pa.position = len(ba.model.Points)
			ba.model.Points = append(ba.model.Points, px)
		} else if pa.position >= 0 && p.Error() != nil {
			// remove a point from the model element
			swap := b.MustPoint(ba.model.Points[len(ba.model.Points)-1].Id).(spi.PointSPI).Anchor().(*pointAnchor)
			ba.model.Points[pa.position] = ba.model.Points[swap.position]
			swap.position = pa.position
			ba.model.Points = ba.model.Points[0 : len(ba.model.Points)-1]
			px = nil
		} else if pa.position >= 0 {
			// update the existing point
			px = ba.model.Points[pa.position]
		}
		if px != nil {
			px.Value = p.MarshalXML()
			if sfp := p.ScaleFactorPoint(); sfp != nil {
				px.ScaleFactor = int16(sfp.ScaleFactor())
			}
			// we could put the description and units (from the SMDX model) in here, but
			// we probably need a configuration option to say whether this is necessary
			// or not
		}
	})

	// We need a special case here (or, at least, somewhere) to fixup the device element with
	// copies of attributes from the model 1 object and a similar case to derive the model 1
	// object from the device object if there is no model 1 object.

	return nil
}
