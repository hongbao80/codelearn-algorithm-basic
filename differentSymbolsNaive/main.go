package main

import "fmt"

func main() {
	result := differentSymbolsNaive("cabca")
	fmt.Println(result)
}

func differentSymbolsNaive(s string) int {
	mapInt := make(map[int32]int)
	for _, v := range s {
		if num, ok := mapInt[v]; !ok {
			mapInt[v] = 1
		} else {
			mapInt[v] = num + 1
		}
	}
	return len(mapInt)
}
