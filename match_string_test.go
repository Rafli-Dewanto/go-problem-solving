package main

import (
	"fmt"
	"testing"
)

func TestMatchString(t *testing.T) {
	stringList := []string{"ab", "ab", "abc"}
	queries := []string{"ab", "abc", "bc"}
	resultMap := make(map[string]int)

	for _, s := range stringList {
		resultMap[s] += 1
	}

	results := make([]int, len(queries))

	for key, v := range queries {
		value, ok := resultMap[v]
		if ok {
			results[key] = value
		} else {
			results[key] = 0
		}
	}

	fmt.Println(results)
}
