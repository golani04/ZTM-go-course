//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operation int

const (
	add Operation = iota
	sub
	mul
	div
)

func (o Operation) calculate(n1, n2 int) int {
	switch o {
	case add:
		return n1 + n2
	case sub:
		return n1 - n2
	case mul:
		return n1 * n2
	case div:
		return n1 / n2
	}
	panic("unhandled operation")
}

func main() {
	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
