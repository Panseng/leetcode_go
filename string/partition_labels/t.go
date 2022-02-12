package main

import (
	"fmt"
	"sort"
)

func main()  {
	s := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(s))
}

func partitionLabels(s string) []int {
	n := len(s)
	s2ares := make(map[byte][]int, n)
	for i := 0; i < n; i++{
		if _, ok := s2ares[s[i]]; ok{
			s2ares[s[i]][1] = i
 		} else {
 			s2ares[s[i]] = []int{i,i}
		}
	}
	ares := [][]int{}
	for _, v := range s2ares{
		ares = append(ares, v)
	}
	if len(ares) > 1{
		ares = merge(ares)
	}

	ans := []int{}
	for _, v := range ares{
		ans = append(ans, v[1]-v[0]+1)
	}
	return ans
}

func merge(intervals [][]int) [][]int {
	iLen := len(intervals)
	if iLen < 2{
		return intervals
	}
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
