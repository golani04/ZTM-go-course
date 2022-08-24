package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)

	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1 // bread key has default 0

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted, new list:", shoppingList)

	cereal, found := shoppingList["cereal"]
	fmt.Print("Need cereal? ")
	if !found {
		fmt.Println("nope")
	} else {
		fmt.Println("yup", cereal, "boxes")
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println("Total items", totalItems)
}
