package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(insideCircle([]int{1, 1}, []int{2, 2}, 1))
}


func insideCircle(point []int, center []int, radius int) bool {
	dis1 := math.Pow(float64(center[0]) - float64(point[0]), 2)
	dis2 := math.Pow(float64(center[1]) - float64(point[1]), 2)
	dis := float64(math.Sqrt(dis1 + dis2))
	if dis <= float64(radius) {
		return true
	}
	return false
}