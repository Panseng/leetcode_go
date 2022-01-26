package main

import "fmt"

func main() {
	a:=[][]int{{1,1,1},{1,0,1},{1,1,1}}
	fmt.Println(a)
	setZeroes(a)
	fmt.Println(a)
}
func setZeroes(matrix [][]int) {
	row, col := map[int]bool{}, map[int]bool{}
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, r := range matrix {
		for j, _ := range r {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}
