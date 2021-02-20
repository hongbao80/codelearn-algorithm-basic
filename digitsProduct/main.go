package main

import "fmt"

func main() {
	result := digitsProduct(100)
	fmt.Println(result)
}

func digitsProduct(product int) int {
	if product == 1 {
		return 1
	}
	if product == 0 {
		return 10
	}
	temp := 0
	for i := 9; i >= 2; i-- {
		for product % i == 0 {
			temp = temp* 10 + i
			product /= i
		}
	}
	fmt.Println(temp)
	result := 0
	for temp > 0 {
		result = result * 10 + temp % 10
		temp /= 10
	}
	if product == 1 {
		return result
	}
	return -1
}