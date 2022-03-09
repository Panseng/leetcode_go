package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
描述
计算字符串最后一个单词的长度，单词以空格隔开，字符串长度小于5000。（注：字符串末尾不以空格为结尾）
输入描述：
输入一行，代表要计算的字符串，非空，长度小于5000。
输出描述：
输出一个整数，表示输入字符串最后一个单词的长度。

输入：hello nowcoder
输出：8
说明：最后一个单词为nowcoder，长度为8
*/
func main()  {
	s, _,_ := bufio.NewReader(os.Stdin).ReadLine()
	n := 0
	for i := len(s)-1; i >=0; i--{
		if s[i] == '\n'{
			continue
		} else if s[i] != ' '{
			n++
		} else {
			break
		}
	}
	fmt.Printf("%d\n", n)
}
