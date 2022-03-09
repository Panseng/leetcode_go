package main

import "fmt"

func main()  {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
	fmt.Println(minRemoveToMakeValid("a)b(c)d"))
	fmt.Println(minRemoveToMakeValid("))(("))
}

func minRemoveToMakeValid(s string) string {
	leftBs := []int{}
	ln := 0
	sn := len(s)
	for i := 0; i < sn; i++{
		if s[i] == '('{
			leftBs = append(leftBs, i)
			ln++
		}
		if s[i] == ')'{
			if ln == 0{
				s = removeIndex(s, sn, i)
				sn = len(s)
				i--
				continue
			}
			leftBs = leftBs[:ln-1]
			ln--
		}
	}
	for i := ln-1; i >= 0; i--{
		s = removeIndex(s, sn, leftBs[i])
		sn = len(s)
	}
	return s
}

func removeIndex (s string, n, index int) string{
	if index == n-1{
		return s[:n-1]
	}
	return s[:index]+s[index+1:]
}
