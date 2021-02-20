package main

func main() {

}

func bishopAndPawn(bishop string, pawn string) bool {
	if (bishop[0] + pawn[1] == bishop[1] + pawn[0] || bishop[0] + bishop[1] == pawn[0] + pawn[1]) && bishop[0] != pawn[0] {
		return true
	}
	return false
}