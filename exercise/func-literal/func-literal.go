//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

type LineCallback func(line string)

func processLine(line []string, callback LineCallback) {
	for i, v := range line {
		fmt.Printf("Line %d: %s\n", i, v)
		callback(v)
	}
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	anon := func(line string) {
		slice := []rune(line)
		letters := 0
		digits := 0
		spaces := 0
		punctuation := 0

		for i := 0; i < len(line); i++ {
			if unicode.IsLetter(slice[i]) {
				letters++
			} else if unicode.IsDigit(slice[i]) {
				digits++
			} else if unicode.IsSpace(slice[i]) {
				spaces++
			} else if unicode.IsPunct(slice[i]) {
				punctuation++
			}
		}
		fmt.Printf("Number of letters: %d\n", letters)
		fmt.Printf("Number of digits: %d\n", digits)
		fmt.Printf("Number of spaces: %d\n", spaces)
		fmt.Printf("Number of punctuation marks: %d\n", punctuation)
		fmt.Println()
	}

	processLine(lines, anon)
}
