package main

import (
	"fmt"
	"strings"
)

func main()  {
	pattern := "abba"
	str := "dog cat cat dog"
	fmt.Println(wordPattern2(pattern, str))
	pattern = "abba"
	str = "dog cat cat fish"
	fmt.Println(wordPattern2(pattern, str))
	pattern = "aaaa"
	str = "dog cat cat dog"
	fmt.Println(wordPattern2(pattern, str))
	pattern = "abba"
	str = "dog dog dog dog"
	fmt.Println(wordPattern2(pattern, str))
	pattern = "aba"
	str = "dog cat cat"
	fmt.Println(wordPattern2(pattern, str))
}

func wordPattern(pattern string, s string) bool {
	ps, ss := strings.Split(pattern, ""), strings.Split(s, " ")
	np, ns := len(ps), len(ss)
	if np != ns{
		return false
	}
	mp, ms, sp := make(map[string]int, np), make(map[string]int, np), make(map[string]string, np)
	for i := 0; i < np; i++{
		mp[ps[i]]++
		ms[ss[i]]++
		if mp[ps[i]] != ms[ss[i]]{
			return false
		}
		if _, ok := sp[ss[i]]; ok && sp[ss[i]] != ps[i]{
			return false
		}
		if _, ok := sp[ss[i]]; !ok{
			sp[ss[i]] = ps[i]
		}
	}
	return true
}

func wordPattern2(pattern string, s string) bool {
	ps, ss := strings.Split(pattern, ""), strings.Split(s, " ")
	np, ns := len(ps), len(ss)
	if np != ns{
		return false
	}
	p2s, s2p := make(map[string]string, np), make(map[string]string, np)
	for i := 0; i < np; i++{
		if _, ok := s2p[ss[i]]; ok && s2p[ss[i]] != ps[i]{
			return false
		}
		if _, ok := p2s[ps[i]]; ok && p2s[ps[i]] != ss[i]{
			return false
		}
		if _, ok := s2p[ss[i]]; !ok{
			s2p[ss[i]] = ps[i]
			p2s[ps[i]] = ss[i]
		}
	}
	return true
}

func wordPattern3(pattern string, s string) bool {
	n := len(pattern)
	words := strings.Split(s, " ") // 生成word值
	pByte2word := make(map[byte]string, n) // pattern单个字节的byte值对应的word值
	word2pByte := make(map[string]byte, n) // s单词string值对应pattern的byte值
	if len(words) != n{
		return false
	}
	for i, v := range words{
		if word2pByte[v] > 0 && word2pByte[v] != pattern[i] || pByte2word[pattern[i]] != "" && pByte2word[pattern[i]] != v{
			return false
		}
		if word2pByte[v] == 0{
			word2pByte[v] = pattern[i]
			pByte2word[pattern[i]] = v
		}
	}
	return true
}
