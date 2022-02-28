# [精选200](https://leetcode-cn.com/problem-list/qg88wci/)

## [156. 上下翻转二叉树](https://leetcode-cn.com/problems/binary-tree-upside-down/)
给你一个二叉树的根节点 root ，请你将此二叉树上下翻转，并返回新的根节点。 \
你可以按下面的步骤翻转一棵二叉树：
- 原来的左子节点变成新的根节点
- 原来的根节点变成新的右子节点
- 原来的右子节点变成新的左子节点

![](../img/156-1.jpg) \
上面的步骤逐层进行。题目数据保证每个右节点都有一个同级节点（即共享同一父节点的左节点）且不存在子节点。


> 示例 1： \
![](../img/156-2.jpg) \
> 输入：root = [1,2,3,4,5] \
> 输出：[4,5,2,null,null,3,1]
>
> 示例 2： \
> 输入：root = [] \
> 输出：[]
>
> 示例 3： \
> 输入：root = [1] \
> 输出：[1]


思路：看图迭代
```go
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
    if root == nil || root.Left == nil{
        return root
    }
    
    x := root
    y := root.Left
    z := root.Right
    root.Left = nil
    root.Right = nil
    yLeft := y.Left
    yRight := y.Right

    y.Left = z
    y.Right = x
    for yLeft != nil{
        x = y
        y = yLeft
        z = yRight

        yLeft = y.Left
        yRight = y.Right

        y.Left = z
        y.Right = x
    }
    return y
}
```

## [159. 至多包含两个不同字符的最长子串](https://leetcode-cn.com/problems/longest-substring-with-at-most-two-distinct-characters/)
给定一个字符串 s ，找出 至多 包含两个不同字符的最长子串 t ，并返回该子串的长度。

> 示例 1: \
> 输入: "eceba" \
> 输出: 3 \
> 解释: t 是 "ece"，长度为3。
>
> 示例 2: \
> 输入: "ccaabbb" \
> 输出: 5 \
> 解释: t 是 "aabbb"，长度为5。

思路：迭代
```go
func lengthOfLongestSubstringTwoDistinct(s string) int {
    n := len(s)
    if n < 3{
        return n
    }
    ans := 2
    for i := 0; i < n; i++{
        var next byte
        for j := i+1; j < n; j++{
            if next == 0{
                if j-i+1 > ans{
                    ans = j-i+1
                }
                if s[j] != s[i]{
                    next = s[j]
                }
                continue
            }
            if s[j] == s[i] || s[j] == next{
                if j-i+1 > ans{
                    ans = j-i+1
                }
                continue
            }
            break
        }
    }
    return ans
}
```