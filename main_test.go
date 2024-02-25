package main

import "testing"

func TestRomanToInt(t *testing.T) {
	testCase := []struct {
		input    string
		expected int
	}{
		{
			input:    "I",
			expected: 1,
		},
		{
			input:    "V",
			expected: 5,
		},
		{
			input:    "X",
			expected: 10,
		},
		{
			input:    "C",
			expected: 100,
		},
		{
			input:    "D",
			expected: 500,
		},
		{
			input:    "M",
			expected: 1000,
		},
		{
			input:    "IV",
			expected: 4,
		},
		{
			input:    "IX",
			expected: 9,
		},
		{
			input:    "XL",
			expected: 40,
		},
		{
			input:    "XC",
			expected: 90,
		},
		{
			input:    "CD",
			expected: 400,
		},
		{
			input:    "CM",
			expected: 900,
		},
		{
			input:    "III",
			expected: 3,
		},
		{
			input:    "LVIII",
			expected: 58,
		},
		{
			input:    "MCMXCIV",
			expected: 1994,
		},
	}

	for _, v := range testCase {
		ac := romanToInt(v.input)
		ep := v.expected
		if ac != ep {
			t.Errorf("expected %d but got %d", ep, ac)
		}
	}
}
