package main

import "testing"

func TestCalulateTotalPrice(t *testing.T) {
	testcases := []struct {
		input    []string
		expected int
	}{
		{
			input:    []string{"Apple"},
			expected: 35,
		},
		{
			input:    []string{"Banana"},
			expected: 20,
		},
		{
			input:    []string{"Melon"},
			expected: 50,
		},
		{
			input:    []string{"Melon", "Melon"}, // Buy one get one free
			expected: 50,
		},
		{
			input:    []string{"Melon", "Melon", "Melon"},
			expected: 100,
		},
		{
			input:    []string{"Apple", "Banana", "Melon"},
			expected: 105,
		},
		{
			input:    []string{"Apple", "Banana", "Melon", "Melon"},
			expected: 105,
		},
		{
			input:    []string{"Apple", "Banana", "Melon", "Melon", "Melon"},
			expected: 155,
		},
		{
			input:    []string{"Lime"},
			expected: 15,
		},
		{
			input:    []string{"Lime", "Lime"},
			expected: 30,
		},
		{
			input:    []string{"Lime", "Lime", "Lime"}, // Buy three for the price of two
			expected: 30,
		},
		{
			input:    []string{"Apple", "Banana", "Lime"},
			expected: 70,
		},
		{
			input:    []string{"Apple", "Banana", "Lime", "Lime"},
			expected: 85,
		},
		{
			input:    []string{"Apple", "Banana", "Lime", "Lime", "Lime"},
			expected: 85,
		},
		{
			input:    []string{"Apple", "Apple", "Banana", "Melon", "Melon", "Lime", "Lime", "Lime"},
			expected: 170,
		},
	}

	for i, tc := range testcases {
		actual := calculateTotalPrice(tc.input)
		if tc.expected != actual {
			t.Fatalf("Case: %d Expected: %d Actual: %d", i, tc.expected, actual)
		}
	}
}
