package main

import (
	"fmt"
	"strconv"
)

func main()  {
	fmt.Println(isHappy(6))
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
	fmt.Println(isHappy2(19))
	fmt.Println(isHappy2(2))
	fmt.Println(isHappy3(19))
	fmt.Println(isHappy3(2))
}

func isHappy(n int) bool {
	if n == 1{
		return true
	}
	s := strconv.Itoa(n)
	intMap := make(map[int]struct{})
	intMap[n] = struct{}{}
	for {
		n = 0
		for _,v := range s{
			n += int(v-'0')*int(v-'0')
		}
		if n == 1{
			return true
		}
		if _, ok := intMap[n]; ok{
			return false
		}
		intMap[n] = struct{}{}
		s = strconv.Itoa(n)
	}
}

func isHappy2(n int) bool {
	if n == 1{
		return true
	}
	intMap := make(map[int]bool)
	for ; n != 1 && !intMap[n];n, intMap[n] = step(n), true{}
	return n == 1
}

func isHappy3(n int) bool {
	if n == 1{
		return true
	}
	slow, fast := n, step(n)
	for fast != 1 && fast != slow {
		slow, fast = step(slow), step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0{
		sum += (n%10) * (n%10)
		n = n/10
	}
	return sum
}
