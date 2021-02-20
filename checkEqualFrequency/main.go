package main

func main() {

}

func checkEqualFrequency(inputArray []int) bool {
	mapInt := make(map[int]int)
	for _, v := range inputArray {
		if num, ok := mapInt[v]; !ok {
			mapInt[v] = 1
		} else {
			mapInt[v] = num + 1
		}
	}
	last := -1
	for _, element := range mapInt {
		if last == -1 {
			last = element
		}
		if element != last {
			return false
		}
	}

	return true
}