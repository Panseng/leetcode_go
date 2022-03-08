package main

import (
	"container/heap"
	"fmt"
	"regexp"
)

func main()  {
	// 最小堆
	t_20220238_3()
	// 切片
	t_20220238_2()
	t_20220238()
}

func t_20220238_3()  {
	h := &IHeap{}
	heap.Init(h)
	nums := []int{1,5,8,3,6,7}
	for _, v := range nums{
		heap.Push(h, v)
		fmt.Println(h)
	}
	fmt.Println(h)
	for h.Len() > 0{
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

func t_20220238_2()  {
	nums1 := []int{0,1,2,3,4,5,6,7,8,9}
	n2 := nums1[1:3]
	n2 = append(n2, 5) // 会改变 nums1 的值
	fmt.Println(n2, nums1)
	addNum(n2) // 会改变nums1的值，但n2不变，因为作为参数时，n2
	fmt.Println(n2, nums1)
	s1 := []int{0,1,2,3}
	s2 := s1[0:2]
	s2 = append(s2, 5,6,7)
	fmt.Println(s1, s2)
	s1 = []int{0,1,2,3}
	s2 = s1[0:2]
	s2 = append(s2, 5)
	s2 = append(s2, 6)
	s2 = append(s2, 7)
	fmt.Println(s1, s2)
}
func addNum(nums []int)  {
	nums[0] = 11 // 通过索引会关联到 n2 & nums1，因为此处的复制还是在原有切片地址上进行的操作
	nums = append(nums, 1) // 会关联到 nums1，不会关联到n2，因为引用地址变了，但切片地址没变
	fmt.Print(nums)
}
func t_20220238(){
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

