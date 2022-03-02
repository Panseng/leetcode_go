# 进阶

## [325. 和等于 k 的最长子数组长度](https://leetcode-cn.com/problems/maximum-size-subarray-sum-equals-k/)
给定一个数组 nums 和一个目标值 k，找到和等于 k 的最长连续子数组长度。如果不存在任意一个符合要求的子数组，则返回 0。

> 示例 1: \
> 输入: nums = [1,-1,5,-2,3], k = 3 \
> 输出: 4 \
> 解释: 子数组 [1, -1, 5, -2] 和等于 3，且长度最长。
>
> 示例 2: \
> 输入: nums = [-2,-1,2,1], k = 1 \
> 输出: 2 \
> 解释: 子数组 [-1, 2] 和等于 1，且长度最长。

思路：前缀和
```go
func maxSubArrayLen(nums []int, k int) int {
    ans := 0
    sum := 0
    n := len(nums)
    sm := make(map[int]int, n) // 记录和的下标
    sm[0] = -1 // 和为0，在左侧
    for i := 0; i < n; i++{
        sum += nums[i]
        if j,ok := sm[sum-k]; ok && i-j > ans{
            ans = i-j
        }
        if _,ok := sm[sum];!ok{ // 这类求的是最大子数组，因此因尽量保持下标在左侧
            sm[sum] = i
        }
    }
    return ans 
}
```

## [1151. 最少交换次数来组合所有的 1](https://leetcode-cn.com/problems/minimum-swaps-to-group-all-1s-together/)
给出一个二进制数组 data，你需要通过交换位置，将数组中 任何位置 上的 1 组合到一起，并返回所有可能中所需 最少的交换次数。

> 示例 1: \
> 输入: data = [1,0,1,0,1] \
> 输出: 1 \
> 解释: \
> 有三种可能的方法可以把所有的 1 组合在一起：
> - [1,1,1,0,0]，交换 1 次；
> - [0,1,1,1,0]，交换 2 次；
> - [0,0,1,1,1]，交换 1 次。
> - 所以最少的交换次数为 1。
>
> 示例  2: \
> 输入：data = [0,0,0,1,0] \
> 输出：0 \
> 解释： \
> 由于数组中只有一个 1，所以不需要交换。
>
> 示例 3: \
> 输入：data = [1,0,1,0,1,0,0,1,1,0,1] \
> 输出：3 \
> 解释： \
> 交换 3 次，一种可行的只用 3 次交换的解决方案是 [0,0,0,0,0,1,1,1,1,1,1]。
>
> 示例 4: \
> 输入: data = [1,0,1,0,1,0,1,1,1,0,1,0,0,1,1,1,0,0,1,1,1,0,1,0,1,1,0,0,0,1,1,1,1,0,0,1] \
> 输出: 8

思路：滑动窗口
```go
func minSwaps(data []int) int {
   c := 0; // 1的总数
   n := len(data)
   for i := 0; i < n; i++{
       c += data[i]
   }
   if c <= 1{
       return 0
   }
   ans := c-1 // 最少移动的次数（这个是最大值，最多移动次数）
   cw := 0 // 窗口，以 c 为长度的窗口
   for i := 0; i < c; i++{ // 0-c 窗口中，1的个数
       cw += data[i]
   }
   if c-cw < ans{ // 取移动次数的最小值
       ans = c-cw
   }
   for i := 1; i+c-1 < n; i++{
       cw += data[i+c-1]-data[i-1] // 窗口移动，则窗口中1的次数变更
       if c-cw < ans{
           ans = c-cw
       }
   }
   return ans
}
```

## [1588. 所有奇数长度子数组的和](https://leetcode-cn.com/problems/sum-of-all-odd-length-subarrays/)
给你一个正整数数组 arr ，请你计算所有可能的奇数长度子数组的和。 \ 
子数组 定义为原数组中的一个连续子序列。 \
请你返回 arr 中 所有奇数长度子数组的和 。

