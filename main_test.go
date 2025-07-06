package main

import (
	"testing"
)

var maximumTestCases = []struct {
	name     string
	data     []int
	expected int
}{
	{
		name:     "empty slice",
		data:     []int{},
		expected: 0,
	},
	{
		name:     "single element",
		data:     []int{7},
		expected: 7,
	},
	{
		name:     "all equal small",
		data:     []int{5, 5, 5, 5},
		expected: 5,
	},
	{
		name:     "ascending 10 elems",
		data:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		expected: 10,
	},
	{
		name:     "descending 12 elems",
		data:     []int{100, 90, 80, 70, 60, 50, 40, 30, 20, 10, 5, 1},
		expected: 100,
	},
	{
		name:     "mixed 15 elems",
		data:     []int{3, 99, 2, 88, 1, 77, 4, 66, 5, 55, 6, 44, 7, 33, 8},
		expected: 99,
	},
	{
		name:     "middle is max",
		data:     []int{1, 2, 3, 4, 100, 4, 3, 2, 1},
		expected: 100,
	},
	{
		name:     "many same max",
		data:     []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		expected: 10,
	},
}

func TestMaximum(t *testing.T) {
	for _, tt := range maximumTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.expected {
				t.Errorf("maximum(%v) = %d; want %d", tt.data, got, tt.expected)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	for _, tt := range maximumTestCases {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.expected {
				t.Errorf("maxChunks(%v) = %d; want %d", tt.data, got, tt.expected)
			}
		})
	}
}

var generateRandomElementsTests = []struct {
	name     string
	size     int
	expected int
}{
	{"ZeroSize", 0, 0},
	{"SmallSize", 3, 3},
	{"NegativeSize", -5, 0},
	{"LargeSize", 1000, 1000},
}

func TestGenerateRandomElement(t *testing.T) {
	for _, tc := range generateRandomElementsTests {
		t.Run(tc.name, func(t *testing.T) {
			result := generateRandomElements(tc.size)
			if len(result) != tc.expected {
				t.Errorf("expected length %d, got %d", tc.expected, len(result))
			}
		})
	}
}

func TestGenerateRandomElementNotZero(t *testing.T) {
	data := generateRandomElements(1000)
	allZero := true
	for _, v := range data {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		t.Errorf("all 1000 elements in randomly filled slice are zeros")
	}
}
