package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(shortestPalindrome("aacecaaa"))
	fmt.Println(shortestPalindrome("abcd"))
}

func shortestPalindrome(s string) string {
	n := len(s)
	b := strings.Split(s, "")
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		b[l], b[r] = b[r], b[l]
	}
	pattern := s + "#" + strings.Join(b, "")
	n = n*2 + 1
	next := make([]int, n)
	getNext(pattern, next)
	return pattern[len(s)+1:n-next[n-1]] + s
}

func getNext(s string, next []int) {
	j := 0
	next[0] = j
	n := len(s)
	for i := 1; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}

func shortestPalindrome0(s string) string {
	l := len(s)
	rs := reverserStr(s)
	i := 0
	for ; i < l; i++ {
		//- 关键代码：反转后的尾部 同 反转前的头部，最大的重叠区间
		//- 那么，不相同的反转头部，为最小字符串
		//- 转换为kmp，则 s + "#" + rs，求 next
		//- 则最后一个字母的前缀和，对应与 s 的头部重叠区间
		if rs[i:] == s[:l-i] {
			break
		}
	}
	return rs[:i] + s
}
func reverserStr(x string) string {
	res := strings.Builder{}
	for i := len(x) - 1; i >= 0; i-- {
		res.WriteByte(x[i])
	}
	return res.String()
}
