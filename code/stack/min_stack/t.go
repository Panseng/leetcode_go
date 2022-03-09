package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack []int
	len int
	min int
}

func main()  {
	ms := Constructor()

	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	fmt.Println(ms.GetMin())
	ms.Pop()
	fmt.Println(ms.Top())
	fmt.Println(ms.GetMin())

	//ms.Push(1)
	//ms.Push(2)
	//fmt.Println(ms.Top())
	//fmt.Println(ms.GetMin())
	//ms.Pop()
	//fmt.Println(ms.GetMin())
	//fmt.Println(ms.Top())

}

func Constructor() MinStack {
	return MinStack{min: math.MaxInt64}
}


func (this *MinStack) Push(val int)  {
	this.stack = append(this.stack, val)
	this.len++
	if this.min > val{
		this.min = val
	}
}


func (this *MinStack) Pop()  {
	if this.len > 0{
		oldTop := this.stack[this.len-1]
		this.stack = this.stack[:this.len-1]
		this.len--
		if this.min == oldTop{
			this.resetMin()
		}
	}
}


func (this *MinStack) Top() int {
	return this.stack[this.len-1]
}


func (this *MinStack) GetMin() int {
	return this.min
}


func (this *MinStack) resetMin() {
	this.min = math.MaxInt64
	for i := 0; i < this.len; i++{
		if this.min > this.stack[i]{
			this.min = this.stack[i]
		}
	}
}





/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
