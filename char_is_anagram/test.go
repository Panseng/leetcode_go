package main

import (
	"fmt"
	"sort"
)

func main()  {
	//s, t:="anagram","nagaram"
	s, t:="rat","car"
	fmt.Println(isAnagram(s,t))
	fmt.Println(isAnagram2(s,t))
	fmt.Println(isAnagram3(s,t))
}

func isAnagram(s string, t string) bool {
	sLen:= len(s)
	if sLen != len(t){
		return false
	}
	sHash, tHash := make(map[byte]int), make(map[byte]int)
	for i:=0;i<sLen;i++{
		sHash[s[i]]++
		tHash[t[i]]++
	}
	for k,_:=range sHash{
		if sHash[k] != tHash[k]{
			return false
		}
	}
	return true
}


func isAnagram2(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	sHash := make(map[rune]int)
	for _,v:=range s{
		sHash[v]++
	}
	for _,v:=range t{
		sHash[v]--
		if sHash[v] < 0{
			return false
		}
	}
	return true
}
func isAnagram3(s string, t string) bool {
	ls1,ls2 := []byte(s),[]byte(t)
	sort.Slice(ls1, func(i, j int) bool {
		return ls1[i]<ls1[j]
	})
	sort.Slice(ls2, func(i, j int) bool {
		return ls2[i]<ls2[j]
	})
	return string(ls1) == string(ls2)
}
