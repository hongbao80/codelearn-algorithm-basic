package main

import (
	"fmt"
)

func main() {
	result := differentSubstringsTrie("abac")
	fmt.Println(result)
}


func differentSubstringsTrie(inputString string) int {
	count := 0
	result := make(map[string]int)
	s := []rune(inputString)
	for i := 0; i < len(s); i++ {
		tempStr := ""
		for j := i; j < len(s); j++ {
			tempStr = fmt.Sprintf("%s%c", tempStr, s[j])
			if _, ok := result[tempStr];!ok {
				count++
				fmt.Println(tempStr)
				result[tempStr] = 1
			}
		}
	}
	return count
}