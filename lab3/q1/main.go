package main

import "fmt"

//1. Struct Person with Associated Methods (Exp 4)
// Define a Go struct named Person with fields: Name (string), Age (int), Job (string), and Salary (float64).
// Add two methods with appropriate receivers: one to read data into the struct from user input, and one to print all fields in a formatted manner.
// Create at least two Person objects in main() and call both methods on each.
// Push the source file to your GitHub repository.

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func (p Person) ReadData() Person {
	fmt.Print("Enter Name: ")
	fmt.Scanln(&p.Name)

	fmt.Print("Enter Age: ")
	fmt.Scanln(&p.Age)

	fmt.Print("Enter Job: ")
	fmt.Scanln(&p.Job)

	fmt.Print("Enter Salary: ")
	fmt.Scanln(&p.Salary)

	return p
}

func (p *Person) PrintData() {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Age: %d\n", p.Age)
	fmt.Printf("Job: %s\n", p.Job)
	fmt.Printf("Salary: %.2f\n", p.Salary)

}

func main() {
	var p1 Person
	// var p2 Person

	fmt.Println("Enter details for Person 1:")
	p1 = p1.ReadData()
	fmt.Println("Details of Person 1:")
	p1.PrintData()

	// fmt.Println("Enter details for Person 2:")
	// p2 = p2.ReadData()
	// fmt.Println("Details of Person 2:")
	// p2.PrintData()
}
