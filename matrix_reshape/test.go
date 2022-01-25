package main

import (
	"fmt"
)

func main()  {

	mat := [][]int{{1,2},{3,4}}
	r,c:=4,1
	fmt.Println(matrixReshape(mat,r,c))
}

// 二维数组一维表示
func matrixReshape(mat [][]int, r int, c int) [][]int {
	m,n:=len(mat), len(mat[0])
	if m*n !=r*c{
		return mat
	}
	newMat := make([][]int, r)
	for i := range newMat{
		newMat[i] = make([]int, c)
	}
	for i:=0;i<m*n;i++{
		newMat[i/c][i%c] = mat[i/n][i%n]
	}
	return newMat
}

func matrixReshape2(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n / r != c{
		return mat
	}
	var nums []int
	for _,v:=range mat{
		nums = append(nums, v...)
	}
	newMat := make([][]int,r)
	for i,j:=0,0; i< m*n;i=i+c{
		newMat[j] = nums[i:i+c]
		j++
	}
	return newMat
}

