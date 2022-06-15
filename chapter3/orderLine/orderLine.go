package orderLine

import (
	"go_training/chapter3/base"
)

type Entity struct {
	base.Entity
	Item      string
	Quantity  int
	UnitPrice float32
}

func New(id int, item string, quantity int, unitPrice float32) Entity {
	return Entity{
		Entity:    base.New(id),
		Item:      item,
		Quantity:  quantity,
		UnitPrice: unitPrice,
	}
}
