package order

import (
	"go_training/chapter3/base"
	"go_training/chapter3/orderLine"
)

type Entity struct {
	base.Entity
	OrderLines []orderLine.Entity
	TotalCost  float64
	User       string
}

func (order *Entity) AddOrderLine(id int, item string, quantity int, unitPrice float32) {
	orderLine := orderLine.New(id, item, quantity, unitPrice)
	order.OrderLines = append(order.OrderLines, orderLine)
	order.UpdateTotalCost()
}

func (order *Entity) UpdateOrderLine(id int, item string, quantity int, unitPrice float32) {
	for _, orderLine := range order.OrderLines {
		if orderLine.Entity.Id == id {
			orderLine.Item = item
			orderLine.Quantity = quantity
			orderLine.UnitPrice = unitPrice
			break
		}
	}
	order.UpdateTotalCost()
}

func (order *Entity) DeleteOrderLine(id int) {
	for i, orderLine := range order.OrderLines {
		if orderLine.Id == id {
			order.OrderLines = append(order.OrderLines[:i], order.OrderLines[i+1:]...)
			break
		}
	}
	order.UpdateTotalCost()
}

func (order *Entity) UpdateTotalCost() {
	total := 0.0
	for _, orderLine := range order.OrderLines {
		total += float64(orderLine.UnitPrice) * float64(orderLine.Quantity)
	}
	order.TotalCost = total
}

func New(id int, user string) Entity {
	return Entity{
		Entity:    base.New(id),
		TotalCost: 0.0,
		User:      user,
	}
}
