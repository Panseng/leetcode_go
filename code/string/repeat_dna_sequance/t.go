package main

import "fmt"

func main()  {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	//s := "AAAAAAAAAAAAA"
	//fmt.Println(s[0:9])
	fmt.Println(findRepeatedDnaSequences(s))
}

func findRepeatedDnaSequences(s string) []string {
	ans := []string{}
	ansMap := make(map[string]int)
	for i:=0; i < len(s)-10; i++{
		rep := s[i:i+10]
		ansMap[rep]++
		if ansMap[rep] == 2{
			ans = append(ans, rep)
		}
	}
	return ans
}
