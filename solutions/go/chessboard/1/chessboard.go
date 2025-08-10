package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	selectedRank := cb[rank]
	occupiedSquares := 0

	for _, v := range selectedRank {
		if v {
			occupiedSquares++
		}
	}

	return occupiedSquares
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	occupiedSquares := 0

	for _, r := range cb {
		fileIndex := file - 1
		if fileIndex >= 0 && fileIndex < len(r) {
			if r[file-1] {
				occupiedSquares++
			}
		}
	}

	return occupiedSquares
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	availableSquares := 0

	for _, r := range cb {
		availableSquares += len(r)
	}

	return availableSquares
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	occupiedSquares := 0

	for i := range cb {
		occupiedSquares += CountInRank(cb, i)
	}

	return occupiedSquares
}
