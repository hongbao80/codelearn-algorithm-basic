package main

import (
	"fmt"
)

func main() {
	result := numberZeroDigits(100)
	fmt.Println(result)
}

//func numberZeroDigits(n int64) int64 {
//	var cur uint64  = 1
//	num2 := 0
//	num5 := 0
//	var i uint64  = 0
//	for i = 2; i <= uint64(n); i++ {
//		cur *= i
//		for cur % 2 == 0 {
//			num2 ++
//			cur /= 2
//		}
//		for cur % 5 == 0 {
//			num5 ++
//			cur /= 5
//		}
//
//	}
//	return int64(math.Min(float64(num2), float64(num5)))
//}

func numberZeroDigits(n int64) int64 {
	var d int64  = 0
	var k int64 = 5
	for k <= n {
		d += n / k
		k *= 5
	}
	return d
}

