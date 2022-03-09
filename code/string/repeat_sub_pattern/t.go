package main

func main()  {
	repeatedSubstringPattern("abcabcabcabc")
	repeatedSubstringPattern("aabaabaa")
	repeatedSubstringPattern("aba")
}

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	for i := 1; i < n;i++{
		if n%i == 0{
			match := true
			for j := i; j<n;j++{
				if s[j] != s[j-i]{
					match = false
					break
				}
			}
			if match{
				return true
			}
		}
	}
	return false
}

func repeatedSubstringPattern2(s string) bool {
	s2 := s+s // 如果是由小段字符串重复组成，那么s2掐头去尾后，仍包含 s
	return kmp(s2[1:len(s2)-1], s) // 有重复，则，由2个该字符串组成的新字符串掐头去尾之后，仍包含原字符串
}

func kmp(haystack string, needle string) bool {
	n := len(needle)
	//if n == 0{ // needle > 0
	//	return 0
	//}
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
			return true
		}
	}
	return false
}
func getNext(s string, next []int)  {
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++{
		for j > 0 && s[i] != s[j]{
			j = next[j-1]
		}
		if s[i] == s[j]{
			j++
		}
		next[i] = j
	}
}
