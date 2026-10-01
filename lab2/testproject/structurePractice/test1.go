package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Marks float64
}

func main() {
	Yuvraj := Student{
		Name:  "Yuvraj",
		Age:   20,
		Marks: 106.5,
	}
	class := []Student{
		{"Yuvraj", 21, 106.5},
		{"Tarun", 20, 20.0},
	}

	for i, value := range class {
		fmt.Printf("Student %d: %s, Age: %d, Marks: %.2f\n", i+1, value.Name, value.Age, value.Marks)
	}

	// fmt.Println(Yuvraj)
	fmt.Printf("Name of the Student : %s\n", Yuvraj.Name)
	fmt.Printf("Age of the Student : %d\n", Yuvraj.Age)
	fmt.Printf("Marks of the Student : %.2f\n", Yuvraj.Marks)

}
