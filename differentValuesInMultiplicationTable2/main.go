package main

import "fmt"

func main() {
	result := differentValuesInMultiplicationTable2(3, 2)
	fmt.Println(result)
}

func differentValuesInMultiplicationTable2(n int, m int) int {
	mapCal := make(map[int]int)
	result := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if _, ok := mapCal[i*j]; !ok {
				result++
				mapCal[i*j] = 1
			}
		}
	}
	return result
}