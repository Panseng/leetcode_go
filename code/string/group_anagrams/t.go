package main

import (
	"fmt"
	"sort"
)

func main()  {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	//strs := []string{""}
	//strs := []string{"a"}
	//fmt.Println(groupAnagrams(strs))
	fmt.Println(groupAnagrams2(strs))
}

func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	mp := make(map[string]int, len(strs))
	for _, v := range strs{
		a := []byte(v)
		sort.Slice(a, func(i, j int) bool {
			return a[i] < a[j]
		})
		sv := string(a)
		if i, ok := mp[sv]; ok{
			ans[i] = append(ans[i], v)
		} else {
			mp[sv] = len(ans)
			ans = append(ans, []string{v})
		}
	}
	return ans
}

func groupAnagrams2(strs []string) [][]string {
	ans := [][]string{}
	mp := make(map[[26]int]int, len(strs))
	for _, v := range strs{
		cn := [26]int{}
		for _, a := range v{
			cn[a-'a']++
		}
		if i, ok := mp[cn]; ok{
			ans[i] = append(ans[i], v)
		} else {
			mp[cn] = len(ans)
			ans = append(ans, []string{v})
		}
	}
	return ans
}
