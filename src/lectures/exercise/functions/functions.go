//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello, ", name)
}

func randomMsg() string {
	return "Random message is printed!"
}

func add(x, y, z int) int {
	return x + y + z
}

func singleAnyNumber() int {
	return 1
}

func twoAnyNumbers() (int, int) {
	return 1, 2
}

func main() {
	greet("Leonid")
	fmt.Println(randomMsg())

	result := add(7, 8, 9)
	fmt.Println("Result is ", result)

	y, z := twoAnyNumbers()
	anotherResult := add(singleAnyNumber(), y, z)
	fmt.Println("Result is ", anotherResult)

}
