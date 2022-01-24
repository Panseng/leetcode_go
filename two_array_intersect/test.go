package main

import "fmt"

func main()  {
	nums1 := []int{4,9,5}
	nums2 := []int{9,4,9,8,4}
	fmt.Println(intersect(nums1,nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
	nums1Map, nums2Map := make(map[int]int), make(map[int]int)
	getNumMap(nums1, nums1Map)
	getNumMap(nums2, nums2Map)
	if len(nums1) < len(nums2){
		return getInters(nums1, nums1Map, nums2Map)
	} else {
		return getInters(nums2, nums1Map, nums2Map)
	}
}

func getNumMap(nums []int, numMap map[int]int){
	for _, v := range nums{
		numMap[v]++
	}
}

func getInters(nums []int, nums1Map map[int]int, nums2Map map[int]int) []int {
	var inters []int
	hasAppend := make(map[int]bool)
	for _, v := range nums{
		if _, ok := hasAppend[v];ok{
			continue
		}
		hasAppend[v] = true
		v1, ok1 := nums1Map[v]
		v2, ok2 := nums2Map[v]
		if ok1 && ok2{
			if v1 < v2{
				appendLen(&inters, v, v1)
			} else {
				appendLen(&inters, v, v2)
			}
		}
	}
	return inters
}

func appendLen(nums *[]int, num int, len int)  {
	for i :=0; i < len; i++{
		*nums = append(*nums, num)
	}
}
