package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	var result int
    for _, r := range cb[rank] {
        if r {
            result++
        }
    }
    return result
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
    if file <= 0 || file >= 9 {
        return 0
    }
	var result int
    for _, v := range cb {
        if v[file-1] {
            result++
        }
    }
    return result
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	var result int
    for _, v := range cb {
        result += len(v)
    }
    return result
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var result int
    for _, v := range cb {
        for _, r := range v {
            if r {
                result++
            }
        }
    }
    return result
}
