//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

func printAssemblyLine(title string, line []Part) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := range line {
		part := line[i]
		fmt.Println(part)
	}

}


func main() {
	// Using a slice, create an assembly line that contains type Part
	assemblyLine := []Part{"Engine", "Transmission", "Chassis"}

	printAssemblyLine("Initial Assembly Line", assemblyLine)

	assemblyLine = append(assemblyLine, "Wheels", "Brakes")

	printAssemblyLine("Assembly Line after adding new parts", assemblyLine)

	slice := assemblyLine[3:5]

	printAssemblyLine("Final Assembly Line with new parts", slice)
	
}
