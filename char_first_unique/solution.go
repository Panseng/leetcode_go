package main

import (
	"fmt"
)

func main() {
	s := "lovelove"
	//s := "love leetcode"
	//fmt.Println(firstUniqChar3(s))
	fmt.Println(firstUniqChar4(s))
	fmt.Println(firstUniqChar5(s))
}

func firstUniqChar3(s string) int {
	a := [26]int{} // 26个字母
	for _, v := range s {
		a[v-'a']++ //
	}
	for i, v := range s {
		if a[v-'a'] == 1 {
			return i
		}
	}
	return -1
}

type loc struct {
	ch  int32 // 字符串被range之后的值
	pos int
}

func firstUniqChar4(s string) int {
	var que []loc
	hashTable := make(map[int32]int)
	for i, v := range s {
		hashTable[v]++
		if hashTable[v] > 1 {
			for len(que) > 0 && hashTable[que[0].ch] > 1 {
				que = que[1:]
			}
		} else {
			que = append(que, loc{v, i})
		}
	}
	if len(que) > 0 {
		return que[0].pos
	}
	return -1
}

func firstUniqChar5(s string) int {
	var que []loc
	hashTable := make(map[int32]int)
	for i, v := range s {
		hashTable[v]++
		if hashTable[v] == 1 {
			que = append(que, loc{v, i})
		}
	}
	for len(que) > 0 && hashTable[que[0].ch] > 1 {
		que = que[1:] // 先进先出
	}
	if len(que) > 0 {
		return que[0].pos
	}
	return -1
}
