# 说明
本项目是基于golang刷leetcode题目，[题库](https://leetcode-cn.com/problemset/all/)

## [217. 存在重复元素](https://leetcode-cn.com/problems/contains-duplicate/solution/cun-zai-zhong-fu-yuan-su-by-leetcode-sol-iedd/)
给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
>示例1\
>输入：nums = [1,2,3,1]\
>输出：true

思路1：先用sort.Int排序，然后比较相邻元素是否相同\
```if nums[i] == nums[i-1]```\
思路2：hash表，判断hash表中是否存在该值，存在返回真，不存在则添加\
实现：[arr_repeat](arr_repeat)

## [53. 最大子数组和](https://leetcode-cn.com/problems/two-sum/solution/liang-shu-zhi-he-by-leetcode-solution/)
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。\
子数组 是数组中的一个连续部分。
>示例1\
>输入：nums = [-2,1,-3,4,-1,2,1,-5,4]\
>输出：6\
>解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

思路：贪心算法，将前值累加，如果大于0则继续累计，同时求最大值\
实现：[max_sub_array](max_sub_array)
```
    max := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] + nums[i-1] > nums[i] {
            nums[i] += nums[i-1]
        }
        if nums[i] > max {
            max = nums[i]
        }
    }
    return max
```

## [1. 两数之和](https://leetcode-cn.com/problems/two-sum/solution/liang-shu-zhi-he-by-leetcode-solution/) 
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target的那两个整数，并返回它们的数组下标。\
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。\
你可以按任意顺序返回答案。
>输入：nums = [2,7,11,15], target = 9\
>输出：[0,1]\
>解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

思路1：暴力枚举，通过两层遍历求解\
思路2 ：hash表法，遍历期间，先求解hash表中是否存在目标值，不存在则将当前值加入hash表
实现：[two_sum](two_sum)
```go
    hashTable := map[int]int{}
    for i, x := range nums {
        if p, ok := hashTable[target-x]; ok {
            return []int{p, i}
        }
        hashTable[x] = i
    }
    return nil
```
## [88. 合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array/) 
给你两个按 非递减顺序 排列的整数数组nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。\
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。\
注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
>输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3\
>输出：[1,2,2,3,5,6]\
>解释：需要合并 [1,2,3] 和 [2,5,6] 。\
>合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。

思路1：暴力排序，合并两个数组后，调用系统排序方法\
思路2：双指针法，题干的数组是有序数组（无序数组可以在排序后使用该方法，但效率可能不如第一个思路），直接用双指针进行排序
```go
    sorted := make([]int, 0, m+n)
    p1, p2 := 0, 0
    for {
        if p1 == m {
            sorted = append(sorted, nums2[p2:]...)
            break
        }
        if p2 == n {
            sorted = append(sorted, nums1[p1:]...)
            break
        }
        if nums1[p1] < nums2[p2] {
            sorted = append(sorted, nums1[p1])
            p1++
        } else {
            sorted = append(sorted, nums2[p2])
            p2++
        }
    }
    copy(nums1, sorted)
```
思路3：逆向双指针法
```go
for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
        var cur int
        if p1 == -1 {
            cur = nums2[p2]
            p2--
        } else if p2 == -1 {
            cur = nums1[p1]
            p1--
        } else if nums1[p1] > nums2[p2] {
            cur = nums1[p1]
            p1--
        } else {
            cur = nums2[p2]
            p2--
        }
        nums1[tail] = cur
    }
```
实现：[merge_increase_array](merge_increase_array)

##  [350. 两个数组的交集 II](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/)
给你两个整数数组nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。
>输入：nums1 = [1,2,2,1], nums2 = [2,2]\
>输出：[2,2]

思路1：hash法，短hash表，长数组校验，各循环一次
```go
    if len(nums1) > len(nums2) {
        return intersect(nums2, nums1)
    }
    m := map[int]int{}
    for _, num := range nums1 {
        m[num]++
    }

    intersection := []int{}
    for _, num := range nums2 {
        if m[num] > 0 {
            intersection = append(intersection, num)
            m[num]--
        }
    }
    return intersection
```
思路2：先排序，再加以双指针法
```go
    sort.Ints(nums1)
    sort.Ints(nums2)
    length1, length2 := len(nums1), len(nums2)
    index1, index2 := 0, 0

    intersection := []int{}
    for index1 < length1 && index2 < length2 {
        if nums1[index1] < nums2[index2] {
            index1++
        } else if nums1[index1] > nums2[index2] {
            index2++
        } else {
            intersection = append(intersection, nums1[index1])
            index1++
            index2++
        }
    }
    return intersection
```
实现：[two_array_intersect](two_array_intersect)

## [121. 买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)
给定一个数组 prices ，它的第i 个元素prices[i]表示一支给定股票第 i 天的价格。\
你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。\
返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
>输入：[7,1,5,3,6,4]\
>输出：5\
>解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。\
>注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

思路1：暴力法，两次遍历（数组长度大时，耗时）\
思路2：历史最低值，历史最大差值
```go
	minPrice := prices[0] + 1
	maxProfit :=0
	for _, v := range prices{
		if v<minPrice{
			minPrice = v
		} else if maxProfit < v-minPrice{
			maxProfit = v-minPrice
		}
	}
	return maxProfit
```
实现：[stock_max_profit](stock_max_profit)

## [566. 重塑矩阵](https://leetcode-cn.com/problems/reshape-the-matrix/)

在 MATLAB 中，有一个非常有用的函数 reshape ，它可以将一个m x n 矩阵重塑为另一个大小不同（r x c）的新矩阵，但保留其原始数据。\
给你一个由二维数组 mat 表示的m x n 矩阵，以及两个正整数 r 和 c ，分别表示想要的重构的矩阵的行数和列数。\
重构后的矩阵需要将原始矩阵的所有元素以相同的 行遍历顺序 填充。\
如果具有给定参数的 reshape 操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。

> ![image](img/566-01.jpg)
>输入：mat = [[1,2],[3,4]], r = 1, c = 4\
>输出：[[1,2,3,4]]\
> 
> ![image](img/566-02.jpg)\
> 输入：mat = [[1,2],[3,4]], r = 2, c = 4\
>输出：[[1,2],[3,4]]

思路：二维数组的一维表示，对于m行n列的数组，i < m*n，则一维表示为mat[i/n][i%n]
```go
	m,n:=len(mat), len(mat[0])
	if m*n !=r*c{
		return mat
	}
	newMat := make([][]int, r)
	for i := range newMat{
		newMat[i] = make([]int, c)
	}
	for i:=0;i<m*n;i++{
		newMat[i/c][i%c] = mat[i/n][i%n]
	}
	return newMat
```
实现：[matrix_reshape](matrix_reshape)

## [118. 杨辉三角](https://leetcode-cn.com/problems/pascals-triangle/)
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。\
在「杨辉三角」中，每个数是它左上方和右上方的数的和。
![image](img/118-0.gif)
>输入: numRows = 5\
>输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
> 
>输入: numRows = 1\
>输出: [[1]]

思路：数学，利用对称性\
特征：第n行有n个数，头尾数值为1，中间数值为前一行按序两两相加，且对称\
```go
    // 初始思路
	ans := make([][]int, numRows)
	for i:=range ans{
		ans[i] = make([]int, i+1)
		ans[i][0],ans[i][i] = 1,1
	}
	for i := 1;i< numRows-1;i++{
		for j:=0;j<i;j++{
			ans[i+1][j+1]=ans[i][j]+ans[i][j+1]
		}
	}
    return ans
```
最终思路
```go
    // 合并循环，利用对称性 code 3
	ans := make([][]int, numRows)
	for i:=range ans{
		ans[i] = make([]int, i+1)
		ans[i][0],ans[i][i] = 1,1
		// 利用对称性，对半缩减循环
		for j:=1;j<i/2+1;j++{
			ans[i][j] = ans[i-1][j-1]+ans[i-1][j]
			ans[i][i-j] = ans[i-1][j-1]+ans[i-1][j]
		}
	}
	return ans
```
实现：[yang_hui_triangle](yang_hui_triangle)

## [36. 有效的数独](https://leetcode-cn.com/problems/valid-sudoku/)

请你判断一个9 x 9 的数独是否有效。只需要 根据以下规则 ，验证已经填入的数字是否有效即可。\
数字1-9在每一行只能出现一次。\
数字1-9在每一列只能出现一次。\
数字1-9在每一个以粗实线分隔的3x3宫内只能出现一次。（请参考示例图）\
![image](img/36-0.png)\

**注意：**\
一个有效的数独（部分已被填充）**不一定是可解**的。\
只需要根据以上规则，验证已经填入的数字是否有效即可。\
**空白格** 用 **'.'** 表示。
>输入：board =
>[["5","3",".",".","7",".",".",".","."]
>,["6",".",".","1","9","5",".",".","."]
>,[".","9","8",".",".",".",".","6","."]
>,["8",".",".",".","6",".",".",".","3"]
>,["4",".",".","8",".","3",".",".","1"]
>,["7",".",".",".","2",".",".",".","6"]
>,[".","6",".",".",".",".","2","8","."]
>,[".",".",".","4","1","9",".",".","5"]
>,[".",".",".",".","8",".",".","7","9"]]
>输出：true

思路1：将二维数组展开为一维数组，再用hash表查重，写完代码发现其实复杂度更高了。
实现见[code](sudoku_is_valid/test.go)

思路2：直接二维展开，通过3个hash表分别查重\
这里注意需要将byte转换为int，才能用于index ```index := v - '1'```
```go
func isValidSudoku(board [][]byte) bool {
	var rows, cols [9][9]int
	var ceils [3][3][9]int
	for i, r := range board {
		for j, v := range r {
			if v == '.' {
				continue
			}
			index := v - '1'
			rows[i][index]++
			cols[j][index]++
			ceils[i/3][j/3][index]++
			if rows[i][index] > 1 || cols[j][index] > 1 || ceils[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
```
实现：[sudoku_is_valid](sudoku_is_valid)

## [73. 矩阵置零](https://leetcode-cn.com/problems/set-matrix-zeroes/)
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 [原地](https://baike.baidu.com/item/%E5%8E%9F%E5%9C%B0%E7%AE%97%E6%B3%95) 算法。\

> 示例1\
> 输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]\
> 输出：[[1,0,1],[0,0,0],[1,0,1]]\
> ![image](img/73-1.jpg)
> 
> 示例2\
> 输入：matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]\
> 输出：[[0,0,0,0],[0,4,5,0],[0,3,1,0]]\
> ![image](img/73-2.jpg)

