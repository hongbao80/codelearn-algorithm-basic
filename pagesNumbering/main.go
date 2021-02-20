package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(pagesNumbering(1000))
}


func pagesNumbering(input int) int {
	if input < 10 {
		return input
	}
	x := len(strconv.Itoa(input))
	count := 0
	k := 1
	for i := 1; i <= x - 1; i++ {
		count += (9 * i) * k
		k *= 10
	}
	rest := (input - int(math.Pow(10, float64(x - 1))) + 1)  * x
	return count + rest
}