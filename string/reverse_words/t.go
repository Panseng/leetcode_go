package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main()  {
	s := "  Bob    Loves  Alice   "
	fmt.Println(reverseWords(s))
	fmt.Println(reverseWords2(s))

	//s = strings.Trim(s, " ")
	//fmt.Println(s)
	//reg := regexp.MustCompile(`\s{2,}`)
	//s = reg.ReplaceAllString(s, " ")
	//fmt.Println(s)
}

func reverseWords(s string) string {
	s = strings.Trim(s, " ")  // 排除两边空格
	reg := regexp.MustCompile(`\s{2,}`)
	s = reg.ReplaceAllString(s, " ") // 所有多空格转换为1个
	ans := strings.Split(s, " ")
	for left, right := 0, len(ans)-1; left < right; left, right = left+1, right-1{
		ans[left], ans[right] = ans[right], ans[left]
	}
	return strings.Join(ans, " ")
}

func reverseWords2(s string) string {
	ans := strings.Split(s, " ")
	for i, n := 0, len(ans); i < n; i++{
		if ans[i] == ""{
			if i == n-1{
				ans = ans[:i]
			} else {
				ans = append(ans[:i], ans[i+1:]...)
				i--
				n = len(ans)
			}
		}
	}
	for left, right := 0, len(ans)-1; left < right; left, right = left+1, right-1{
		ans[left], ans[right] = ans[right], ans[left]
	}
	return strings.Join(ans, " ")
}
