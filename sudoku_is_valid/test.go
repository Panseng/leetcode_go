package main

import "fmt"

func main() {
	a := make([][]byte, 9)
	a[0] = []byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'}
	a[1] = []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}
	a[2] = []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}
	a[3] = []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}
	a[4] = []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}
	a[5] = []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}
	a[6] = []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}
	a[7] = []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}
	a[8] = []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}

	fmt.Println(a)
	fmt.Println(isValidSudokuSelf(a))
}

func isValidSudokuSelf(board [][]byte) bool {
	return isValidSudoku2(board, true)
}

func isValidSudoku2(board [][]byte, isRoot bool) bool {
	for _, v := range board {
		if isRepeat(v) {
			return false
		}
	}
	if !isRoot {
		return true
	}

	rows := make([][]byte, 9)
	ceils := make([][]byte, 9)
	for i := 0; i < 9*9; i++ {
		rows[i%9] = append(rows[i%9], board[i/9][i%9])
		switch i {
		case 0, 1, 2, 9, 10, 11, 18, 19, 20:
			ceils[0] = append(ceils[0], board[i/9][i%9])
		case 3, 4, 5, 12, 13, 14, 21, 22, 23:
			ceils[1] = append(ceils[1], board[i/9][i%9])
		case 6, 7, 8, 15, 16, 17, 24, 25, 26:
			ceils[2] = append(ceils[2], board[i/9][i%9])
		case 27, 28, 29, 36, 37, 38, 45, 46, 47:
			ceils[3] = append(ceils[3], board[i/9][i%9])
		case 30, 31, 32, 39, 40, 41, 48, 49, 50:
			ceils[4] = append(ceils[4], board[i/9][i%9])
		case 33, 34, 35, 42, 43, 44, 51, 52, 53:
			ceils[5] = append(ceils[5], board[i/9][i%9])
		case 54, 55, 56, 63, 64, 65, 72, 73, 74:
			ceils[6] = append(ceils[6], board[i/9][i%9])
		case 57, 58, 59, 66, 67, 68, 75, 76, 77:
			ceils[7] = append(ceils[7], board[i/9][i%9])
		case 60, 61, 62, 69, 70, 71, 78, 79, 80:
			ceils[8] = append(ceils[8], board[i/9][i%9])
		}
	}
	if !isValidSudoku2(rows, false) {
		return false
	}
	return isValidSudoku2(ceils, false)
}

func isRepeat(bytes []byte) bool {
	byteExist := make(map[byte]bool, len(bytes))
	for _, v := range bytes {
		if v != '.' { // 非空
			if byteExist[v] {
				return true
			}
			byteExist[v] = true
		}
	}
	return false
}
