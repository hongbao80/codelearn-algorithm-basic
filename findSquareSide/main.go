package main

import (
	"fmt"
	"math"
)

func main() {
	result := findSquareSide([]int{0, 1, 0, 1}, []int{0, 1, 1, 0})
	fmt.Println(result)
}

func findSquareSide(x []int, y []int) int {
	canhFirst1 := math.Pow(float64(x[1]) - float64(x[0]), 2)
	canhSecond1 := math.Pow(float64(y[1]) - float64(y[0]), 2)
	first:= int(canhFirst1 + canhSecond1)

	canhFirst2 := math.Pow(float64(x[2]) - float64(x[0]), 2)
	canhSecond2 := math.Pow(float64(y[2]) - float64(y[0]), 2)
	second:= int(canhFirst2 + canhSecond2)

	canhFirst3 := math.Pow(float64(x[3]) - float64(x[0]), 2)
	canhSecond3 := math.Pow(float64(y[3]) - float64(y[0]), 2)
	third := int(canhFirst3 + canhSecond3)
	if first == second {
		return first
	}
	return third
}
