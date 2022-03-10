package main

func main() {

}
func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	sums := make([]int, n+1)
	// 获取前n项和，
	for i, v := range nums {
		sums[i+1] = sums[i] + v
	}

	ans := n + 1
	// 滑动窗口
	// 递增元素的下标
	queue := make([]int, 0, n+1)
	nq := 0
	for i, v := range sums {
		// 确保区间内递增
		for nq > 0 && v <= sums[queue[nq-1]] {
			queue = queue[:nq-1]
			nq--
		}
		for nq > 0 && v-sums[queue[0]] >= k {
			if ans > i-queue[0] {
				ans = i - queue[0]
			}
			queue = queue[1:]
			nq--
		}
		queue = append(queue, i)
		nq++
	}
	if ans == n+1 {
		return -1
	}
	return ans
}

// O(n^2)超时
func shortestSubarray2(nums []int, k int) int {
	n := len(nums)
	sums := make([]int, n)
	for i, v := range nums {
		if v >= k {
			return 1
		}
		if i == 0 {
			sums[i] = v
			continue
		}
		sums[i] = sums[i-1] + v
	}
	if n < 2 { // 不存在
		return -1
	}
	for i := 2; i <= n; i++ {
		for j := 0; j+i-1 < n; j++ {
			left := 0
			if j > 0 {
				left = sums[j-1]
			}
			if k <= sums[j+i-1]-left {
				return i
			}
		}
	}
	return -1
}
