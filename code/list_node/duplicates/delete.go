package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//a := []int{1, 1, 2, 3, 3}
	a := []int{1, 1, 1,2,2,2}
	head := generator(a)
	head.Show()
	//deleteDuplicates(head)
	deleteDuplicates2(head)
	fmt.Println()
	head.Show()
}
func generator(list []int) *ListNode {
	head := &ListNode{Val: list[0]} // 哨兵
	tail := head                    // 游标
	for _, v := range list {
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

// 迭代
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	prev := head
	for prev != nil && prev.Next != nil {
		for prev.Next != nil && prev.Val == prev.Next.Val {
			prev.Next = prev.Next.Next
		}
		prev = prev.Next
	}
	return head
}

// 迭代2
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prev := head
	for prev.Next != nil {
		if prev.Val == prev.Next.Val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return head
}