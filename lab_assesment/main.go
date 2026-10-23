package main

import (
	"fmt"
)

func main() {
	items := []Item{}

	items = AddItem(items, Item{"1", "Laptop", 5, 500000})
	fmt.Println("After adding:", items)

	items = AddItem(items, Item{"2", "Iphone", 1, 50000})
	fmt.Println("After adding:", items)

	items = RemoveItem(items, "1")
	fmt.Println("After removing:", items)

	UpdateItem(items, "1", 10, 1000000)
	fmt.Println("After updating:", items)

	stats := ComputeStats(items)
	fmt.Println("Item count:", stats.ItemCount)
	fmt.Println("Total quantity:", stats.TotalQuantity)
	fmt.Println("Total value:", stats.TotalValue)
}
