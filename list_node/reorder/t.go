package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := []int{1, 2, 3, 4,5}
	head := generator(a)
	reorderList(head)
	head.Show()
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l1.Show()
	l2.Show()
	l2 = reverseList(l2)
	mergeList(l1, l2)
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
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
	fmt.Println()
}
