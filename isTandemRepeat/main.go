package main

import "fmt"

func main() {
	isTandemRepeat("tandemtandem")
}

func isTandemRepeat(inputString string) bool {
	if len(inputString) % 2 != 0 {
		return false
	}
	ru := []rune(inputString)
	len := len(ru)
	head := ru[0:len/2]
	tail := ru[len/2:len]
	fmt.Print(string(head)," ", string(tail))
	if string(head) ==  string(tail) {
		return true
	}
	return false
}
