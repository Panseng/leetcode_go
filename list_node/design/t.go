package main

import (
	"fmt"
	"math"
)

type MyLinkedList struct {
	Val int
	Next *MyLinkedList
}

func main()  {
	//linkedList := new(MyLinkedList)
	//linkedList.AddAtHead(1)
	//linkedList.AddAtTail(3)
	//linkedList.Show()
	//linkedList.AddAtIndex(1,2)  //链表变为1-> 2-> 3
	//linkedList.Show()
	//fmt.Println(linkedList.Get(1)  )        //返回2
	//linkedList.DeleteAtIndex(1)  //现在链表是1-> 3
	//linkedList.Show()
	//fmt.Println(linkedList.Get(1) )         //返回3

	//a := math.MinInt64
	//b := math.MinInt64
	//if a == b{
	//	fmt.Println("可等")
	//}

	ll := new(MyLinkedList)
	fmt.Println("AddAtHead 7")
	ll.AddAtHead(7)
	fmt.Println("AddAtHead 2")
	ll.AddAtHead(2)
	fmt.Println("AddAtHead 1")
	ll.AddAtHead(1)
	fmt.Println("AddAtHead 0")
	ll.Show()
	ll.AddAtIndex(3,0)
	ll.DeleteAtIndex(2)
	ll.AddAtHead(6)
	ll.AddAtTail(4)
	fmt.Println(ll.Get(4))
	ll.AddAtHead(4)
	ll.AddAtIndex(5,0)
	ll.AddAtHead(6)
	ll.Show()
}

func Constructor() MyLinkedList {
	return MyLinkedList{Val: math.MinInt64}
}

func (this *MyLinkedList) Get(index int) int {
	nodes := this.GetSlice()
	if index < 0 || index >= len(nodes){
		return -1
	}
	return nodes[index].Val
}


func (this *MyLinkedList) AddAtHead(val int)  {
	if this.Val == math.MinInt64{
		this.Val = val
		return
	}
	nodes := this.GetSlice()

	this.Val = val
	this.Next = nodes[0]
}


func (this *MyLinkedList) AddAtTail(val int)  {
	nodes := this.GetSlice()
	n := len(nodes)
	if n == 0{
		this.Val = val
		return
	}
	nodes[len(nodes)-1].Next = &MyLinkedList{Val: val}
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index <= 0{
		this.AddAtHead(val)
		return
	}
	nodes := this.GetSlice()
	n := len(nodes)
	if index > n{
		return
	}
	if index == n{
		this.AddAtTail(val)
		return
	}
	node := nodes[index-1]
	node.Next = &MyLinkedList{Val: val, Next: node.Next}
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index == 0{
		this = this.Next
		return
	}
	nodes := this.GetSlice()
	n := len(nodes)
	if index < n && index > 0{
		node := nodes[index-1]
		node.Next = node.Next.Next
	}
}

func (this *MyLinkedList) GetSlice() (ans []*MyLinkedList ){
	prev := this
	for prev != nil{
		if prev.Val == math.MinInt64{
			break
		}
		ans = append(ans, prev)
		prev = prev.Next
	}
	return
}

func (this *MyLinkedList)Show()  {
	prev := this
	for prev != nil{
		if prev.Val == math.MinInt64{
			break
		}
		fmt.Print(prev.Val)
		prev = prev.Next
	}
	fmt.Println()
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */