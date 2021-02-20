package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(truncateString("33334536"))
}


func truncateString(s string) string {
	ru := []rune(s)
	i, _ := strconv.Atoi(string(ru[0]))
	for i % 3 == 0 {
		ru = remove(ru, 0)
		i, _ = strconv.Atoi(string(ru[0]))
	}
	i, _ = strconv.Atoi(string(ru[len(ru) - 1]))
	for i % 3 == 0 {
		ru = remove(ru, len(ru) - 1)
		i, _ = strconv.Atoi(string(ru[len(ru) - 1]))
	}
	first, _ := strconv.Atoi(string(ru[0]))
	last, _ := strconv.Atoi(string(ru[len(ru) - 1]))
	for (first + last) % 3 == 0 {
		ru = remove(ru, 0)
		ru = remove(ru, len(ru) - 1)
		first, _ = strconv.Atoi(string(ru[0]))
		last, _ = strconv.Atoi(string(ru[len(ru) - 1]))
	}
	return string(ru)
}

func remove(slice []rune, s int) []rune {
	return append(slice[:s], slice[s+1:]...)
}