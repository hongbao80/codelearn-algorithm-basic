package main

import (
	"fmt"
	"math"
)

func main() {
	result := commonCharacterCount("aabcc", "adcaa")
	fmt.Println(result)
}

func commonCharacterCount(s1 string, s2 string) int {
	mapS1 := make(map[int32]int)
	mapS2 := make(map[int32]int)
	for _, ch  := range s1 {
		if val, ok := mapS1[ch]; !ok {
			mapS1[ch] = 1
		} else {
			mapS1[ch] = val + 1
		}
	}

	for _, ch  := range s2 {
		if val, ok := mapS2[ch]; !ok {
			mapS2[ch] = 1
		} else {
			mapS2[ch] = val + 1
		}
	}
	result := 0
	var i int32
	for  i  = 97; i <= 122; i++ {
		result += int(math.Min(float64(mapS1[i]), float64(mapS2[i])))
	}

	return result
 }