> 示例 1： \
> 输入：arr = [1,4,2,5,3] \
> 输出：58 \
> 解释：所有奇数长度子数组和它们的和为： \
> [1] = 1 \
> [4] = 4 \
> [2] = 2 \
> [5] = 5 \
> [3] = 3 \
> [1,4,2] = 7 \
> [4,2,5] = 11 \
> [2,5,3] = 10 \
> [1,4,2,5,3] = 15 \
> 我们将所有值求和得到 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58
>
> 示例 2： \
> 输入：arr = [1,2] \
> 输出：3 \
> 解释：总共只有 2 个长度为奇数的子数组，[1] 和 [2]。它们的和为 3 。
>
> 示例 3： \
> 输入：arr = [10,11,12] \
> 输出：66

思路：暴力求解
```go
func sumOddLengthSubarrays(arr []int) int {
    sum := 0
    n := len(arr)
    for l:= 1; l <= n; l+=2{
        for i := 0; i+l-1 < n; i++{
            for j := i; j < i+l; j++{
                sum += arr[j]
            }
        }
    }
    return sum
}
```

思路2：前缀和
```go
func sumOddLengthSubarrays(arr []int) int {
    n := len(arr)
    prefixSum := make([]int, n+1)
    prefixSum[0] = 0 // 第0位的前缀和为0
    for i,v := range arr{
        prefixSum[i+1] = prefixSum[i]+v
    }
    
    sum := 0
    for start := range arr{
        for length := 1; start+length-1 < n; length += 2{
            end := start+length-1
            sum += prefixSum[end+1]-prefixSum[start] // 当前区间的和 = 右侧前缀和-左侧前缀和
        }
    }
    return sum
}
```

## [452. 用最少数量的箭引爆气球](https://leetcode-cn.com/problems/minimum-number-of-arrows-to-burst-balloons/)
在二维空间中有许多球形的气球。对于每个气球，提供的输入是水平方向上，气球直径的开始和结束坐标。由于它是水平的，所以纵坐标并不重要，因此只要知道开始和结束的横坐标就足够了。开始坐标总是小于结束坐标。 \
一支弓箭可以沿着 x 轴从不同点完全垂直地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被引爆。可以射出的弓箭的数量没有限制。 弓箭一旦被射出之后，可以无限地前进。我们想找到使得所有气球全部被引爆，所需的弓箭的最小数量。 \
给你一个数组 points ，其中 points [i] = [xstart,xend] ，返回引爆所有气球所必须射出的最小弓箭数。

> 示例 1： \
> 输入：points = [[10,16],[2,8],[1,6],[7,12]] \
> 输出：2 \
> 解释：对于该样例，x = 6 可以射爆 [2,8],[1,6] 两个气球，以及 x = 11 射爆另外两个气球
>
> 示例 2： \
> 输入：points = [[1,2],[3,4],[5,6],[7,8]] \
> 输出：4
>
> 示例 3： \
> 输入：points = [[1,2],[2,3],[3,4],[4,5]] \
> 输出：2
>
> 示例 4： \
> 输入：points = [[1,2]] \
> 输出：1
>
> 示例 5： \
> 输入：points = [[2,3],[2,3]] \
> 输出：1

思路：排序+贪心
- 注意，这类仅能合并存在公共交集的区间，所有被合并的区间均存在该交集
  - 如：`[1,5],[2,4],[4,5]`，合并前公共交集为`4`
  - 如：`[1,5],[2,3],[4,5]`，则仅前两项存在公共交集`[2,3]`
```go
func findMinArrowShots(points [][]int) int {
    n := len(points)
    if n < 2{
        return n
    }
    // 排序，按左侧从小到达排序
    sort.Slice(points, func(i,j int)bool{
        return points[i][0] < points[j][0]
    })
    var leftEnd int
    for i, x := range points{
        if i == 0{
            leftEnd =points[0][1]
            continue
        }
        if x[0] <= leftEnd{
            n--
            // 注意，这类并不是要合并所有有重叠的区间
            // 而是，合并的区间，合并前必须存在共同的交集
            // 如：[1,5],[2,4],[4,5]，合并前公共交集为4
            // 如：[1,5],[2,3],[4,5]，则仅前两项存在公共交集[2,3]
            if x[1] < leftEnd{
                leftEnd = x[1]
            }
            continue
        }
        leftEnd = x[1]
    }
    return n
}
```

