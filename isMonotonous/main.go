package main

import "fmt"

func main() {
	fmt.Print(isMonotonous([]int{5, 3, -2}))
}

func isMonotonous(sequence []int) bool {
	size := len(sequence)
	if size == 1 {
		return true
	}
	if size == 2 {
		return sequence[0] != sequence[1]
	}
	inc, dec := true, true
	for i := 1; i < size && (inc || dec); i++ {
		inc = inc && sequence[i-1] <= sequence[i]
		dec = dec && sequence[i-1] >= sequence[i]
	}

	return inc || dec
}
