package main

import "fmt"

// Division function
func divide(a, b float64) float64 {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

// Multiplication function
func multiply(a, b float64) float64 {
	return a * b
}

// Addition function
func add(a, b float64) float64 {
	return a + b
}

// Subtraction function
func subtract(a, b float64) float64 {
	return a - b
}

// Hello World function
func helloWorld() {
	fmt.Println("Hello, World!")
}

// FlyHigh function
func flyHigh() {
	fmt.Println("FlyHigh")
}

// MyName function
func myName() {
	fmt.Println("My Name")
}

func main() {
	// Division
	result := divide(10, 2)
	fmt.Println("Division:", result)

	// Multiplication
	result = multiply(5, 6)
	fmt.Println("Multiplication:", result)

	// Addition
	result = add(3, 4)
	fmt.Println("Addition:", result)

	// Subtraction
	result = subtract(8, 5)
	fmt.Println("Subtraction:", result)

	// Hello World
	helloWorld()

	// FlyHigh
	flyHigh()

	// My Name
	myName()
}