> 函数输入```func setZeroes(matrix [][]int) ```

思路1：变量记录需要变更的行列索引数据```row, col := map[int]bool{}, map[int]bool{}```\
> 两次遍历，第一次获取索引数据，第二次变更\

```go
	row, col := map[int]bool{}, map[int]bool{}
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, r := range matrix {
		for j, _ := range r {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
```
思路2：用矩阵的第一行和第一列代替方法一中的两个标记数组，以达到 O(1)O(1) 的额外空间。但这样会导致原数组的第一行和第一列被修改，无法记录它们是否原本包含 00。因此我们需要额外使用两个标记变量分别记录第一行和第一列是否原本包含 00。
```go
    n, m := len(matrix), len(matrix[0])
    row0, col0 := false, false
    for _, v := range matrix[0] {
        if v == 0 {
            row0 = true
            break
        }
    }
    for _, r := range matrix {
        if r[0] == 0 {
            col0 = true
            break
        }
    }
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }
    if row0 {
        for j := 0; j < m; j++ {
            matrix[0][j] = 0
        }
    }
    if col0 {
        for _, r := range matrix {
            r[0] = 0
        }
    }
```
实现：[matrix_set_zeros](matrix_set_zeros)

## [387. 字符串中的第一个唯一字符](https://leetcode-cn.com/problems/first-unique-character-in-a-string/)
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

