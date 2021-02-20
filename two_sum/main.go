package main

import "fmt"

func main() {
	result := twoSum([]int{3, 2, 4}, 6)
	fmt.Print(result)
	//maps := make(map[int]int)
	//value, ok := maps[9]
	//fmt.Println(ok)
	//fmt.Println(value)
}

func twoSum(nums []int, target int) []int {
	var maps = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		rest := target - nums[i]
		if _, ok := maps[rest]; !ok {
			maps[nums[i]] = i
			continue
		}
		return []int{maps[rest], i}
	}
	return nil
}
