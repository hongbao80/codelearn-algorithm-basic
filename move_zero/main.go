package main

import "fmt"

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}

func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && nums[len(nums) - 1] !=0 {
			nums[i], nums[len(nums) - 1] = nums[len(nums) - 1], 0
		}
	}
	fmt.Println(nums)
}