> 输入：s = "leetcode"\
> 输出：返回 0\
> 
> 输入：s = "loveleetcode"\
> 输出：返回 2\

**提示**：你可以假定该字符串只包含小写字母。

思路1：hash表记录出现次数，两次遍历获取首个出现一次的索引\
注意，range字符串返回的是int32的值
```go
func firstUniqChar2(s string)int  {
	a := make(map[int32]int)
	for _, v:=range s{
		a[v]++
	}
	for i,v:=range s{
		if a[v] == 1{
			return i
		}
	}
	return -1
}
```
思路2：hash表法，利用先进先出的数组，嵌套遍历\
关键思路，对首个非唯一的字符串进行排除
```go
	for len(que) > 0 && hashTable[que[0].ch] >1{
		que = que[1:] // 先进先出
	}
```
先进先出关键代码```que = que[1:]```
代码：
```go
func firstUniqChar5(s string) int {
	var que []loc
	hashTable := make(map[int32]int)
	for i,v:=range s{
		hashTable[v]++
		if hashTable[v] == 1{
			que = append(que, loc{v, i})
		}
	}
	for len(que) > 0 && hashTable[que[0].ch] >1{
		que = que[1:] // 先进先出
	}
	if len(que)>0{
		return que[0].pos
	}
	return -1
}
```
实现：[char_first_unique](char_first_unique)\
**_注意_**：力扣官方解法，仅限定小写字母，所以可以声明```[26]int{}```

