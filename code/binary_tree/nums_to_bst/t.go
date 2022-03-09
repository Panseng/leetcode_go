package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main()  {
	fmt.Println(sortedArrayToBST([]int{-10,-3,0,5,9}))
}

func sortedArrayToBST(nums []int) *TreeNode {
	var node *TreeNode
	for _,v:=range nums{
		node = add(node, v)
	}
	return node
}

// 计算层高
func  getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(getHeight(node.Right), getHeight(node.Left)) + 1 // 每层选最大的长度+1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 计算平衡因子
func  getBalanceFactor(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return getHeight(node.Left) - getHeight(node.Right)
}

///////////////////////////////////////////////////
// LL T1<Z<T2< X <T3<Y<T4                        //
//        y                             x       //
//       / \                           /   \     //
//      x   T4 向右旋转 (y)      z     y    //
//     / \      - - - - - - - ->    / \    / \   //
//    z   T3                    T1 T2 T3 T4  //
//   / \                                         //
// T1   T2                                       //
///////////////////////////////////////////////////
// 右旋
func  rightRotate(y *TreeNode) *TreeNode {
	x := y.Left
	T3 := x.Right

	x.Right = y
	y.Left = T3
	return x
}

////////////////////////////////////////////////
// RR T1<Y<T2< X <T3<Z<T4                     //
//    y                             x         //
//  /  \                          /   \       //
// T1 x 向左旋转 (y)       y     z      //
//     / \  - - - - - - - ->  / \    / \     //
//   T2  z                T1 T2 T3 T4    //
//       / \                                  //
//      T3 T4                                 //
////////////////////////////////////////////////
// 左旋
func  leftRotate(y *TreeNode) *TreeNode {
	x := y.Right
	T2 := x.Left

	x.Left = y
	y.Right = T2
	return x
}

// 添加
func  add(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: value}
	}
	if value < node.Val {
		node.Left = add(node.Left, value)
	} else if value > node.Val {
		node.Right = add(node.Right, value)
	}

	balanceFactor := getBalanceFactor(node)
	if balanceFactor > 1 && getBalanceFactor(node.Left) >= 0 {
		return rightRotate(node)
	}
	if balanceFactor > 1 && getBalanceFactor(node.Left) < 0 {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}
	if balanceFactor < -1 && getBalanceFactor(node.Right) <= 0{
		return leftRotate(node)
	}
	if balanceFactor < -1 && getBalanceFactor(node.Right) >0{
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}
	return node
}
