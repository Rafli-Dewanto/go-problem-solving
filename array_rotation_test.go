package main

import (
	"fmt"
	"testing"
)

func TestArrayRotation(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	var d int = 6

	newSlice := make([]int, len(arr))
	n := len(arr)

	for idx, value := range arr {
		rawDifference := idx - d

		elementIndex := findIndex(rawDifference, n)

		newSlice[elementIndex] = value
	}
	fmt.Println(newSlice)
}

func findIndex(val, size int) int {
	if val >= 0 {
		return val
	}
	return findIndex(val+size, size)
}
