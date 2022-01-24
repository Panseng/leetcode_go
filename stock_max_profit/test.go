package main

import (
	"fmt"
	"time"
)

func main()  {
	nums := []int{7,1,5,3,6,4}
	times := 10000000

	start := time.Now()
	for n := 0; n < times; n++ {
		maxProfit1(nums) // 先将数组进行排序
	}
	cost := time.Since(start)
	fmt.Printf("1 cost: %s\n", cost)

	start = time.Now()
	for n := 0; n < times; n++ {
		maxProfit2(nums) // 先将数组进行排序
	}
	cost = time.Since(start)
	fmt.Printf("2 cost: %s\n", cost)
}

func maxProfit1(prices []int) int {
	max := 0
	for i :=0;i< len(prices)-1;i++{
		for n:=i+1;n<len(prices);n++{
			if prices[n]-prices[i]>max{
				max = prices[n]-prices[i]
			}
		}
	}
	return max
}

func maxProfit2(prices []int) int {
	max := 0
	for i :=0;i< len(prices)-1;i++{
		for n:=i+1;n<len(prices);n++{
			profit :=prices[n]-prices[i]
			if profit>max{
				max = profit
			}
		}
	}
	return max
}