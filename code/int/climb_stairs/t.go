package main

import "fmt"

func main()  {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairs(6))
}
func climbStairs(n int) int {
	if n == 1{
		return 1
	}
	if n == 2{
		return 2
	}
	steps := make([]int, n+1)
	steps[1], steps[2] = 1, 2 // 边界

	for i := 3; i <= n; i++ {
		steps[i] = steps[i-2] + steps[i-1] // 状态转移
	}

	return steps[n] // 最优子结构
}
