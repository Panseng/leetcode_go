package main

import "fmt"

func main()  {
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6}
	n := 3
	merge(nums1,m, nums2, n)
	fmt.Println(nums1)
}

// 双指针法
func merge(nums1 []int, m int, nums2 []int, n int)  {
	sorted := make([]int, 0, m+n)
	i1, i2 := 0,0
	for{
		if i1 == m{
			sorted = append(sorted, nums2[i2:]...)
			break
		}
		if i2 == n{
			sorted = append(sorted, nums1[i1:]...)
			break
		}
		if nums1[i1] < nums2[i2]{
			sorted = append(sorted, nums1[i1])
			i1++
		} else{
			sorted = append(sorted, nums2[i2])
			i2++
		}
	}
	copy(nums1,sorted)
}
