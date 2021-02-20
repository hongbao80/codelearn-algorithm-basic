package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 30, 4}
	arr2 := []int{6, 4, 20, 2}
	arr := merge(arr1, arr2)
	result := sort(arr)
	fmt.Println(result)
}


func merge(arr1, arr2 []int) []int{
	for i := 0; i < len(arr2); i++ {
		arr1 = append(arr1, arr2[i])
	}
	return arr1
}

func sort(arr []int) []int {
	for i := 0; i < len(arr) - 1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] >= arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}