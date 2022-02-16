package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main()  {
	na := new(ListNode)
	nb := new(ListNode)
	if na.Next == nb.Next{
		fmt.Println("相等")
	}

	var naa ListNode
	var nbb ListNode
	if naa == nbb{
		fmt.Println(naa,"相等",nbb)
	}

	a :=[]int{1,2,3,4,5}
	head1 := generator(a)
	b :=[]int{6,7,8,9,10}
	head2 := generator(b)

	fmt.Println(getIntersectionNode(head1,head2))

	head1.Show()
	fmt.Println()
	head2.Show()
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	prevA, prevB := headA, headB
	// 如果两条链长度一样，而不相交，则同时指向 nill
	for prevA != prevB{
		if prevA == nil{
			prevA = headB
		} else{
			prevA = prevA.Next
		}
		if prevB == nil{
			prevB = headA
		} else{
			prevB = prevB.Next
		}
	}
	if prevA == prevB{
		fmt.Println(prevA, " == ", prevB)
	}
	return prevA
}

func generator(list []int)  *ListNode{
	head := &ListNode{Val: list[0]} // 哨兵
	tail := head // 游标
	for _,v:=range list{
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return head.Next
}
func (h *ListNode) Show() {
	fmt.Print(h.Val)
	for h.Next != nil {
		h = h.Next
		fmt.Print(h.Val)
	}
}


