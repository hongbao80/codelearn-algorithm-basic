package main

import "fmt"

func main() {
	rotate([]int{1,2,3,4,5,6,7}, 3)
}

func rotate(nums []int, k int)  {
	for i := 0; i < k; i++ {
		last := nums[len(nums) - 1]
		nums = remove(nums, len(nums) - 1)
		nums = append([]int{last}, nums...)
	}
	fmt.Println(nums)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

