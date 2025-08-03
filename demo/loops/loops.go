package main

import "fmt"

func main() {
	sum := 0
	fmt.Println("Sum is", sum)

	for i := 1; i <= 10; i++ {
		sum += i
		fmt.Println("Sum is", sum)
	}

	//While loop
	for sum > 10 {
		sum -= 5
		fmt.Println("Decrement. Sum is", sum)
	}

	// Infinite loop
	for {
		fmt.Println("Infinite loop. Sum is", sum)
		sum -= 1
		
		if sum <= 0 {
			fmt.Println("Breaking out of infinite loop. Final sum is", sum)
			break
		}
	}
}
