package main

import "fmt"

func main() {
	fmt.Print(chessBoardCellColor("A1", "B3"))
}

func chessBoardCellColor(cell1 string, cell2 string) bool {

	if (cell1[0] + cell1[1] + cell2[0] + cell2[1])%2 == 0 {
		return true
	}
	return false
}