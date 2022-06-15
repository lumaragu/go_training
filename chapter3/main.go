package main

import (
	"fmt"
	"go_training/chapter3/order"
)

func main() {
	order := order.New(1, "testOrder1")
	fmt.Println(order)
	order.AddOrderLine(1, "testItem1", 2, 23.32)
	order.AddOrderLine(2, "testItem2", 3, 17.45)
	fmt.Println(order)
	order.UpdateOrderLine(1, "testItem1Updated", 21, 3.42)
	order.DeleteOrderLine(2)
	fmt.Println(order.TotalCost)
}
