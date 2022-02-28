package main

import "fmt"

func main()  {
	s := "abc"
	ss := [2]byte{s[1]}
	fmt.Println(ss)
}