思路2：排序+贪心
- 按右侧排序，按照右边界排序，就要从左向右遍历，因为右边界越小越好，只要右边界越小，留给下一个区间的空间就越大，所以从左向右遍历，优先选右边界小的。
- 同 [435. 无重叠区间](https://leetcode-cn.com/problems/non-overlapping-intervals/solution/435-wu-zhong-die-qu-jian-tan-xin-jing-di-qze0/)
```go
func findMinArrowShots(points [][]int) int {
	if len(points) < 2{
		return len(points)
	}
	// 以右侧端点排序
	sort.Slice(points, func(i, j int) bool{
		return points[i][1] < points[j][1]
	})
	// 从左往右，贪心算法
	num, end := 1, points[0][1]
	for _,v := range points[1:]{
		if v[0] > end{
			num++
			end = v[1]
		}
	}
	return num
}
```

## [128. 最长连续序列](https://leetcode-cn.com/problems/longest-consecutive-sequence/)

给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。 \
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

> 示例 1： \
> 输入：nums = [100,4,200,1,3,2] \
> 输出：4 \
> 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
>
> 示例 2： \
> 输入：nums = [0,3,7,2,5,8,4,6,0,1] \
> 输出：9

思路：排序(nlogn)+计数（取最大值）
```go
func longestConsecutive(nums []int) int {
    n := len(nums)
    if n < 1{
        return 0
    }
    sort.Ints(nums)
    ans := 1
    s := 0
    
    for i := 1; i < n; i++{
        if nums[i] == nums[i-1]{ // 删除重复项，避免累计错误
            if i == 1{
                nums = nums[1:]
            } else if i == n-1{
                nums = nums[:i]
            } else{
                nums = append(nums[:i], nums[i+1:]...)
            }
            i--
            n--
            continue
        }
        if nums[i] == nums[i-1]+1{
            s++
            if s+1 > ans{
                ans = s+1
            }
        } else{
            s = 0
        }
    }
    return ans
}
```

思路2：hash法，对每个头部数据进行递增匹配，找最大长度
```go
func longestConsecutive(nums []int) int {
    numMap := make(map[int]bool)
    for _,v := range nums{
        numMap[v] = true
    }
    ans := 0
    for k := range numMap{
        if ans > len(nums)/2{ // 如果过半数，则不可能产生更大的连续序列了
            break
        }
        // 保证只从最小值开始寻找，一直找到最大值，来保障最大长度
        // 同时避免重复遍历子集
        if _, okP := numMap[k-1]; !okP{ 
            next := k+1
            for numMap[next]{ // 通过迭代找最大
                next++
            }
            if next-k > ans{
                ans = next-k
            }
        }
    }
    return ans
}
```

## [454. 四数相加 II](https://leetcode-cn.com/problems/4sum-ii/)
给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：
- 0 <= i, j, k, l < n
- nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

> 示例 1： \
> 输入：nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2] \
> 输出：2 \
> 解释： \
> 两个元组如下：
> 1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
> 2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0
>
> 示例 2： \
> 输入：nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0] \
> 输出：1

思路：hash法，记住前2项和、后2项和，然后分别比较两项和是否构成和为0，符合条件则累计次数
```go
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    n := len(nums1)
    sumOT := make(map[int]int)
    for i := 0; i < n; i++{ // 前两项和
        for j := 0; j < n; j++{
            sumOT[nums1[i]+nums2[j]]++
        }
    }

    sumTF := make(map[int]int)
    for i := 0; i < n; i++{ // 后两项和
        for j := 0; j < n; j++{
            sumTF[nums3[i]+nums4[j]]++
        }
    }

    ans := 0
    for kOT, vOT := range sumOT{ // 匹配
        if vTF, ok := sumTF[-kOT]; ok{
            ans += vOT*vTF
        }
    }
    return ans
}
```
优化，减少迭代次数：在获取后两项和时，直接匹配前两项和
```go
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
    n := len(nums1)
    sumOT := make(map[int]int)
    for i := 0; i < n; i++{
        for j := 0; j < n; j++{
            sumOT[nums1[i]+nums2[j]]++
        }
    }

    ans := 0
    for i := 0; i < n; i++{
        for j := 0; j < n; j++{
            if v,ok := sumOT[-nums3[i]-nums4[j]]; ok{ // 直接匹配后两项与前两项和
                ans += v
            }
        }
    }
    
    return ans
}
```

