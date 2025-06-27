package main

import "fmt"

func arrayManipulation(n int32, queries [][]int32) int64 {
	// Create an array of size n+2 to handle the difference array logic
	arr := make([]int64, n+2)
	//     1 2 100
	//     2 5 100
	//     3 4 100

	// Apply the difference array technique
	for _, query := range queries {
		a := query[0]        // 1
		b := query[1]        // 2
		k := int64(query[2]) // 100
		fmt.Println("a", a)
		fmt.Println("b", b)
		fmt.Println("k", k)

		arr[a] += k
		arr[b+1] -= k
	}

	// Compute the prefix sum and find the maximum value
	var max, current int64
	for i := 1; i <= int(n); i++ {
		current += arr[i]
		if current > max {
			max = current
		}
	}

	return int64(max)
}
