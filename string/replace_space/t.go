package main

import "fmt"

func main()  {
	fmt.Println(replaceSpace(" We are happy.  "))
}

func replaceSpace(s string) string {
	for i, n := 0, len(s); i < n; i++{
		if s[i] == " "[0]{
			if i == n-1{
				s = s[:i]+"%20"
				break
			}
			s = s[:i]+"%20"+s[i+1:]
			i += 2
			n = len(s)
		}
	}
	return s
}
