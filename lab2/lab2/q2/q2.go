package main

import "fmt"

func main() {
	// The slice records the books currently on a library display shelf.
	shelf := []string{"Clean Code", "The Pragmatic Programmer", "Go in Action"}
	fmt.Println("Initial shelf:", shelf)

	shelf = append(shelf, "Design Patterns")
	fmt.Println("After adding a book:", shelf)

	removedIndex := 1
	removedBook := shelf[removedIndex]
	shelf = append(shelf[:removedIndex], shelf[removedIndex+1:]...)
	fmt.Printf("After removing index %d (%s): %v\n", removedIndex, removedBook, shelf)

	shelf[0] = "Effective Go"
	fmt.Println("After updating index 0:", shelf)

	// The map stores the number of copies available for each book.
	stock := map[string]int{
		"Effective Go":    3,
		"Go in Action":    2,
		"Design Patterns": 1,
	}
	fmt.Println("Initial stock:", stock)

	stock["The Go Programming Language"] = 4
	fmt.Println("After inserting a title:", stock)

	delete(stock, "Go in Action")
	fmt.Println("After deleting Go in Action:", stock)

	if copies, exists := stock["Effective Go"]; exists {
		fmt.Printf("Lookup Effective Go: %d copies available\n", copies)
	} else {
		fmt.Println("Lookup Effective Go: title not found")
	}
}
