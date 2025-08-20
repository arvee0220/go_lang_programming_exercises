//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operation byte

const (
	add Operation = iota
	sub
	mul
	div
)

func (c Operation) calculate(a, b int) (int, error){
	switch c {
	case add:
		return a + b, nil
	case sub:
		return a - b, nil
	case mul:
		return a * b, nil
	case div:
		if b == 0 {
			panic("error: cannot divide by zero")
		}
		return a / b, nil
	}
	panic("error: unknown operation")
}

func main() {
	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
