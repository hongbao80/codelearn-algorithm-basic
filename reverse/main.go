package main

import (
	"fmt"
	"log"
)

func main() {
	input := "abcdef"
	result := reverse(input)
	fmt.Println(result)
}

func getReverseString(input string) string {
	if len(input) == 0{
		log.Fatal("Input string is empty")
	}
	//string to []rune
	ru := []rune(input)

	// get length of rune
	lengthInput := len(ru)

	// loop to length/2 of string input
	// assign character values which have index i to character values which have index length - i - 1 and reverse
	var ru2 []rune
	for i := lengthInput - 1; i >=0; i-- {
		//ru[lengthInput - 1 - i], ru[i] = ru[i], ru[lengthInput - i - 1]
		ru2 = append(ru2, ru[i])
	}
	return string(ru2)
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
