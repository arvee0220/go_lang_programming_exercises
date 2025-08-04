//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	Length float64
	Width  float64
}

func Area(rect Rectangle) float64 {
	return rect.Length * rect.Width
}

func main() {

	rect1 := Rectangle{Length: 5, Width: 3}
	area1 := Area(rect1)
	fmt.Printf("Area of rectangle 1: %.2f\n", area1)

	rect1.Length *= 2
	rect1.Width *= 2
	area1 = Area(rect1)
	fmt.Printf("New area of rectangle 1 after doubling size: %.2f\n", area1)

}
