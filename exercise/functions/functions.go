//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello,", name)
}

func message(message string) string {
	return message
}

func addThreeNumbers(a,b,c int) int {
	return a + b + c
}

func returnNumber() int {
	return 11
}

func returnTwoNumbers() (int, int) {
	return 20, 30
}

func main() {
	greet("Alice")

	messageToPrint := message("Welcome to the Go programming world!")

	fmt.Println(messageToPrint)
	num1, num2 := returnTwoNumbers()
	num3 := returnNumber()
	fmt.Println("First number:", num1, "\n","Second number:", num2, "\n", "Third number:", num3)
	total := addThreeNumbers(num1, num2, num3)
	fmt.Println("Total of three numbers:", total)

}
