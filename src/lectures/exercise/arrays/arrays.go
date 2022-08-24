//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price float32
}

func totalCost(products [4]Product) float32 {
	var sum float32
	for i := 0; i < len(products); i++ {
		sum += products[i].price
	}

	return sum
}

func printInfo(products [4]Product) {
	fmt.Println(products[3])
	fmt.Println(len(products))
	fmt.Println(totalCost(products))
}

func main() {
	products := [4]Product{
		{name: "Butter", price: 1.23},
		{name: "Bread", price: 1.02},
		{name: "Milk", price: 1.57},
	}

	printInfo(products)

	products[3] = Product{name: "Cheese", price: 1.41}

	printInfo(products)
}
