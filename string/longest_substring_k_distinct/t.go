package main

import "fmt"

func main()  {
	fmt.Println(lengthOfLongestSubstringKDistinct("eceba", 2))
	fmt.Println(lengthOfLongestSubstringKDistinct("aa", 1))
}

// 340. 至多包含 K 个不同字符的最长子串
// 最小堆+字典+滑动窗口
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

//type hp []addr
//
//func (h hp)Len() int{return len(h)}
//func (h hp)Less(i, j int) bool {return h[i].i < h[j].i}
//func (h hp)Swap(i,j int){h[i], h[j] = h[j], h[i]}
//func (h *hp) Push(v interface{}){*h = append(*h, v.(addr))}
//func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[1:], a[0]; return }

