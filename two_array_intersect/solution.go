package main

import (
	"fmt"
	"sort"
	"time"
)

func main()  {
	nums1 := []int{4,9,5,1,0,4,2,3,7,5,9,8,11,21,33}
	nums2 := []int{9,4,9,8,4,5,6,9,7,10,12,13,11,11,15,16,18,19,12,35,6,89,74,2,0,12,34,55,64,12,34,15,19,18,16}
	times := 100000

	start := time.Now()
	for n := 0; n < times; n++ {
		intersectHash(nums1,nums2) // 先将数组进行排序
	}
	cost := time.Since(start)
	fmt.Printf("哈希法 cost: %s\n", cost)

	start = time.Now()
	for n := 0; n < times; n++ {
		intersectDoubleIndex(nums1,nums2) // 先将数组进行排序
	}
	cost = time.Since(start)
	fmt.Printf("双指针 cost: %s\n", cost)
}

// 简短化，同时空间使用更少，遍历次数更少
func intersectHash(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2){ // 用更短的数组作nums1，实现对更短数组记录hash表
		return intersectHash(nums2, nums1)
	}
	m := make(map[int]int) // m := map[int]int{}
	for _, num := range nums1{
		m[num]++
	}
	intersection := []int{}
	for _, num := range nums2{
		if m[num] > 0{
			intersection = append(intersection, num)
			m[num]--
		}
	}
	return intersection
}

// 双指针法 复杂度高了
func intersectDoubleIndex(nums1 []int, nums2 []int) []int  {
	sort.Ints(nums1)
	sort.Ints(nums2)
	len1, len2 := len(nums1), len(nums2)
	index1, index2 := 0,0
	intersetion := []int{}
	for index1 < len1 && index2 < len2{
		if nums1[index1] < nums2[index2]{
			index1++
		} else if nums1[index1] > nums2[index2]{
			index2++
		} else {
			intersetion = append(intersetion, nums1[index1])
			index1++
			index2++
		}
	}
	return intersetion
}