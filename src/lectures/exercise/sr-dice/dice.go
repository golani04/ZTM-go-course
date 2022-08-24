//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDice(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	// time.Now().UnixNano(), which yields a constantly-changing number.
	rand.Seed(time.Now().UnixNano())

	dices, rolls, sides := 2, 1, 6

	total := 0
	for i := 0; i < rolls; i++ {
		dd := dices
		for dd > 0 {
			num := rollDice(sides)
			fmt.Println("Rolled #", (i + 1), "dice is", num)
			total += num
			dd--
		}

		fmt.Println("Total is", total)
		if dices == 2 && total == 2 {
			fmt.Println("Snake Eyes!")
		} else if total == 7 {
			fmt.Println("Lucky 7")
		} else if total%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

}
