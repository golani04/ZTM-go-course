package main

import "fmt"

func printInfo(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")

	for i := 0; i < len(slice); i++ {
		item := slice[i]
		fmt.Println(item)
	}
}

func main() {
	route := []string{"Grocery", "Department Store", "Salon"}
	printInfo("Route 1", route)

	route = append(route, "Home")
	printInfo("Route 2", route)

	fmt.Println()
	fmt.Println(route[0], "visited")
	fmt.Println(route[1], "visited")

	route = route[2:]
	printInfo("Remaining Locations", route)
}
