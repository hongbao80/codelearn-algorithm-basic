package main

import "fmt"

func main() {
	result := isPangram("The quick brown fox jumps over the lazy dog.")
	fmt.Println(result)
}


func isPangram(sentence string) bool {
	listCha := make(map[int32]int)
	for _, ch := range sentence {
		listCha[ch] = 1
	}
	var i int32
	for i = 65; i <= 90; i++ {
		_, ok := listCha[i]
		_, ok2 := listCha[i + 32]
		if  !ok && !ok2 {
			return false
		}
	}
	return true
}