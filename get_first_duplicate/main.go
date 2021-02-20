package main

import "fmt"

func main() {
	result := getFirstElementRecurring([]int {2, 1, 1, 2, 3, 5,1 ,2, 4})
	fmt.Println(result)
}

func getFirstElementRecurring(arr []int) int {
	hash := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if _, ok := hash[arr[i]]; !ok {
			hash[arr[i]] = true
			continue
		}
		return arr[i]
	}
	return -1
}