## [383. 赎金信](https://leetcode-cn.com/problems/ransom-note/)
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。\
如果可以，返回 true ；否则返回 false 。\
magazine 中的每个字符只能在 ransomNote 中使用一次。\
> 示例 1：\
> 输入：ransomNote = "a", magazine = "b"\
> 输出：false
>
> 示例 2：\
> 输入：ransomNote = "aa", magazine = "ab"\
> 输出：false
>
> 示例 3：\
> 输入：ransomNote = **"aa"**, magazine = **"aba"**\
> 输出：true

关键点，```ransomNote```中同一字符出现的次数要小于等于```magazine```中字符

思路：hash表法，先统计```magazine```中字符出现次数，遍历```ransomNote```，如果```ransomNote```出现其他字符或者字符多于```magazine```则返回false
```go
func canConstruc2(ransomNote string, magazine string) bool{
	mHash:=make(map[int32]int)
	for _,v:=range magazine{
		mHash[v]++
	}
	for _,v:=range ransomNote{
		mHash[v]--
		if mHash[v] < 0{
			return false
		}
	}
	return true
}
```
实现：[char_can_construct](char_can_construct)

## [242. 有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。\
**注意**：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。
> 示例1:\
> 输入: s = "anagram", t = "nagaram"\
> 输出: true
>
> 示例 2:\
> 输入: s = "rat", t = "car"\
> 输出: false

思路1：hash表法，同383. 赎金信
```go
func isAnagram2(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	sHash := make(map[rune]int)
	for _,v:=range s{
		sHash[v]++
	}
	for _,v:=range t{
		sHash[v]--
		if sHash[v] < 0{
			return false
		}
	}
	return true
}
```
思路2：对两个字符串进行排序后比较
```go
func isAnagram3(s string, t string) bool {
	ls1,ls2 := []byte(s),[]byte(t)
	sort.Slice(ls1, func(i, j int) bool {
		return ls1[i]<ls1[j]
	})
	sort.Slice(ls2, func(i, j int) bool {
		return ls2[i]<ls2[j]
	})
	return string(ls1) == string(ls2)
}
```
实现：[char_is_anagram](char_is_anagram)

