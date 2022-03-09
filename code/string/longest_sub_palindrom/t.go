package main

import (
	"fmt"
)

func main()  {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	n := len(s)
	str := s[0:1]
	for i := 0; i < n;i++{
		ns := len(str)
		for j := i+ns+1; j <= n; j++{
			if isPali(s[i:j]){
				str = s[i:j]
			}
		}
	}
	return str
}

func isPali(s string) bool{
	for i := 0; i < len(s)/2; i++{
		if s[i] != s[len(s)-1-i]{
			return false
		}
	}
	return true
}

func longestPalindrome2(s string) string {
	if s == ""{
		return ""
	}
	start, end := 0,0
	for i := 0; i < len(s); i++{
		// 中心拓展
		start1, end1 := expandPali(s, i, i) // 回文数目为奇数
		start2, end2 := expandPali(s, i, i+1) // 回文数目为偶数
		if end1-start1 > end-start{
			start, end = start1, end1
		}
		if end2-start2 > end-start{
			start, end = start2, end2
		}
	}
	return s[start:end+1]
}
func expandPali(s string,start, end int) (int, int){
	// 由中心往两边同步拓展
	// 左侧-1，右侧+1，拓展后左右两侧相等构成回文则继续拓展
	// 否则返回缩窄左右两侧指针
	for ; start >= 0 && end < len(s) && s[start] == s[end]; start, end = start-1, end+1{}
	return start+1, end-1
}
