package main

import "fmt"

func main()  {
	fmt.Println(largestRectangleArea([]int{2,1,5,6,2,3}))
	fmt.Println(largestRectangleArea([]int{2,4}))
	fmt.Println(largestRectangleArea2([]int{2,1,5,6,2,3}))
	fmt.Println(largestRectangleArea2([]int{2,4}))
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	stack := make([]int, 0, n) // 栈，记录可取值左侧下标
	ns := 0 // 栈长度
	for i := 0; i < n; i++{
		// [2,6,1,6,2,8,4,7,9]
		// i = 0时，stack = []，表示 i = 0时，可取值左侧下标(不包含)为-1
		// i = 1时，循环后stack = [0]，表示 i = 1时，可取值左侧下标(不包含)为0
		// i = 2时，循环后stack = []，表示 i = 2时，可取值左侧下标(不包含)为-1
		// 只保留小于自身的下标
		for ns > 0 && heights[stack[ns-1]] >= heights[i]{
			stack = stack[:ns-1]
			ns--
		}
		if ns == 0{
			left[i] = -1
		} else {
			left[i] = stack[ns-1]
		}
		stack = append(stack, i)
		ns++
	}

	stack = make([]int, 0, n)
	ns = 0
	// [2,6,1,6,2,8,4,7,9]
	// i = 8时，stack = []，表示 i = 8时，可取值右侧下标(不包含)为9
	// i = 7时，循环后stack = []，表示 i = 7时，可取值右侧下标(不包含)为9
	// i = 6时，循环后stack = []，表示 i = 6时，可取值右侧下标(不包含)为9
	// i = 5时，循环后stack = [6]，表示 i = 6时，可取值右侧下标(不包含)为9
	for i := n-1; i >= 0; i-- {
		for ns > 0 && heights[stack[ns-1]] >= heights[i]{
			stack = stack[:ns-1]
			ns--
		}
		if ns == 0{
			right[i] = n
		} else {
			right[i] = stack[ns-1]
		}
		stack = append(stack, i)
		ns++
	}

	ans := 0
	for i := 0; i < n; i++{
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func largestRectangleArea2(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	for i := range right{
		right[i] = n
	}

	stack := make([]int, 0, n) // 栈，记录可取值左侧下标
	ns := 0 // 栈长度
	for i := 0; i < n; i++{
		// [2,6,1,6,2,8,4,7,9]
		// i = 0时，stack = []，表示 i = 0时，可取值左侧下标(不包含)为-1
		// i = 1时，循环后stack = [0]，表示 i = 1时，可取值左侧下标(不包含)为0
		// i = 2时，循环后stack = []，表示 i = 2时，可取值左侧下标(不包含)为-1
		// 只保留小于自身的下标
		for ns > 0 && heights[stack[ns-1]] >= heights[i]{
			right[stack[ns-1]] = i
			stack = stack[:ns-1]
			ns--
		}
		if ns == 0{
			left[i] = -1
		} else {
			left[i] = stack[ns-1]
		}
		stack = append(stack, i)
		ns++
	}

	ans := 0
	for i := 0; i < n; i++{
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}
	return ans
}

func max(a, b int)  int{
	if a > b {
		return a
	}
	return b
}