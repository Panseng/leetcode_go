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

	var i interface{} = 77
	if value, ok := i.(int); ok {
		fmt.Println("类型匹配整型：%d\n", value)
	} else if value, ok := i.(string); ok {
		fmt.Printf("类型匹配字符串:%s\n", value)
	}
	a := 5
	c := float64(a)/2
	fmt.Println(c)
}

