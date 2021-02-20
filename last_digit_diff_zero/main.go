package main

import (
	"fmt"
)

func main() {
	num:= numberOfZero(5)
	fmt.Println(num)
}

func numberOfZero(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
		for result%10 == 0 {
			result /= 10
		}
	}
	return result % 10
}