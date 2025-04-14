package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	fileSquares := cb[file]
	var sumSquares int
	for _, x := range fileSquares {
		if x {
			sumSquares++
		}
	}
	return sumSquares
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var sumSquares int
	for _,v := range cb {
		if rank-1 >= 0 && rank-1 < len(v) {
			if v[rank-1] {
				sumSquares+=1
			}
		}
	}
	return sumSquares
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var countRanks int
	for _,v := range cb {
		countRanks+=len(v)
	}
	return countRanks
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var occupied int 
	for _,v := range cb {
		for _, x := range v {
			if x {
				occupied++
			}
		}
	}
	return occupied
}
