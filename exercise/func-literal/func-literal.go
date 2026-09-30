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

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	func() {
		letters := 0
		digits := 0
		spaces := 0
		punctuation := 0

		for _, v := range lines {
			for _, r := range v {
				if unicode.IsLetter(r) {
					letters++
				} else if unicode.IsDigit(r) {
					digits++
				} else if unicode.IsSpace(r) {
					spaces++
				} else if unicode.IsPunct(r) {
					punctuation++
				}
			}
		}

		fmt.Printf("Number of letters: %d\n", letters)
		fmt.Printf("Number of digits: %d\n", digits)
		fmt.Printf("Number of spaces: %d\n", spaces)
		fmt.Printf("Number of punctuation marks: %d\n", punctuation)
	}()
}
