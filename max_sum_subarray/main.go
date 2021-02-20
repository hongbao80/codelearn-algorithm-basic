package main

import (
	"fmt"
	"math"
)

func main() {
	result := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	max := nums[0]
	cur := 0
	for _, v := range nums {
		if cur < 0 {
			cur = 0
		}
		cur += v
		if cur > max {
			max = cur
		}
		max = int(math.Max(float64(max), float64(cur)))
	}
	return max
}

//func maxSubArray(nums []int) int {
//	max := nums[0]
//	for i := 0; i < len(nums); i++ {
//		current := nums[i]
//		for j := i + 1; j < len(nums); j++ {
//			current += nums[j]
//			if current > max {
//				max = current
//			}
//		}
//		if current > max {
//			max = current
//		}
//	}
//	return max
//}
