## [4. 寻找两个正序数组的中位数]

给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。 \
算法的时间复杂度应该为 O(log (m+n)) 。

> 示例 1： \
> 输入：nums1 = [1,3], nums2 = [2] \
> 输出：2.00000 \
> 解释：合并数组 = [1,2,3] ，中位数 2
>
> 示例 2： \
> 输入：nums1 = [1,2], nums2 = [3,4] \
> 输出：2.50000 \
> 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

思路：遍历
```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n1 := len(nums1)
    n2 := len(nums2)
    if n1 == 0 && n2 == 0{
        return float64(0)
    }
    isOdd := (n1+n2)%2 == 1
    var left, right float64
    mid := (n1+n2)/2
    i := 0
    for ; len(nums1) > 0 && len(nums2) > 0; i++{
        var num int
        if nums1[0] > nums2[0]{
            num = nums2[0]
            nums2 = nums2[1:]
        } else{
            num = nums1[0]
            nums1 = nums1[1:]
        }
        if isOdd && i == mid{
            return float64(num)
        }
        if !isOdd && i == mid-1{
            left = float64(num)
        }
        if !isOdd && i == mid{
            right = float64(num)
            i++
            break
        }
    }
    if i > mid{ // 说明已经取到值
        return (left+right)/2
    }
    if len(nums2) > 0{ // 未取到值，其中一个数组长度不够
        nums1 = nums2
    }
    if i == mid{ // 刚好卡在中间站
        if isOdd{ // 奇数长度时
            return float64(nums1[0])
        }
        return (left + float64(nums1[0]))/2
    }
    for ; len(nums1) > 0; i++{
        num := nums1[0]
        nums1 = nums1[1:]
        
        if isOdd && i == mid{
            return float64(num)
        }
        if !isOdd && i == mid-1{
            left = float64(num)
        }
        if !isOdd && i == mid{
            right = float64(num)
            break
        }
    }
    return (left+right)/2
}
```
思路2：合并后排序
```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n1 := len(nums1)
    n2 := len(nums2)
    if n1 < n2{
        nums1, nums2 = nums2, nums1
    }
    nums1 = append(nums1, nums2...)
    n1 = n1+n2
    sort.Ints(nums1)
    
    isOdd := n1%2 == 1
    var left, right float64
    mid := n1/2
    for i := 0; i <= mid; i++{
        if isOdd && i == mid{
            return float64(nums1[i])
        }
        if !isOdd && i == mid-1{
            left = float64(nums1[i])
        }
        if !isOdd && i == mid{
            right = float64(nums1[i])
        }
    }
    return (left+right)/2
}
```