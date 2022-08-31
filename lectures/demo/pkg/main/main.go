package main

import (
	"lectures/lectures/demo/pkg/display"
	"lectures/lectures/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("hello from display")
	msg.Exciting("an exciting hello")
}
