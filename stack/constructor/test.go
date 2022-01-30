package main

import "fmt"

func main() {
	a := new(MyQueue)
	a.Push(1)
	a.Push(2)
	fmt.Println(a.Peek())
	fmt.Println(a.Pop())
	fmt.Println(a.Empty())
}

type MyQueue struct {
	mq []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.mq = append(this.mq, x)
}

func (this *MyQueue) Pop() int {
	v := this.mq[len(this.mq)-1]
	this.mq = this.mq[:len(this.mq)-1]
	return v
}

func (this *MyQueue) Peek() int {
	return this.mq[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.mq) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
