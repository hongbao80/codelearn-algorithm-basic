package main

import "fmt"

func main() {
	arr1 := []interface{} {1, 2, 3, 4, nil}
	arr2 := []interface{} {5, 6, 9, 99, nil}
	result := findCommonItems2(arr1, arr2)
	fmt.Print(result)

}

// this is not best, cause time complexity is O(n^2)
func findCommonItem(arr1, arr2 []interface{}) bool {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[1] == arr2[j] {
				return true
			}
		}
	}
	return false
}

func findCommonItems2(arr1, arr2 []interface{}) bool {
	hash := make(map[interface{}]bool)
	for i := 0; i < len(arr1); i++ {
		if !hash[arr1[i]] {
			hash[arr1[i]] = true
		}
	}

	for i := 0; i < len(arr2); i++ {
		if hash[arr2[i]] {
			return true
		}
	}
	return false
}
