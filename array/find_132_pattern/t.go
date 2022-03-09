package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(find132pattern([]int{1,2,3,4}))
	fmt.Println(find132pattern([]int{3,1,4,2}))
	fmt.Println(find132pattern([]int{-1,3,2,0}))
}

func find132pattern(nums []int) bool {
	stack := make([]int, 0, len(nums))
	ns := 0
	maxK := math.MinInt // 存放132模式中的次大值
	for i := len(nums)-1; i >= 0; i--{
		if nums[i] < maxK{
			return true
		}
		// 如果栈中有值，并且栈顶的值还小于当前元素
		// 此时，有了nums[j](nums[i]值) & nums[k](栈顶值)
		// 只要栈顶比当前元素小就出栈，保证了栈内元素是升序的，栈顶是【栈中】最大值
		for ns > 0 && stack[ns-1] < nums[i] {
			// 保存次大值
			maxK = stack[ns-1]
			stack = stack[:ns-1]
			ns--
		}
		stack = append(stack, nums[i])
		ns++
	}
	return false
}