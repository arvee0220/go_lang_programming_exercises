package timeparse

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	test := []struct{
		input       string
		expected    Time
		expectError bool
	} {
		{"14:07:33", Time{Hour: 14, Minute: 7, Second: 33}, false},
		{"12:00", Time{}, true},		
		{"25:00:00", Time{Hour: 25, Minute: 0, Second: 0}, true},
	}

	for _, tc := range test {
		result, err := ParseTime(tc.input)
		if tc.expectError {
			if err == nil {
				t.Errorf("ParseTime(%q) expected error, got nil", tc.input)
			}
		} else {
			if err != nil {
				t.Errorf("ParseTime(%q) unexpected error: %v", tc.input, err)
			} else if result != tc.expected {
				t.Errorf("ParseTime(%q) = %v; want %v", tc.input, result, tc.expected)
			}			
		}
	}
}