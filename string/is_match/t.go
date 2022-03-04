package main

import (
	"fmt"
)

func main()  {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "*"))
	fmt.Println(isMatch("cb", "?a"))
	fmt.Println(isMatch("adceb", "*a*b"))
	fmt.Println(isMatch("acdcb", "a*c?b"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("ab", "?*"))
}

func isMatch(s string, p string) bool {
	ns, np := len(s), len(p)
	dp := make([][]bool, ns+1) // 0 位保存长度为0的结果， ns位保存尾部匹配结果
	for i := range dp{
		dp[i] = make([]bool, np+1)
	}

	dp[0][0] = true // 长度均为 0 时，返回真
	for j := 1; j <= np; j++{ // 对于 s 为 0的情况，如果p的前j项为*，则j+1均为真
		if p[j-1] == '*'{
			dp[0][j] = true
		} else{
			break
		}
	}

	for i := 1; i <= ns; i++{
		for j := 1; j <= np; j++{
			if p[j-1] == '*'{ // 对于 * ，我们可以用或不用
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			} else if p[j-1] == '?' || p[j-1] == s[i-1]{
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[ns][np]
}