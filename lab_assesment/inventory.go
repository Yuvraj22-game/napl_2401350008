package main

type Item struct {
	ID       string
	Name     string
	Quantity int
	Price    float64
}

func findIndex(items []Item, id string) int {
	for i, item := range items {
		if item.ID == id {
			return i
		}
	}
	return -1
}

func AddItem(items []Item, item Item) []Item {
	return append(items, item)
}

func UpdateItem(items []Item, id string, quantity int, price float64) {
	index := findIndex(items, id)
	if index != -1 {
		items[index].Quantity = quantity
		items[index].Price = price
	}
}

func RemoveItem(items []Item, id string) []Item {
	index := findIndex(items, id)
	if index != -1 {
		return append(items[:index], items[index+1:]...)
	}
	return items
}

type Stats struct {
	ItemCount     int
	TotalQuantity int
	TotalValue    float64
}

func ComputeStats(items []Item) Stats {
	ch := make(chan Stats)

	go func() {
		var s Stats
		for _, item := range items {
			s.ItemCount++
			s.TotalQuantity += item.Quantity
			s.TotalValue += float64(item.Quantity) * item.Price
		}
		ch <- s
	}()

	return <-ch
}
