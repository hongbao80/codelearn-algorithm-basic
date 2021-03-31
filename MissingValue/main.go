package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(missingValue([]int{7, 2, 5, 3, 5, 3}, []int{7, 2, 5, 4, 6, 3, 5, 3}))
}

func missingValue(a []int, b []int) []int {
	m := make(map[int]bool)
	result := []int{}
	for i := 0; i < len(a); i++ {
		m[a[i]] = true
	}
	for i := 0; i < len(b); i++ {
		if _, ok := m[b[i]]; !ok {
			result = append(result, b[i])
		}
	}
	sort.Ints(result)
	return result
}
