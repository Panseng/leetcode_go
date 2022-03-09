package main

import (
	"fmt"
	"sort"
)

func main()  {
	nums := []int{-4,-1,0,3,10}
	//fmt.Println(sortedSquares(nums))
	fmt.Println(sortedSquares2(nums))
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for l, r, i := 0, n-1, n-1; r >= l; i--{
		left, right := nums[l]*nums[l], nums[r]*nums[r]
		if left > right{
			ans[i] = left
			l++
		} else {
			ans[i] = right
			r--
		}
	}
	return ans
}

func sortedSquares2(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i, v := range nums{
		ans[i] = v*v
	}
	sort.Ints(ans)
	return ans
}