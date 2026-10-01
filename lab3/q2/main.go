package main

// 1. Pointer Manipulation — Referencing and Dereferencing (Exp 5)
// Write a Go program demonstrating pointer usage:
// (1) declare a variable, print its address using & and access its value using *;
// (2) write a function that accepts a pointer parameter and modifies the original variable (pass-by-reference) — print before and after values;
// (3) allocate a struct using new() and access/modify its fields through the pointer.

import "fmt"

type Person struct {
	Name string
	Age  int
	Job  string
}

// part 1
func modifyValue(a *int) {
	*a = 100
}

func modifyStructure(p *Person) {
	p.Name = "default person"
	p.Age = 18
	p.Job = "default job"

}

func main() {
	//part 2
	x := 200
	//pointer declaration
	ptr := &x
	//pointer print
	println("value of ptr : ", ptr)
	//pointer value after variable modification
	modifyValue(ptr)
	println("\nvalue of variable : ", x)

	//part 3
	p1 := new(Person)

	fmt.Println("\nperson details before modification:")
	fmt.Println("Name : ", p1.Name)
	fmt.Println("Age : ", p1.Age)
	fmt.Println("Job : ", p1.Job)

	modifyStructure(p1)
	fmt.Println("\nperson details after modification:")
	fmt.Println("Name : ", p1.Name)
	fmt.Println("Age : ", p1.Age)
	fmt.Println("Job : ", p1.Job)
}