## [141. 环形链表](https://leetcode-cn.com/problems/linked-list-cycle/)
给你一个链表的头节点 head ，判断链表中是否有环。\
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递。仅仅是为了标识链表的实际情况。\
如果链表中存在环，则返回 true 。 否则，返回 false 。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
```
> ![](img/141-1.png)\
> 输入：head = [3,2,0,-4], pos = 1\
> 输出：true\
> 解释：链表中有一个环，其尾部连接到第二个节点。
> 
> ![](img/141-2.png)\
> 输入：head = [1,2], pos = 0\
> 输出：true\
> 解释：链表中有一个环，其尾部连接到第一个节点。
> 
> ![](img/141-3.png)\
> 输入：head = [1], pos = -1\
> 输出：false\
> 解释：链表中没有环。

思路1：快慢指针，如果快指针先到达nil值，则没有循环，如果快指针和慢指针相等则说明有循环
```go
func hasCycle(head *ListNode) bool {
    fast, slow:= head, head
    for{
        if fast == nil || fast.Next == nil{
            return false
        }
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow{
            break
        }
    }
    return true
}
```
思路2：hash表法
```go
func hasCycle(head *ListNode) bool {
    seen := map[*ListNode]bool{}
    for head != nil {
        if seen[head] {
            return true
        }
        seen[head] = true
        head = head.Next
    }
    return false
}
```

## [21. 合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/)
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
```
![](img/21-0.jpg)
> 输入：l1 = [1,2,4], l2 = [1,3,4]\
> 输出：[1,1,2,3,4,4]

思路1：递归调用
```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil{
        return list2
    }
    if list2 == nil{
        return list1
    }
    if list1.Val < list2.Val{
        list1.Next = mergeTwoLists(list1.Next, list2)
        return list1
    }
    list2.Next = mergeTwoLists(list2.Next, list1)
    return list2
}
```
[思路2](https://leetcode-cn.com/problems/merge-two-sorted-lists/solution/he-bing-liang-ge-you-xu-lian-biao-by-leetcode-solu/) ：双指针、哨兵 & 游标法

```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var ls *ListNode = new (ListNode)
    var p = ls // ls作为哨兵，p作为游标，游标不断Next前进，直到最后遍历完
    for list1 != nil && list2 != nil{
        if list1.Val < list2.Val{
            p.Next = list1
            list1 = list1.Next
        } else{
            p.Next = list2
            list2 = list2.Next
        }
        p = p.Next

    }
    if list1 == nil{
        p.Next = list2
    } else{
        p.Next = list1
    }
    return ls.Next
}
```

## [203. 移除链表元素](https://leetcode-cn.com/problems/remove-linked-list-elements/)
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

> ![](img/203-0.jpg)
> 输入：head = [1,2,6,3,4,5,6], val = 6\
> 输出：[1,2,3,4,5]

思路1：遍历法，通过配置游标的方式
code1：
```go
func removeElements(head *ListNode, val int) *ListNode {
    for head != nil && head.Val == val{ // 去除头部可能相同值的节点
        head = head.Next
    }
    if head == nil{
        return head
    }
    prev := head // 配置游标
    for prev.Next != nil{
        if prev.Next.Val == val{
            prev.Next = prev.Next.Next // 删除相同值节点
        } else{
            prev = prev.Next // 游往下一个节点
        }
    }
    return head
}
```
code2：
```go
func removeElements(head *ListNode, val int) *ListNode {
    var preHead *ListNode = &ListNode{Next: head}
    for temp := preHead; temp.Next != nil;{ // 游标
        if temp.Next.Val == val{
            temp.Next = temp.Next.Next
        } else{
            temp = temp.Next
        }
    }
    return preHead.Next
}
```

思路2：递归
```go
func removeElements(head *ListNode, val int) *ListNode {
    if head == nil{
        return head
    }
    head.Next = removeElements(head.Next, val)
    if head.Val == val{
        return head.Next
    }
    return head
}
```