//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name  string
	state SecurityTag
}

func changeState(state SecurityTag, tag *Item) {
	tag.state = state
}

func checkout(tags []Item) {
	for i := 0; i < len(tags); i++ {
		changeState(Inactive, &tags[i])
	}
}

func main() {

	//  - Create at least 4 items, all with active security tags
	//  - Store them in a slice or array
	tags := []Item{
		{name: "T-Shirt", state: Active},
		{name: "Shorts", state: Active},
		{name: "Socks", state: Active},
		{name: "Watch", state: Active},
	}

	//  - Deactivate any one security tag in the array/slice
	changeState(Inactive, &tags[2])
	fmt.Println(tags)

	//  - Call the checkout() function to deactivate all tags
	checkout(tags)
	fmt.Println(tags)
}
