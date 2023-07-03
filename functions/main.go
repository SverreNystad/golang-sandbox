package main

import (
	"fmt"
)

const Pi = 3.14

// This is a simple calculator program which can perform basic arithmetic operations
// The program uses higher order functions to perform the operations
// The following are the operations performed by the program: Addition, Subtraction, Multiplication, Division, Modulus, Exponent. But only for integers
func calculator(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Addition
func add(x int, y int) int {
	return x + y
}

// Subtraction
func sub(x int, y int) int {
	return x - y
}

// Multiplication
func mul(x int, y int) int {
	return x * y
}

// Division
func div(x int, y int) int {
	return x / y
}

// Modulus
func mod(x int, y int) int {
	return x % y
}

// One can also declare the variables like this
func exponent(x, y int) int {
	if y == 0 {
		return 1
	}
	current_x := x
	for i := 1; i < y; i++ {
		current_x *= x
	}
	return current_x
}

// Square root of a number
func sqrt(x float64) float64 {
	root := 1.0
	root = float64(1)

	for i := 1; i < 10; i++ {
		root -= (root*root - x) / (2 * root)
	}
	return root
}

func Swap(x string, y string) (string, string) {
	return y, x
}

// Named return values. As x and y are defined in the functions expected output, we can just use return but the z will not be returned
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	z := 10
	z++
	return // This will return x and y but not z
}

func Calculate(x int, y int, operationInput string) int {
	operation := interpretOperation(operationInput)
	return calculator(x, y, operation)
}

func interpretOperation(input string) func(int, int) int {
	switch input {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return mul
	case "/":
		return div
	case "%":
		return mod
	case "^":
		return exponent
	default:
		panic("Unsupported operation")

	}
}
func main() {
	fmt.Println("Calculate 2 add 3 =  ", calculator(2, 3, add))
	fmt.Println("Calculate 2 sub 3 =  ", calculator(2, 3, sub))
	fmt.Println("Calculate 2 mul 3 =  ", calculator(2, 3, mul))
	fmt.Println("Calculate 2 div 3 =  ", calculator(2, 3, div))
	fmt.Println("Calculate 2 mod 3 =  ", calculator(2, 3, mod))
	fmt.Println("Calculate 2 exp 3 =  ", calculator(2, 3, exponent))
	fmt.Print("Swap 2 and 3 = ")
	fmt.Println(Swap("2", "3"))
	fmt.Println("Named return values")
	fmt.Println(split(17))
}
