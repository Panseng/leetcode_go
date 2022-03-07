package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
描述
写出一个程序，接受一个由字母、数字和空格组成的字符串，和一个字符，然后输出输入字符串中该字符的出现次数。（不区分大小写字母）
数据范围：1≤n≤1000
输入描述：第一行输入一个由字母和数字以及空格组成的字符串，第二行输入一个字符。
输出描述：输出输入字符串中含有该字符的个数。（不区分大小写字母）

输入：
   ABCabc
   A
输出：2
 */
func main()  {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	input.Scan()
	key := input.Text()

	str = strings.ToLower(str)
	key = strings.ToLower(key)

	k := key[0]
	n := 0
	for i := range str{
		if k == str[i]{
			n++
		}
	}
	fmt.Printf("%d\n", n)
}
