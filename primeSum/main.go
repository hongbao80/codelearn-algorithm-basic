package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(2 % 22082018)
	fmt.Println(primeSum(5))
}

func primeSum(n int) int {

	if n == 2 {
		return 2
	}
	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

func isPrime(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}