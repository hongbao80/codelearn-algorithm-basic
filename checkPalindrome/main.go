package main

import "fmt"

func main() {
	result := checkPalindrome("aabaaa");
	fmt.Println(result)
}


func checkPalindrome(inputString string) bool {
	len := len(inputString)
	for i := 0; i < len; i++ {
		if inputString[i] != inputString[len-i-1] {
			return false
		}
	}
	return true
}