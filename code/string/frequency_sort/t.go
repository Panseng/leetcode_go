package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	var big int
	big = 10000*10000
	fmt.Println(big)
	fmt.Println("z"[0]-'a',"B"[0]-'a')
	//fmt.Println(frequencySort("tree"))
	//fmt.Println(frequencySort("cccaaa"))
	//fmt.Println(frequencySort("Aabb"))

	fmt.Println(frequencySort4("tree"))
	fmt.Println(frequencySort4("cccaaa"))
	fmt.Println(frequencySort4("Zabb"))

	//a := "abc"
	//var b byte
	//var i32 int32
	//var r rune
	//for i,v := range a{
	//	i32 = v
	//	b = a[i]
	//	r = v
	//}
}

func frequencySort4(s string) string {
	cnt := make(map[byte]int)
	maxFre := 0 // 最高出现频率
	for i := range s{
		cnt[s[i]]++
		if cnt[s[i]] > maxFre{
			maxFre = cnt[s[i]]
		}
	}

	buckets := make([][]byte, maxFre+1) // 记录各个频率出现的字母byte值
	for k,v := range cnt{
		buckets[v] = append(buckets[v], k)
	}

	ans := make([]byte, 0, len(s))
	for i := maxFre; i > 0; i--{
		for n := 0; n < len(buckets[i]); n++{
			ans = append(ans,bytes.Repeat([]byte{buckets[i][n]}, i)...)
		}
	}
	return string(ans)
}

func frequencySort3(s string) string {
	cnt := make(map[byte]int)
	for i := range s{
		cnt[s[i]]++
	}

	type ac struct {
		b byte
		c int
	}
	cts := []ac{}
	for k,v := range cnt{
		cts = append(cts, ac{k,v})
	}

	sort.Slice(cts, func(i, j int) bool {
		return cts[i].c > cts[j].c
	})
	ans := make([]byte, 0, len(s))
	for _,v := range cts{
		ans = append(ans,bytes.Repeat([]byte{v.b}, v.c)...)
	}
	return string(ans)
}

// 未考虑数字的情况
func frequencySort2(s string) string {
	type ac struct {
		i int32
		c int
	}
	alps := make([]ac, 26*2)
	for i := range alps{
		alps[i] = ac{c:0}
	}
	for _,v := range s{
		i := v-'a'
		if i < 0{
			i += 32+26 // 对于大写字母，定位到小写字母后的数组
		}
		if alps[i].c == 0{
			alps[i].i = v
		}
		alps[i].c++
	}

	sort.Slice(alps, func(i, j int) bool {
		return alps[i].c > alps[j].c // 从大到小
	})
	ans := ""
	for n := 0; n < 52; n++{
		if alps[n].c == 0{
			break
		}
		for m := 0; m < alps[n].c; m++{
			ans += string(alps[n].i)
		}
	}

	return ans
}

func frequencySort(s string) string {
	n := len(s)
	ac := make(map[int32]int, n)
	for _,v := range s{
		ac[v]++
	}

	ah := &alphaH{}
	heap.Init(ah) // 初始化堆
	for k,v := range ac{
		heap.Push(ah, alphaC{k,v})
	}

	ans := ""
	for ah.Len() > 0{
		tem := heap.Pop(ah).(alphaC) // 使用堆时注意转换 pop 出来的格式
		for i := 0; i < tem.c;i++{
			ans += string(tem.i)
		}
	}

	return ans
}

type alphaC struct {
	i int32 // 字符串 int32 码
	c int  // 字符串出现次数
}

type alphaH []alphaC

func (a alphaH) Len() int { return len(a) }
func (a alphaH) Less(i, j int) bool { return a[i].c > a[j].c } // 当前是从小到大排序，如果要从大到小，则 a[i].c < a[j].c
func (a alphaH) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a *alphaH) Push(x interface{}) {
	*a = append(*a, x.(alphaC))
}

func (a *alphaH) Pop() interface{} {
	old := *a
	n := len(old)
	x := old[n-1]
	*a = old[0 : n-1]
	return x
}