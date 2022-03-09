package main

import (
	"fmt"
)

func main() {
	a := make([][]byte, 9)
	a[0] = []byte{'8', '3', '.', '.', '7', '.', '.', '.', '.'}
	a[1] = []byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'}
	a[2] = []byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'}
	a[3] = []byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'}
	a[4] = []byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'}
	a[5] = []byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'}
	a[6] = []byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'}
	a[7] = []byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'}
	a[8] = []byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}

	fmt.Println(isValidSudoku(a))
}

func isValidSudoku(board [][]byte) bool {
	var rows, cols [9][9]int
	var ceils [3][3][9]int
	for i, r := range board {
		for j, v := range r {
			if v == '.' {
				continue
			}
			index := v - '1' // 将byte转换为int，顺便-1作为index下标
			rows[i][index]++
			cols[j][index]++
			ceils[i/3][j/3][index]++
			if rows[i][index] > 1 || cols[j][index] > 1 || ceils[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
