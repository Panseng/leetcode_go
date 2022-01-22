//题库链接：https://leetcode-cn.com/problems/contains-duplicate/solution/cun-zai-zhong-fu-yuan-su-by-leetcode-sol-iedd/
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	arrLen      = 1000  // 数组长度
	circleTimes = 10000 // 循环调用排序方法次数
)

func main() {
	start := time.Now()
	for n := 0; n < circleTimes; n++ {
		containsDuplicateSort(randArray()) // 先将数组进行排序
	}
	cost := time.Since(start)
	fmt.Printf("sort cost: %s\n", cost)

	start = time.Now()
	for n := 0; n < circleTimes; n++ {
		containsDuplicateHash(randArray()) // 不排序
	}
	cost = time.Since(start)
	fmt.Printf("hash cost: %s\n", cost)
}

// 生成数组
func randArray() []int {
	var arr [arrLen]int
	for i := 0; i < arrLen; i++ {
		arr[i] = rand.Int()
	}
	return arr[:]
}

// 自己的方法
func containsDuplicateHash(nums []int) bool {
	numExist := make(map[int]bool, len(nums))
	for _, v := range nums {
		if numExist[v] {
			return true
		}
		numExist[v] = true
	}
	return false
}

// 官方解法
func containsDuplicateSort(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
