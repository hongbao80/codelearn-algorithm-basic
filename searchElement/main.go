package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{100, 100, 99, 98, 102, 103}
	//result := searchElement(a, 4)
	a = remove(a, 0)
	fmt.Println(a)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func searchElement(a []int, n int) int {
	sort.Ints(a)
	for i := 0; i < len(a)-1; i++ {
		if a[i] == a[i+1] {
			a = remove(a, i+1)
		}
	}
	h := 0
	for i := len(a) - 1; i > 0; i-- {
		if a[i] != a[i-1] {
			if h == n-1 {
				return a[i]
			}
			h++
		}
	}
	return -1
}
