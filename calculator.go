package main

import "fmt"

// Functions
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mult(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

// Main
func main() {
	var num1 int
	var num2 int
	var operator string

	fmt.Print("Type a number: ")
	fmt.Scan(&num1)

	fmt.Print("Type an operator: ")
	fmt.Scan(&operator)

	fmt.Print("Type a number: ")
	fmt.Scan(&num2)

	var result int
	switch operator {
	case "+":
		result = add(num1, num2)
	case "-":
		result = sub(num1, num2)
	case "*":
		result = mult(num1, num2)
	case "/":
		result = div(num1, num2)
	}

	fmt.Println("The result is:", result)
}
