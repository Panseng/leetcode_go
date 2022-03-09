package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := "123"
	fmt.Println(a[1])
	fmt.Println(int(a[1]))
	num1 := "0"
	num2 := "0"
	//fmt.Println(strings.Split(num1,""), strings.Split(num2, ""))
	fmt.Println(addStrings(num1, num2))
}

func addStrings(num1 string, num2 string) string {
	strArr := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	numsStr := "0123456789"

	ns1 := strings.Split(num1, "")
	ns2 := strings.Split(num2, "")
	n1, n2 := len(ns1), len(ns2)
	// 补位
	if n1 > n2 {
		ns2 = fillZero(ns2, n1-n2)
	} else if n1 < n2 {
		ns1 = fillZero(ns1, n2-n1)
	}

	ans := []string{}
	n := len(ns1)
	carry := 0 // 进位索引
	for i := n - 1; i >= 0; i-- {
		addIndex := carry + strings.Index(numsStr, ns1[i]) + strings.Index(numsStr, ns2[i])
		ans = append([]string{strArr[addIndex%10]}, ans...)
		carry = addIndex / 10
	}

	if carry > 0 {
		ans = append([]string{strArr[carry]}, ans...)
	}

	return strings.Join(ans, "")
}

func fillZero(strArr []string, n int) []string {
	zeros := make([]string, n)
	for i := range zeros {
		zeros[i] = "0"
	}
	return append(zeros, strArr...)
}

// 优化
func addStrings2(num1 string, num2 string) string {
	ans := ""
	n1 := len(num1)
	n2 := len(num2)
	if n1 > n2 {
		num2 = fillZero2(n1-n2) + num2
	} else if n1 < n2 {
		num1 = fillZero2(n2-n1) + num1
	}
	n1 = len(num1)
	carry := 0
	for i := n1 - 1; i >= 0; i-- {
		res := carry + int(num1[i]-'0') + int(num2[i]-'0')
		ans = strconv.Itoa(res%10) + ans
		carry = res / 10
	}
	if carry > 0 {
		ans = strconv.Itoa(carry) + ans
	}
	return ans
}

func fillZero2(n int) string {
	zeroStr := ""
	for i := 0; i < n; i++ {
		zeroStr += "0"
	}
	return zeroStr
}
