package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	head := generator(a)
	head.Show()
	//head = reverseList(head)
	//head = reverseList2(head)
	head = reverseList3(head)
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
	fmt.Println()
}

// 迭代1，通过数组辅助
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	//nodeMap := make(map[*ListNode]bool)
	nodeList := []*ListNode{}
	for head != nil {
		//if nodeMap[head] {
		//	break
		//}
		//nodeMap[head] = true
		nodeList = append(nodeList, head)
		head = head.Next
	}
	preHead := new(ListNode)
	prev := preHead
	for i := len(nodeList) - 1; i >= 0; i-- {
		prev.Next = nodeList[i]
		prev = prev.Next
	}
	prev.Next = nil
	return preHead.Next
}

// 迭代2，无辅助
func reverseList2(head *ListNode) *ListNode {
	var preHead *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = preHead // 倒叙排列
		preHead = curr
		curr = next
	}
	return preHead
}

// 迭代2，无辅助
func reverseList4(head *ListNode) *ListNode {
	var nilHead *ListNode
	node1 := head
	for node1 != nil {
		node2 := node1.Next
		node1.Next = nilHead // 倒叙排列
		nilHead = node1
		node1 = node2
	}
	return nilHead
}

// 递归
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := reverseList3(head.Next) //  一直递归到head.Next = 5，cur为5
	head.Next.Next = head          // head为游标，将cur为5的next指向当前head（4）
	head.Next = nil
	return cur
}
