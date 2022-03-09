package main

import "fmt"

func main()  {
	//nums := []int{3,2,2,3}
	//n := 3

	nums := []int{0,1,2,2,3,0,4,2}
	n := 2

	//nums := []int{2,2,2}
	//n := 2

	//a := append(nums[:1], nums[2:]...)

	fmt.Println(removeElement(nums, n))
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums);{
		if nums[i] == val{
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
	}
	return len(nums)
}

func removeElement2(nums []int, val int) int {
	left := 0
	for i, v := range nums { // v 即 nums[right]
		if v != val {
			nums[left], nums[i] = v, nums[left] // 这样会交换多次，不划算
			left++
		}
	}
	return left
}