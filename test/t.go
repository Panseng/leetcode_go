package main

import (
	"fmt"
	"regexp"
)

func main()  {
	s := "abc"
	ss := [2]byte{s[1]}
	fmt.Println(ss)

	s2 := "e:\\code\\golang\\leetcode_go\\notes\\base.md"
	if found, _ := regexp.MatchString("^[a-z]:\\\\", s2); found{
		fmt.Println("regexp absolute addr right")
	} else {
		fmt.Println("regexp absolute addr wrong")
	}
}
