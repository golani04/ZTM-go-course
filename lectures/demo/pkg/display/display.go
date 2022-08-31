package display

import "fmt"

// if function is capitalized this function will exported
// capitalized functions are public and lowercased are private
func Display(msg string) {
	fmt.Println(msg)
}
