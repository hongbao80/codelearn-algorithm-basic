package main

import (
	"fmt"
)

func main() {
	result := formatString("   abc    a aa  a a      ")
	fmt.Println(result)
}

func formatString(input string) string {
	s := []rune(input)
	for s[0] == ' ' {
		s = delChar(s, 0)
	}

	for s[len(s) - 1] == ' ' {
		s = delChar(s, len(s) - 1)
	}
	for i := 0; i < len(s) - 1; i++ {
		if s[i] == ' ' && s[i+1] == ' ' {
			s = delChar(s, i)
			i--
		}
	}
	return string(s)
}

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}