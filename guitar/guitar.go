package guitar

import (
	"container/list"
)

type Guitar struct {
	serialNumber string
	builder      string
	model        string
	_type        string
	backWood     string
	topWood      string
	price        float32
}

func NewGuitar(serialNumber, builder, model, _type, backWood, topWood string, price float32) Guitar {
	return Guitar{
		serialNumber: serialNumber,
		builder:      builder,
		model:        model,
		_type:        _type,
		backWood:     backWood,
		topWood:      topWood,
		price:        price}
}

func (g Guitar) GetSerialNumber() string {
	return g.serialNumber
}

func (g Guitar) setPrice(price float32) {
	g.price = price
}

func (g Guitar) getPrice() float32 {
	return g.price
}

func (g Guitar) getBuilder() string {
	return g.builder
}

func (g Guitar) getModel() string {
	return g.model
}

func (g Guitar) getType() string {
	return g._type
}

func (g Guitar) getBackWood() string {
	return g.backWood
}

func (g Guitar) getTopWood() string {
	return g.topWood
}

type Inventory struct {
	list *list.List
}

func NewInventory() *Inventory {
	l := list.New()
	return &Inventory{list: l}
}

func (i Inventory) AddGuitar(guitar Guitar) {
	i.list.PushBack(guitar)
}

func (i Inventory) GetGuitar(serialNumber string) (Guitar, bool) {
	l := i.list
	for e := l.Front(); e != nil; e = e.Next() {
		guitar := e.Value.(Guitar)
		if guitar.GetSerialNumber() == serialNumber {
			return guitar, true
		}
	}
	return Guitar{}, false
}
