package main

import "fmt"

func main() {
	slice := make([]string, 3)
	slice = []string{"Hello", "World", "!"}

	for i, elem := range slice {
		fmt.Print(i, ") ", elem, ":")

		for _, ch := range elem {
			fmt.Printf(" %q", ch)
		}
		fmt.Println()
	}
}
