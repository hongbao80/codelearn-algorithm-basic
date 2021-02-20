package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {
	maps := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if !maps[nums[i]] {
			maps[nums[i]] = true
			continue
		}
		return true
	}
	return false
}
