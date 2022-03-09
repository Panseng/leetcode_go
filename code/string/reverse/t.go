package main

import "fmt"

func main() {
	s := "abcd"
	fmt.Println(reverseStr(s, 2))
}
func reverseStr(s string, k int) string {
	str := ""
	n := len(s)
	for i := 0; i < n; i += 2*k {
		end := i + k
		if end >= n {
			str = str + reverseString([]byte(s[i:]))
		} else {
			if end+k >= n{
				str = str + reverseString([]byte(s[i:end])) + s[end:]
			} else {
				str = str + reverseString([]byte(s[i:end])) + s[end:end+k]
			}
		}
	}
	return str
}

func reverseString(s []byte) string {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
	return string(s)
}
