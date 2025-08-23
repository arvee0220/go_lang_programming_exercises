package main

import (
	"coursecontent/demo/pkg/display" // refer to go.mod module path
	"coursecontent/demo/pkg/msg"
)

func main() {
	msg.Hi()
	display.Display("display.Display(msg string) Hello from display")
	msg.Exciting("msg.Exciting(msg string) an exciting message")
}