## [448. 找到所有数组中消失的数字](https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/)
给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

> 示例 1： \
> 输入：nums = [4,3,2,7,8,2,3,1] \
> 输出：[5,6]
>
> 示例 2： \
> 输入：nums = [1,1] \
> 输出：[2]

思路：通过数组记录出现次数
- 在可以产生连续的排列时，数组下标可以作为特殊的key值
- 比如连续的大写字母、小写字母、数值（注意，大小写字母直接有空隙，需要排除）
```go
func findDisappearedNumbers(nums []int) []int {
    nm := make([]int, len(nums)+1)
    for _,v := range nums{
        nm[v]++
    }
    ans := []int{}
    for i,v := range nm{
        if i == 0 {
            continue
        }
        if v == 0{
            ans = append(ans, i)
        }
    }
    return ans
}
```

思路2：原位操作
- 数组中的值是乱序
- 数组中的值在 1-n之间，可以构成下标（需要v-1）
- 值指向的下标，对应值需要+n，来区别没有值对应的下标值
- +n是便于取模定位
```go
func findDisappearedNumbers(nums []int) []int {
    n := len(nums)
    // 注意，nums是乱序
    for _,v := range nums{
        // 需要将值对应到长度为 n 的数组，那么 v 作为下标，需要-1
        v = (v-1)%n // 获取 v值 的前一项（此次 v 被看着下标，取n的余数，是因为v可能大于n
        nums[v] += n // 此次操作，使得 v 可能大于 n
    }
    ans := []int{}
    for i,v := range nums{
        if v <= n{ // v
            ans = append(ans, i+1)
        }
    }
    return ans
}
```

## [1427. 字符串的左右移](https://leetcode-cn.com/problems/perform-string-shifts/)
给定一个包含小写英文字母的字符串 s 以及一个矩阵 shift，其中 shift[i] = [direction, amount]：
- direction 可以为 0 （表示左移）或 1 （表示右移）。
- amount 表示 s 左右移的位数。
- 左移 1 位表示移除 s 的第一个字符，并将该字符插入到 s 的结尾。
- 类似地，右移 1 位表示移除 s 的最后一个字符，并将该字符插入到 s 的开头。

对这个字符串进行所有操作后，返回最终结果。

> 示例 1： \
> 输入：s = "abc", shift = [[0,1],[1,2]] \
> 输出："cab" \
> 解释： \
> [0,1] 表示左移 1 位。 "abc" -> "bca" \
> [1,2] 表示右移 2 位。 "bca" -> "cab"
>
> 示例 2： \
> 输入：s = "abcdefg", shift = [[1,1],[1,1],[0,2],[1,3]] \
> 输出："efgabcd" \
> 解释： \
> [1,1] 表示右移 1 位。 "abcdefg" -> "gabcdef" \
> [1,1] 表示右移 1 位。 "gabcdef" -> "fgabcde" \
> [0,2] 表示左移 2 位。 "fgabcde" -> "abcdefg" \
> [1,3] 表示右移 3 位。 "abcdefg" -> "efgabcd"

思路：先获取最终左右移动的情况
```go
func stringShift(s string, shift [][]int) string {
    amount := 0
    for _, v := range shift{
        if v[0] == 0{
            amount += v[1]
        } else{ // 右移转换为负值
            amount -= v[1]
        }
    }
    amount = amount % len(s) // 可能会存在移动次数超过字符串长度的情况
    if amount == 0{
        return s
    } else if amount > 0{
        return s[amount:] + s[:amount]
    }
    return s[len(s)+amount:] + s[:len(s)+amount]
}
```
