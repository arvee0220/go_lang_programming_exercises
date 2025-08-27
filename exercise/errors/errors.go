//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"	
	"strings"
	"strconv"	
)

type Time struct {
	Hour int
	Minute int
	Second int
}

type TimeParseError struct {
	msg string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParseTime(timeStr string) (Time, error) {
	parts := strings.Split(timeStr, ":")

	if len(parts) != 3 {
		return Time{}, &TimeParseError{
			msg: "time string must be in HH:MM:SS format",
			input: timeStr,
		}
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour:%v", err), timeStr}
	}
	if hour < 0 || hour > 23 {
		return Time{}, &TimeParseError{"Hour must be between 0 and 23", timeStr}
	}

	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("Error parsing minute:%v", err), timeStr}
	}
	if minute < 0 || minute > 59 {
		return Time{}, &TimeParseError{"Minute must be between 0 and 59", timeStr}
	}

	second, err := strconv.Atoi(parts[2])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("Error parsing second:%v", err), timeStr}
	}
	if second < 0 || second > 59 {
		return Time{}, &TimeParseError{"Second must be between 0 and 59", timeStr}
	}


	return Time{Hour: hour, Minute: minute, Second: second}, nil
}

