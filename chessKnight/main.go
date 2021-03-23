package main

import "fmt"

func main() {
	count := chessKnight("c2")
	fmt.Println(count)
}

func chessKnight(Cell string) int {
	count := 0

	if nextX := Cell[0] + 2; nextX <= 'h' {
		if nextY := Cell[1] + 1; nextY <= '8' {
			count++
		}
		if nextY := Cell[1] - 1; nextY >= '1' {
			count++
		}
	}
	if nextX := Cell[0] - 2; nextX >= 'a' {
		if nextY := Cell[1] + 1; nextY <= '8' {
			count++
		}
		if nextY := Cell[1] - 1; nextY >= '1' {
			count++
		}
	}

	if nextY := Cell[1] + 2; nextY <= '8' {
		if nextX := Cell[0] + 1; nextX <= 'h' {
			count++
		}

		if nextX := Cell[0] - 1; nextX >= 'a' {
			count++
		}
	}

	if nextY := Cell[1] - 2; nextY >= '1' {
		if nextX := Cell[0] + 1; nextX <= 'h' {
			count++
		}

		if nextX := Cell[0] - 1; nextX >= 'a' {
			count++
		}
	}
	return count
}
