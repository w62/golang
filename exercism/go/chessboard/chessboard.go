package chessboard
import (
	"fmt"
)

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {

	result :=0

	if _, ok := cb[file]; ok {


		for i:=0; i<len(cb[file]); i++ {

			if cb[file][i] {
				result ++
				fmt.Println("occupied.")
			}
		}

	} else {

	fmt.Println("file", file, " not found in chessboard.")
	return 0
	}
//	fmt.Println(cb, file)
	fmt.Println(file, "len(file)=", len(file), "Chessboard file", cb[file], "result=", result)
	return result
	panic("Please implement CountInFile()")
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {

	result := 0

	myrank := rank - 1

	if myrank >=0 && myrank <8 {

	for _, file := range cb {
		fmt.Println("file=",file, " rank=", myrank)
		if file[myrank] {
			result ++
		}
		fmt.Println(file)
	}
}

	return result
	panic("Please implement CountInRank()")
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb) * 8
	panic("Please implement CountAll()")
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {

	result := 0
	for i := 1; i<=8; i++ {
		result += CountInRank(cb, i)
	}

	return result
	panic("Please implement CountOccupied()")
}
