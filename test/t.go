package main

import (
	"container/heap"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	placeholderFmt()
	//t_20220309()
	//// 最小堆
	//t_20220238_3()
	//// 切片
	//t_20220238_2()
	//t_20220238()
}

// fmt 占位符
func placeholderFmt() {
	//通用verbs
	//  %v       　值的默认格式
	//  %+v      　类似%v,但输出结构体时会添加字段名
	//	%#v 　　 Go语法表示值
	//	%T  　　 Go语法表示类型
	//	%%     　 百分号表示
	//如下示例：
	type Sample struct {
		Title string
		name  string
		Age   int
	}
	s := Sample{"测试", "wentao", 26}
	fmt.Printf("%v\n", s)       // {测试 wentao 26}
	fmt.Printf("%+v\n", s)      // {Title:测试 name:wentao Age:26}
	fmt.Printf("%#v\n", s)      // main.Sample{Title:"测试", name:"wentao", Age:26}
	fmt.Printf("%T\n", s)       //  main.Sample
	fmt.Printf("%v%%\n", s.Age) //  26%

	// bool值
	fmt.Printf("%t\n", true) // true

	// 整数
	// %b  表示二进制
	// %c  该值对应的unicode吗值
	// %d  表示十进制
	// %o  表示八进制
	// %q  该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	// %x  表示为十六进制，使用a-f
	// %X  表示为十六进制，使用A-F
	// %U  表示为Unicode格式：U+1234，等价于"U+%04X"
	fmt.Printf("%b\n", 11)     // 1011
	fmt.Printf("%c\n", 0x4E2D) // 中
	fmt.Printf("%d\n", 0x4E2D) // 20013
	fmt.Printf("%o\n", 0x4E2D) // 47055
	fmt.Printf("%q\n", 0x4E2D) // '中'
	fmt.Printf("%x\n", 0x4E2D) // 4e2d
	fmt.Printf("%X\n", 0x4E2D) // 4E2D
	fmt.Printf("%U\n", 0x4E2D) // U+4E2D

	// 浮点数与复数
	// %b  无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
	// %e  科学计数法，例如 -1234.456e+78
	// %E  科学计数法，例如 -1234.456E+78
	// %f  有小数点而无指数，例如 123.456
	// %F  等价于%f
	// %g  根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	// %G  根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
	fmt.Printf("%b\n", 10.45) //5882827013252710p-49
	fmt.Printf("%e\n", 10.45) //1.045000E+01
	fmt.Printf("%E\n", 10.45) //1.045000E+01
	fmt.Printf("%f\n", 10.45) //10.450000
	fmt.Printf("%F\n", 10.45) //10.450000
	fmt.Printf("%g\n", 10.45) //10.45
	fmt.Printf("%G\n", 10.45) //10.45

	// string与[]byte
	// %s  输出字符串表示（string类型或[]byte)
	// %q  双引号围绕的字符串，由Go语法安全地转义
	// %x  十六进制，小写字母，每字节两个字符 （使用a-f）
	// %X  十六进制，大写字母，每字节两个字符 （使用A-F）
	fmt.Printf("%s\n", []byte("go开发"))      //go开发
	fmt.Printf("%s\n", "go开发")          //go开发
	fmt.Printf("%q\n", []byte("go开发"))          //"go开发"
	fmt.Printf("%q\n", "go开发")          //"go开发"
	fmt.Printf("%x\n", "go开发")          //676fe5bc80e58f91
	fmt.Printf("%X\n", "go开发")          //676FE5BC80E58F91

	// Slice
	// %p       切片第一个元素的指针，或相应值的指针
	// %v 切片的值
	fmt.Printf("%p\n", []byte("go开发"))          // 0xc00000a110
	fmt.Printf("%v\n", []byte("go开发"))          // [103 111 229 188 128 229 143 145]
	fmt.Printf("%p\n", []int{1, 2, 3, 45, 65})  // 0xc00000c420
	fmt.Printf("%v\n", []int{1, 2, 3, 45, 65})  // [1 2 3 45 65]
	strP := "go开发"
	fmt.Printf("%p\n", &strP)            // 0xc000044260
}
func t_20220309() {
	s := "$22.65美元"
	m, _ := regexp.Compile("[0-9.]+")
	ns := m.FindString(s)
	fmt.Println(strconv.ParseFloat(ns, 64))
}
func t_20220238_3() {
	h := &IHeap{}
	heap.Init(h)
	nums := []int{1, 5, 8, 3, 6, 7}
	for _, v := range nums {
		heap.Push(h, v)
		fmt.Println(h)
	}
	fmt.Println(h)
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
}

type IHeap []int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i] < h[j] } // 最小堆
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func t_20220238_2() {
	nums1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	n2 := nums1[1:3]
	n2 = append(n2, 5) // 会改变 nums1 的值
	fmt.Println(n2, nums1)
	addNum(n2) // 会改变nums1的值，但n2不变，因为作为参数时，n2
	fmt.Println(n2, nums1)
	s1 := []int{0, 1, 2, 3}
	s2 := s1[0:2]
	s2 = append(s2, 5, 6, 7)
	fmt.Println(s1, s2)
	s1 = []int{0, 1, 2, 3}
	s2 = s1[0:2]
	s2 = append(s2, 5)
	s2 = append(s2, 6)
	s2 = append(s2, 7)
	fmt.Println(s1, s2)
}
func addNum(nums []int) {
	nums[0] = 11           // 通过索引会关联到 n2 & nums1，因为此处的复制还是在原有切片地址上进行的操作
	nums = append(nums, 1) // 会关联到 nums1，不会关联到n2，因为引用地址变了，但切片地址没变
	fmt.Print(nums)
}
func t_20220238() {
	s := "abc"
	ss := [2]byte{s[1]}
	fmt.Println(ss)

	s2 := "e:\\code\\golang\\leetcode_go\\notes\\base.md"
	if found, _ := regexp.MatchString("^[a-z]:\\\\", s2); found {
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
	c := float64(a) / 2
	fmt.Println(c)
}
