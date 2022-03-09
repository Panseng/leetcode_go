package main

import "fmt"

func main()  {
	nums := []int{2,3,1,2,4,3}
	target := 7
	fmt.Println(minSubArrayLen(target, nums))
}

func minSubArrayLen(target int, nums []int) int {
	result := len(nums) + 1
	left := 0
	sum := 0
	for right, v := range nums{ // 不断将 right 指针往右边递增
		sum += v
		for sum >= target{ // 不断将 left 指针往右边压缩
			if right-left+1 < result{
				result = right-left+1
			}
			sum -= nums[left]
			left++
		}
	}
	if result == len(nums)+1{
		return 0
	}
	return result
}

func minSubArrayLen2(target int, nums []int) int {
	n := len(nums)
	result := n+1
	for i, v := range nums{
		if v >= target{
			return 1
		}
		sum := v
		for j := i+1; j < n; j++{
			sum += nums[j]
			if sum >= target{
				if j-i+1 < result{
					result = j-i+1
				}
			}
		}
	}
	if result == n+1{
		return 0
	}
	return result
}
