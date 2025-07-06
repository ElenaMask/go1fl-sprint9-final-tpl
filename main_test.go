package main

import (
	"testing"
)

func TestGenerateRandomElement(t *testing.T) {
	s := []int{1000, 20, 10, 15, 0}
	for _, v := range s {
		r := len(generateRandomElements(v))
		if v != r {
			t.Error("Excepted a slice with size:", v, "received with size:", r)
		}
	}
	allZero := true
	for _, v := range generateRandomElements(s[0]) {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		t.Error("1000 elemets were filled with zero, does random function work?")
	}
}

func TestMaximum(t *testing.T) {
	slices := [][]int{
		{19, 0, 71, 49, 8},
		{56, 29, 100, 2, 1, 7},
		{71, 8, 4, 999, 781, 9},
		{},
	}
	maxs := []int{71, 100, 999, 0}
	for i, slice := range slices {
		if maximum(slice) != maxs[i] {
			t.Error(i, ":", maximum(slice), "!=", maxs[i])
		}
	}
}

func TestMaxChunks(t *testing.T) {
	data := generateRandomElements(100_000_000)
	max1 := maximum(data)
	max2 := maxChunks(data)
	if max1 != max2 {
		t.Error("maximum and maxChunks found different max on the same data, maxinum:", max1, "maxChunks:", max2)
	}
}
