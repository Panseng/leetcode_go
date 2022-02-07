package main

import (
	"fmt"
	"sort"
)

func main()  {
	a := [][]int{{1,4},{0,0}}
	fmt.Println(a)
	fmt.Println(merge(a))
}
func merge(intervals [][]int) [][]int {
	iLen := len(intervals)
	if iLen < 2{
		return intervals
	}
	// 二维数组排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]//按照每行的第一个元素排序
	})
	ans := [][]int{intervals[0]}
	for i,v := range intervals{
		if i == 0{
			continue
		}
		aLen := len(ans)
		aLast := ans[aLen-1]
		if aLast[1] >= v[0]{
			aLast[0] = min(aLast[0], v[0])
			aLast[1] = max(aLast[1], v[1])
			continue
		}
		ans = append(ans, v)
	}
	return ans
}

func min(a, b int) int{
	if a < b{
		return a
	}
	return b
}

func max(a, b int) int{
	if a < b{
		return b
	}
	return a
}
