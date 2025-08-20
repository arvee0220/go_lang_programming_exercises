package main

import "fmt"

// Variadic function to sum an arbitrary number of integers
func sum(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func main() {

	a := []int{1, 2, 3, 4, 5}
	b := []int{4, 5, 6}

	all := append(a, b...)

	fmt.Println("All numbers:", all)

	answer := sum(all...)

	fmt.Println("The sum is:", answer)

}
