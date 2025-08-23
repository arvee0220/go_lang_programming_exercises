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

func quotient(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	quot := nums[0]
	for _, num := range nums[1:] {
		if num == 0 {
			panic("error: cannot divide by zero")
		}
		quot /= num
	}

	return quot
}

func main() {

	a := []int{1, 2, 3, 4, 5}
	b := []int{4, 5, 6}

	ab := append(a, b...)

	c := []float64{1, 2, 3, 4, 5}
	d := []float64{4, 5, 6}
	cd := append(c, d...)

	fmt.Println("All numbers:", ab)

	answer := sum(ab...)

	fmt.Println("The sum is:", answer)

	quotientResult := quotient(cd...)
	fmt.Println("The quotient is:", quotientResult)

}
