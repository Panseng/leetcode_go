package main

//节点结构体
type Node struct {
	Val int
	Left *Node
	Right *Node
}
//二叉树结构体
type BinaryTree struct {
	root *Node
	count int
}
//查询时记录下标
var i int=0
//添加节点
func (node *Node)addNode(newNode *Node){
	if node.Val>newNode.Val{
		if node.Left==nil{
			node.Left=newNode
		}else{
			node.Left.addNode(newNode)
		}
	}else{
		if node.Right==nil{
			node.Right=newNode
		}else{
			node.Right.addNode(newNode)
		}
	}
}
func (bt *BinaryTree)add(Val int){
	var newNode Node
	newNode.Val=Val
	if bt.root==nil{
		bt.root=&newNode
	}else{
		bt.root.addNode(&newNode)
	}
	bt.count++
}
//遍历二叉树（中序：左-中-右），返回切片类型
func (node *Node)toArrayNode(Vals []int){
	if node.Left!=nil{
		node.Left.toArrayNode(Vals)
	}
	Vals[i]=node.Val
	i++
	if node.Right!=nil{
		node.Right.toArrayNode(Vals)
	}
}
func (bt *BinaryTree)toArray()[]int{
	if bt.root==nil{
		return nil
	}
	Vals:=make([]int,bt.count)
	bt.root.toArrayNode(Vals)
	return Vals
}
//返回该树的长度
func (bt *BinaryTree)size()int{
	return bt.count
}