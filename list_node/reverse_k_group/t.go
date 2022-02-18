package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := []int{1, 2, 3, 4, 5,6,7,8,9,10}
	head := generator(a)
	head.Show()
	head = reverseKGroup2(head, 2)
	fmt.Println()
	head.Show()
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	prev := head
	n := 0
	for prev != nil{
		prev = prev.Next
		n++
	}
	if n < k || n == 1 || k == 1{
		return head
	}

	var leftTail, rightTail, firstHead, nilHead, node1, node2 *ListNode
	node1 = head
	i := 0
	// 整体思路是：以k个节点为一组，分头尾，取余后，尾下标为0，头下标为k-1，上一组的尾接下一组的头
	for ;i < n && i < (n/k)*k;i++ {
		if i%k == 0{
			nilHead = nil
		}
		node2 = node1.Next
		node1.Next = nilHead
		nilHead = node1
		node1 = node2
		if i % k == 0{
			rightTail = nilHead // 获取当前组的尾
			continue
		}
		if i%k != k-1{ // 如果当前节点不是新的头节点，则进入下一个循环
			continue
		}
		if i == k-1{
			firstHead = nilHead
		} else {
			leftTail.Next = nilHead
		}
		leftTail = rightTail // 记录当前组的尾，因为下一个组的尾会覆盖rightTail值
	}

	if i < n{
		leftTail.Next = node1
	}
	return firstHead
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	nodes := []*ListNode{}
	prev := head
	for prev != nil{
		nodes = append(nodes, prev)
		prev = prev.Next
	}
	n := len(nodes)
	if n < k || n == 1{
		return head
	}
	node := &ListNode{}
	prev = node
	i := 0
	for ;i < n && i < (n/k)*k;i+=k{
		next, tail := slice2Node(nodes[i:i+k])
		prev.Next = next
		prev = tail
	}
	for ; i < n; i++{
		prev.Next = nodes[i]
		prev = prev.Next
	}
	prev.Next = nil
	return node.Next
}

func slice2Node(nodes []*ListNode) (*ListNode, *ListNode){
	node := &ListNode{Next:nil}
	prev := node
	n := len(nodes)
	for n > 0{
		prev.Next = nodes[n-1]
		prev = prev.Next
		nodes = nodes[:n-1]
		n = len(nodes)
	}
	return node.Next, prev
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
