package main

import "fmt"

func main()  {
	a := []int{1,2,3,4}
	fmt.Println(productExceptSelf(a))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n < 2{
		return nums
	}
	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; i++{ // 先获取左侧乘积
		ans[i] = nums[i-1]*ans[i-1]
	}
	r := 1
	for i := n-1; i >= 0; i--{
		ans[i] = ans[i] * r
		r *= nums[i]
	}
	return ans
}
