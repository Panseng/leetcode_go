package main

import (
	"fmt"
	"time"
)

func main()  {
	times := 10000000
	nums := []int{2,7,11,15}

	start := time.Now()
	for n := 0; n < times; n++ {
		twoSum(nums, 26) // 先将数组进行排序
	}
	cost := time.Since(start)
	fmt.Printf("暴力 cost: %s\n", cost)

	start = time.Now()
	for n := 0; n < times; n++ {
		twoSumHash(nums, 26) // 先将数组进行排序
	}
	cost = time.Since(start)
	fmt.Printf("哈希 cost: %s\n", cost)
}

func twoSum(nums []int, target int) []int {
	for left:=0;left <= len(nums)-2;left++{
		for right :=left+1;right<len(nums);right++{
			if nums[left]+nums[right]==target{
				return []int{left,right}
			}
		}
	}
	return []int{-1,-1}
}

func twoSumHash(nums []int, target int) []int {
	hashTable := map[int]int {}
	for i, x :=range nums{
		if p, ok := hashTable[target-x];ok{
			return []int {p,i}
		}
		hashTable[x] = i
	}
	return nil
}