package main

import (
	"fmt"
)

func main() {
	//a :="bg"
	//b:="efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"

	//a :="a"
	//b:="b"

	a := "aac"
	b := "aab"
	c := "aA"
	for _, v := range c {
		fmt.Println(v)
	}
	fmt.Println(canConstruct(a, b))
	fmt.Println(canConstruc2(a, b))
}

func canConstruct(ransomNote string, magazine string) bool {
	rnHash := make(map[int32]int)
	mHash := make(map[int32]int)
	for _, v := range magazine {
		mHash[v]++
	}
	for _, v := range ransomNote {
		if mHash[v] == 0 {
			return false
		}
		rnHash[v]++
	}
	for _, v := range magazine {
		if rnHash[v] <= mHash[v] || rnHash[v] == 0 {
			continue
		}
		return false
	}
	return true
}

func canConstruc2(ransomNote string, magazine string) bool {
	mHash := make(map[int32]int)
	for _, v := range magazine {
		mHash[v]++
	}
	for _, v := range ransomNote {
		mHash[v]--
		if mHash[v] < 0 {
			return false
		}
	}
	return true
}
