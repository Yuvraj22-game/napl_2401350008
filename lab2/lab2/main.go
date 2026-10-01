package main

import (
	mathutil "MyProject1/test"
	factorial "MyProject1/test1"
	"fmt"
)

func main() {

	// var a string

	// var b string

	// fmt.Print("Enter first string: ")

	// fmt.Scan(&a)

	// fmt.Print("Enter second string: ")

	// fmt.Scan(&b)

	// result := mathutil.AddString(a, b)

	// fmt.Printf("The combined string is: %s\n", result)
	// var a string

	// var b string

	// fmt.Print("Enter first string: ")

	// fmt.Scan(&a)

	// fmt.Print("Enter second string: ")

	// fmt.Scan(&b)

	// result := mathutil.AddString(a, b)

	// fmt.Printf("The combined string is: %s\n", result)
	result := mathutil.ReverseString("HelloWorld")

	fmt.Printf("The reversed string is: %s\n", result)

	vowelCount := mathutil.CountVowels("YuvrajSaini")

	fmt.Printf("The number of vowels in the string is: %d\n", vowelCount)

	fact := factorial.Factorial(5)

	fmt.Printf("The factorial of 5 is: %d\n", fact)

	pow := factorial.Power(2, 3)

	fmt.Printf("The result of 2 raised to the power of 3 is: %d\n", pow)

}
