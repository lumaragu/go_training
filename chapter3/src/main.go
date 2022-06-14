package main

import (
	"fmt"
	"time"
)

type BaseEntity struct {
	id      int
	created time.Time
}

func newBaseEntity(id int) BaseEntity {
	return BaseEntity{id: id, created: time.Now()}
}

type OrderLine struct {
	BaseEntity
	item      string
	quantity  int
	unitPrice float32
}

func newOrderLine(id int, item string, quantity int, unitPrice float32) OrderLine {
	return OrderLine{
		BaseEntity: newBaseEntity(id),
		item:       item,
		quantity:   quantity,
		unitPrice:  unitPrice,
	}
}

type Order struct {
	BaseEntity
	orderLines []OrderLine
	totalCost  float64
	user       string
}

func (o *Order) addOrderLine(id int, item string, quantity int, unitPrice float32) {
	orderLine := newOrderLine(id, item, quantity, unitPrice)
	o.orderLines = append(o.orderLines, orderLine)
	o.updTotalCost()
}

func (o *Order) updOrderLine(id int, item string, quantity int, unitPrice float32) {
	for _, orderLine := range o.orderLines {
		if orderLine.id == id {
			orderLine.item = item
			orderLine.quantity = quantity
			orderLine.unitPrice = unitPrice
			break
		}
	}
	o.updTotalCost()
}

func (o *Order) delOrderLine(id int) {
	for i, orderLine := range o.orderLines {
		if orderLine.id == id {
			o.orderLines = append(o.orderLines[:i], o.orderLines[i+1:]...)
			break
		}
	}
	o.updTotalCost()
}

func (o *Order) updTotalCost() {
	total := 0.0
	for _, orderLine := range o.orderLines {
		total += float64(orderLine.unitPrice) * float64(orderLine.quantity)
	}
	o.totalCost = total
}

func newOrder(id int, user string) Order {
	return Order{
		BaseEntity: newBaseEntity(id),
		totalCost:  0.0,
		user:       user,
	}
}

func main() {
	order := newOrder(1, "testOrder1")
	fmt.Println(order)
	order.addOrderLine(1, "testItem1", 2, 23.32)
	order.addOrderLine(2, "testItem2", 3, 17.45)
	fmt.Println(order)
	order.updOrderLine(1, "testItem1Updated", 21, 3.42)
	order.delOrderLine(2)
	fmt.Println(order.totalCost)
}
