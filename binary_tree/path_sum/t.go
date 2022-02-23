package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := [][]int{}
	b :=[]int{1,-2}
	al := append(b, 1,0)
	a = append(a,al)
	fmt.Println(a,al)
	ar := append(b, 3,2)
	a = append(a,ar)
	fmt.Println(a,ar)

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: -2,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: -1,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -2,
			},
		},
	}
	fmt.Println(pathSum(root, 1))
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	nodes := []*TreeNode{root}
	valSum := [][]int{} // 记录每个节点路径，最后一位是当前路径和
	valSum = append(valSum, []int{root.Val, root.Val})
	n := len(nodes) // 当前层节点数目
	level := 1      // 当前节点为第几层，root为第1层
	for n > 0 {
		temN := []*TreeNode{}
		temV := make([][]int, 0)
		for j := 0; j < n; j++ {
			node := nodes[j]
			ps := valSum[j][level]    // 路径和
			path := valSum[j][:level] // 路径
			if node.Left == nil && node.Right == nil && ps == targetSum {
				ans = append(ans, path)
			}
			if node.Left != nil {
				temN = append(temN, node.Left)
				tLeft := append(path, node.Left.Val, ps+node.Left.Val)
				temV = append(temV, tLeft)
			}
			if node.Right != nil {
				temN = append(temN, node.Right)
				// 注意，go的slice切片，会在原地址操作
				// 如果 tRight 直接 = append(path, node.Right.Val, ps+node.Right.Val)
				// 则 上方的 tLeft 值会被 tRight 值覆盖
				// 因为 tLeft 值与 tRight 值都是在path切片原地址上进行的操作，
				tRight := make([]int, level)
				copy(tRight, path)
				tRight = append(tRight, node.Right.Val, ps+node.Right.Val)
				temV = append(temV, tRight)
			}
		}
		level++
		nodes = temN
		valSum = temV
		n = len(nodes)
	}
	return ans
}
