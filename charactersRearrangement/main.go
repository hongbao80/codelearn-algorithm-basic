package main

import "fmt"

func main() {
	result := charactersRearrangement("b", "a")
	fmt.Println(result)
}

func charactersRearrangement(string1 string, string2 string) bool {
	if len(string1) != len(string2) {
		return false
	}
	mapCmp := make(map[int32]bool)
	for _, ch := range string1 {
		mapCmp[ch] = true
	}
	for _, ch := range string2 {
		if _, ok := mapCmp[ch]; !ok {
			return false
		}
	}
	return true
 }