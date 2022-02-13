package main

import (
	"fmt"
	"strconv"
)

func main()  {
	n1 := "2"
	n2 := "3"
	fmt.Println(multiply(n1, n2))
}

// 大数不可行
// "498828660196"
// "840477629533"
func multiply(num1 string, num2 string) string {
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)
	n2 = n1*n2
	return strconv.Itoa(n2)
}

func multiply2(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0"{
		return "0"
	}
	n, m := len(num1), len(num2)
	multiArr := make([]int,m+n)
	for i := n-1; i >= 0; i--{
		x := int(num1[i]-'0')
		for j := m-1; j >= 0; j--{
			y := int(num2[j]-'0')
			multiArr[i+j+1] += x*y
		}
	}
	for i := n+m-1; i>0; i--{
		multiArr[i-1] += multiArr[i]/10
		multiArr[i] %= 10
	}
	ind := 0
	str := ""
	if multiArr[0] == 0{
		ind = 1
	}
	for ; ind < n+m; ind++{
		str += strconv.Itoa(multiArr[ind])
	}
	return str
}
