package main

import "fmt"

func main()  {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("", ""))
	strStr2("", "abcabceabc")
}

func strStr(haystack string, needle string) int {
	n1, n2 := len(haystack), len(needle)
	if n2 == 0{
		return 0
	}
	if n1 == 0 || n1<n2 || (n1 == n2 && haystack != needle){
		return -1
	}
	for i := 0; i <= n1-n2; i++{
		if haystack[i:i+n2] == needle{
			return i
		}
	}
	return -1
}

func strStr2(haystack string, needle string) int {
	n := len(needle)
	if n == 0{
		return 0
	}
	next := make([]int, n)
	getNext(needle, next)
	j := 0;
	for i := 0; i < len(haystack); i++{
		for j > 0 && haystack[i] != needle[j]{
			j = next[j-1] // 根据next回退
		}
		if haystack[i] == needle[j]{
			j++
		}
		if j == n{
			return i-n+1
		}
	}
	return -1
}

func getNext(s string, next []int){
	j := 0;
	next[0] = j
	for i := 1; i < len(s); i++{
		for j > 0 && s[j] != s[i]{
			j = next[j-1]
		}
		if s[i]==s[j]{
			j++
		}
		next[i] = j
	}
}
