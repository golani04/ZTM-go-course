package main

import "fmt"

func sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	return sum
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 6, 8}

	all := append(a, b...)
	fmt.Println(sum(all...))

	answer := sum(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(answer)

}
