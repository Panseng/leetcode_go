# [精选200](https://leetcode-cn.com/problem-list/qg88wci/)
* [精选200](#精选200)
  * [156-上下翻转二叉树](#156-上下翻转二叉树)
  * [159-至多包含两个不同字符的最长子串](#159-至多包含两个不同字符的最长子串)
  * [161-相隔为-1-的编辑距离](#161-相隔为-1-的编辑距离)
  * [163-缺失的区间](#163-缺失的区间)
  * [170-两数之和-iii---数据结构设计](#170-两数之和-iii---数据结构设计)
  * [253-会议室-ii](#253-会议室-ii)
  * [340-至多包含-k-个不同字符的最长子串](#340-至多包含-k-个不同字符的最长子串)

## [156. 上下翻转二叉树](https://leetcode-cn.com/problems/binary-tree-upside-down/)
给你一个二叉树的根节点 root ，请你将此二叉树上下翻转，并返回新的根节点。 \
你可以按下面的步骤翻转一棵二叉树：
- 原来的左子节点变成新的根节点
- 原来的根节点变成新的右子节点
- 原来的右子节点变成新的左子节点

![](img/156-1.jpg) \
上面的步骤逐层进行。题目数据保证每个右节点都有一个同级节点（即共享同一父节点的左节点）且不存在子节点。

> 示例 1： \
![](img/156-2.jpg) \
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

## [161. 相隔为 1 的编辑距离](https://leetcode-cn.com/problems/one-edit-distance/)
给定两个字符串 s 和 t，判断他们的编辑距离是否为 1。

注意： \
满足编辑距离等于 1 有三种可能的情形：
1. 往 s 中插入一个字符得到 t
2. 从 s 中删除一个字符得到 t
3. 在 s 中替换一个字符得到 t

> 示例 1： \
> 输入: s = "ab", t = "acb" \
> 输出: true \
> 解释: 可以将 'c' 插入字符串 s 来得到 t。
>
> 示例 2: \
> 输入: s = "cab", t = "ad" \
> 输出: false \
> 解释: 无法通过 1 步操作使 s 变为 t。
>
> 示例 3: \
> 输入: s = "1203", t = "1213" \
> 输出: true \
> 解释: 可以将字符串 s 中的 '0' 替换为 '1' 来得到 t。

思路：遍历，按字符串长度差异情况进行字符对比
```go
func isOneEditDistance(s string, t string) bool {
    ns, nt := len(s), len(t)
    if ns > nt{
        s, t, ns, nt = t, s, nt, ns
    }
    difN := 0
    if ns == nt{
        for i := 0; i < ns; i++{
            if s[i] == t[i]{
                continue
            }
            difN++
            if difN > 1{
                return false
            }
        }
        if difN == 1{ // 位数相同，则一定要有一个差异
            return true
        }
        return false
    } else if ns == nt-1{
        // 位数差1，则最多1个差异
        // 如果差异在末尾，则 difN 为0
        for i := 0; i < ns; i++{
            if s[i] == t[i+difN]{
                continue
            }
            difN++
            if difN > 1{
                return false
            }
            i--
        }
        return true
    }  else{ // 位数差大于1
        return false
    }
}
```
## [163. 缺失的区间](https://leetcode-cn.com/problems/missing-ranges/)
给定一个排序的整数数组 nums ，其中元素的范围在 闭区间 [lower, upper] 当中，返回不包含在数组中的缺失区间。

> 示例： \
> 输入: nums = [0, 1, 3, 50, 75], lower = 0 和 upper = 99, \
> 输出: ["2", "4->49", "51->74", "76->99"]

思路1：分情况
```go
func findMissingRanges(nums []int, lower int, upper int) []string {
    n := len(nums)
    if n == 0{
        if lower == upper{
            return []string{strconv.Itoa(lower)}
        } else{
            return []string{strconv.Itoa(lower)+"->"+strconv.Itoa(upper)}
        }
    }
    ans := []string{}
    if nums[0] > lower{
        if nums[0]-lower == 1{
            ans = append(ans, strconv.Itoa(lower))
        } else{
            ans = append(ans, strconv.Itoa(lower) + "->" + strconv.Itoa(nums[0]-1))
        }
    }

    for i := 0; i < n-1; i++{
        cur, next := nums[i], nums[i+1]
        if next-cur == 2{
            ans = append(ans, strconv.Itoa(cur+1))
        } else if next - cur > 2{
            ans = append(ans, strconv.Itoa(cur+1) + "->" + strconv.Itoa(next-1))
        }
    }

    last := nums[n-1]
    if upper-last == 1{
        ans = append(ans, strconv.Itoa(upper))
    } else if upper-last > 1{
        ans = append(ans, strconv.Itoa(last+1)+"->"+strconv.Itoa(upper))
    }
    return ans
}
```

思路2：补充数组
- 用`upper+1`补充数组的最后一位，因为要比较最后一位到upper之间的区间
  - `strconv.Itoa(cur) + "->" + strconv.Itoa(v-1)`
- 用游标逐个比较，`cur := lower`作为移动的游标
  - 右移`cur++` or `cur = v+1`
```go
func findMissingRanges(nums []int, lower int, upper int) []string {
    ans := []string{}
    // 注意，这里要upper+1，因为会比较最后一位到upper之间的区间
    // strconv.Itoa(cur) + "->" + strconv.Itoa(v-1)
    nums = append(nums, upper+1)
    cur := lower // 作为移动的游标
    for _, v := range nums{
        // 这里没有针对 cur > v的情况
        // 是因为cur > v说明，v重复，可以忽略
        if cur == v{ // 比较是否有缺失
            cur++ // 不断右移
        } else if cur < v{
            if cur == v-1{
                ans = append(ans, strconv.Itoa(cur))
            } else {
                ans = append(ans, strconv.Itoa(cur) + "->" + strconv.Itoa(v-1))
            }
            cur = v+1 // 右移
        }
    }
    return ans
}
```

## [170. 两数之和 III - 数据结构设计](https://leetcode-cn.com/problems/two-sum-iii-data-structure-design/)
设计一个接收整数流的数据结构，该数据结构支持检查是否存在两数之和等于特定值。 \
实现 TwoSum 类：
- TwoSum() 使用空数组初始化 TwoSum 对象
- void add(int number) 向数据结构添加一个数 number
- boolean find(int value) 寻找数据结构中是否存在一对整数，使得两数之和与给定的值相等。如果存在，返回 true ；否则，返回 false 。

> 示例： \
> 输入： \
> ["TwoSum", "add", "add", "add", "find", "find"] \
> [[], [1], [3], [5], [4], [7]] \
> 输出： \
> [null, null, null, null, true, false] \
> 解释： \
> TwoSum twoSum = new TwoSum(); \
> twoSum.add(1);   // [] --> [1] \
> twoSum.add(3);   // [1] --> [1,3] \
> twoSum.add(5);   // [1,3] --> [1,3,5] \
> twoSum.find(4);  // 1 + 3 = 4，返回 true \
> twoSum.find(7);  // 没有两个整数加起来等于 7 ，返回 false \

思路1：hash法
```go
type TwoSum struct {
    Val []int
}


func Constructor() TwoSum {
    return TwoSum{}
}


func (this *TwoSum) Add(number int)  {
    this.Val = append(this.Val, number)
}


func (this *TwoSum) Find(value int) bool {
    mp := make(map[int]struct{})
    for _, num := range this.Val{
        if _, ok := mp[value-num]; ok{
            return true
        }
        mp[num] = struct{}{}
    }
    return false
}
```
优化，只存map数据，不存数组
```go
type TwoSum struct {
    Val map[int]int
}


func Constructor() TwoSum {
    return TwoSum{
        make(map[int]int),
    }
}


func (this *TwoSum) Add(number int)  {
    this.Val[number]++
}


func (this *TwoSum) Find(value int) bool {
    for k, v := range this.Val{
        if value-k == k{
            if v > 1{
                return true
            }
            continue
        }
        if _, ok := this.Val[value-k]; ok{
            return true
        }
    }
    return false
}
```

思路2：双指针法
- 双指针法依赖排序，需要是否排序的tag
- 每次add后，需要重新排序
```go
type TwoSum struct {
    Val []int // 值数组
    isSort bool // 排序标识
}


func Constructor() TwoSum {
    return TwoSum{}
}


func (this *TwoSum) Add(number int)  {
    this.Val = append(this.Val, number)
    this.isSort = false
}


func (this *TwoSum) Find(value int) bool {
    if !this.isSort{
        sort.Ints(this.Val)
        this.isSort = true
    }
    for left, right := 0, len(this.Val)-1; left < right; { // 双指针
        sum := this.Val[left] + this.Val[right]
        if sum > value{ // sum更大，则需要右侧指针左移，减小和
            right--
        } else if sum < value{ // sum更小，则需要左侧指针右移，增加和
            left++
        } else{
            return true
        }
    }
    return false
}
```

## [253. 会议室 II](https://leetcode-cn.com/problems/meeting-rooms-ii/)
给你一个会议时间安排的数组 intervals ，每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi] ，返回 所需会议室的最小数量 。

> 示例 1： \
> 输入：intervals = [[0,30],[5,10],[15,20]] \
> 输出：2
>
> 示例 2： \
> 输入：intervals = [[7,10],[2,4]] \
> 输出：1

思路：组会议室
- 对于有交集的会议，需要错开会议室
- 没有交集的会议，可以合并在同一个会议室
- 一个会议结束时间 == 另一个会议的开始时间，可以在同一个会议室
- 可以合并的会议，对应会议室更新结束时间
- 会议室，按结束时间从小到大排序
```go
func minMeetingRooms(intervals [][]int) int {
    sort.Slice(intervals, func(i,j int)bool{ // 从小到大排序
        return intervals[i][0] < intervals[j][0]
    })
    // 对于有交集的会议，需要错开会议室
    // 没有交集的会议，可以合并在同一个会议室
    // 一个会议结束时间 == 另一个会议的开始时间，可以在同一个会议室
    // 可以合并的会议，对应会议室更新结束时间
    // 会议室，按结束时间从小到大排序
    arranges := make([][]int, 0, len(intervals)) // 会议室数组
    arranges = append(arranges, intervals[0])
    for _, v := range intervals[1:]{
        if arranges[0][1] > v[0]{
            arranges = append(arranges, v)
        } else {
            for _, a := range arranges{
                if a[1] <= v[0]{ // 题目对结束时间 = 会议开始时间，认为可以同一个会议室
                    a[1] = v[1]
                    break
                }
            }
        }
        sort.Slice(arranges, func(i, j int)bool{ // 排序，对已安排的会议，按结束时间从小到大排序
            return arranges[i][1] < arranges[j][1]
        })
    }
    return len(arranges)
}
```

## [340. 至多包含 K 个不同字符的最长子串](https://leetcode-cn.com/problems/longest-substring-with-at-most-k-distinct-characters/)
给定一个字符串 s ，找出 至多 包含 k 个不同字符的最长子串 T。

> 示例 1:
> 输入: s = "eceba", k = 2
> 输出: 3
> 解释: 则 T 为 "ece"，所以长度为 3。
>
> 示例 2:
> 输入: s = "aa", k = 1
> 输出: 2
> 解释: 则 T 为 "aa"，所以长度为 2。

思路：hash法+滑动窗口计数
- hash表记录字符在当前窗口最后出现位置
- 先进先出的队列，记录字符和索引信息
  - 在出栈时（从小到大，按字符位置），不断对比与当前hash表是否相等
  - 相等则说明找到最小位置
```go
type addr struct{ b byte; i int } // b 为字符byte，i 为索引
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	n := len(s)
	if n == 0 || k == 0{
		return 0
	}
	if n <= k{
		return n
	}
	hash := make(map[byte]int) // 记录符合要求的字符在窗口最后出现的索引
	hash[s[0]] = 0
	addrs := []addr{{ // 记录位置信息
		b: s[0],
		i: 0,
	}}
	count := 1 // 记录当前窗口字符种类
	maxLen := 1 // 记录最大窗口
	for left, right := 0, 1; right < n && left < n; {
		if _, ok := hash[s[right]]; ok || count < k{ // 当前字符在符合的字符类集中、字符种类数小于k
			if !ok { // 字符种类+1
				count++
			}
			hash[s[right]] = right
			addrs = append(addrs, addr{b: s[right], i: right})
			right++
			if right - left > maxLen{
				maxLen = right-left
			}
			continue
		}
		a := addrs[0]
		addrs = addrs[1:]
		for a.i != hash[a.b]{
			a = addrs[0]
			addrs = addrs[1:]
		}
		left = a.i+1
		delete(hash, a.b)
		hash[s[right]] = right
		addrs = append(addrs, addr{b: s[right], i: right})
		right++
	}
	return maxLen
}
```