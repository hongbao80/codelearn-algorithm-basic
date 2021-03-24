package main

import "fmt"

func main() {
	result := sortByLength([]string{"abc", "", "aaa", "a", "zz"})
	fmt.Print(result)
}

func sortByLength(inputArray []string) []string {
	for {
		count := 0
		for i := 0; i < len(inputArray)-1; i++ {
			if len(inputArray[i]) > len(inputArray[i+1]) {
				count++
				inputArray[i], inputArray[i+1] = inputArray[i+1], inputArray[i]
			}
		}
		if count == 0 {
			break
		}
	}

	return inputArray
}
