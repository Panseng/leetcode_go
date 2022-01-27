package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "loveleetcode"
	ss := strings.Split(s, "")
	aa := []int{1, 2, 3}
	fmt.Println(ss, aa)
	a := make(map[int]int)
	fmt.Println(a[0])
	fmt.Println(firstUniqChar2(s))
}

func firstUniqChar1(s string) int {
	index := -1
	ss := strings.Split(s, "")
	hasTab := make(map[string]bool)
	repeatTab := make(map[string]bool)
	for _, v := range ss {
		if hasTab[v] {
			repeatTab[v] = true
		} else {
			hasTab[v] = true
		}
	}
	for i := len(ss) - 1; i >= 0; i-- {
		if repeatTab[ss[i]] {
			continue
		}
		index = i
	}
	return index
}

func firstUniqChar2(s string) int {
	a := make(map[int32]int)
	for _, v := range s {
		a[v]++
	}
	for i, v := range s {
		if a[v] == 1 {
			return i
		}
	}
	return -1
}
