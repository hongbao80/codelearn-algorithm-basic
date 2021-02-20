package main

import (
	"fmt"
	"strings"
)

func main() {
	result := amendTheSentence("XinChao")
	fmt.Println(result)
}

func amendTheSentence(s string) string {
	var result string = ""
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			result += " "
		}
		result += string(v)
	}
	return strings.TrimSpace(result)
}