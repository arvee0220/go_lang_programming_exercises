package display

import "fmt"

func hello(msg string) {
	fmt.Println(msg) // Private function
}

func Display(msg string) {
	fmt.Println(msg) // Public function
	hello("hello() is a private function in display package") // Call to private function
}

