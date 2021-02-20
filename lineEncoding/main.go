package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "aabbbc"
	result := lineEncoding(s)
	fmt.Println(result)
}

func lineEncoding(s string) string {
	s += " "
	ru := []rune (s)
	count := 1
	result := ""
	for i := 0; i < len(ru) - 1; i++ {
		if ru[i] == ru[i + 1] {
			count++
		} else {
			if count > 1 {
				result += strconv.Itoa(count)
			}
			result += string(ru[i])
			count = 1
		}
	}
	return result
}