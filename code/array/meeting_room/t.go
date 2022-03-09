package main

import (
	"fmt"
	"sort"
)

func main()  {
	intervals := [][]int{{1293,2986},{848,3846},{4284,5907},{4466,4781},{518,2918},{300,5870}}
	minMeetingRooms(intervals)
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i,j int)bool{ // 从小到大排序
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
	arranges := make([][]int, 0, len(intervals))
	arranges = append(arranges, intervals[0])
	for _, v := range intervals[1:]{
		fmt.Println(arranges)
		if arranges[0][1] > v[0]{
			arranges = append(arranges, v)
		} else {
			for _, a := range arranges{
				if a[1] <= v[0]{
					a[1] = v[1]
					break
				}
			}
		}
		sort.Slice(arranges, func(i, j int)bool{
			return arranges[i][1] < arranges[j][1]
		})
	}
	fmt.Println(arranges)
	return len(arranges)
}