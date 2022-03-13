# hash表

## [217. 存在重复元素](https://leetcode-cn.com/problems/contains-duplicate/solution/cun-zai-zhong-fu-yuan-su-by-leetcode-sol-iedd/)
给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
>示例1\
>输入：nums = [1,2,3,1]\
>输出：true

思路1：先用sort.Int排序，然后比较相邻元素是否相同\
```if nums[i] == nums[i-1]```\
思路2：hash表，判断hash表中是否存在该值，存在返回真，不存在则添加\
实现：[arr_repeat](../code/array/arr_repeat)

## [706. 设计哈希映射](https://leetcode-cn.com/problems/design-hashmap/submissions/)
不使用任何内建的哈希表库设计一个哈希映射（HashMap）。

实现 MyHashMap 类：
- MyHashMap() 用空映射初始化对象
- void put(int key, int value) 向 HashMap 插入一个键值对 (key, value) 。如果 key 已经存在于映射中，则更新其对应的值 value 。
- int get(int key) 返回特定的 key 所映射的 value ；如果映射中不包含 key 的映射，返回 -1 。
- void remove(key) 如果映射中存在 key 的映射，则移除 key 和它所对应的 value 。

思路1：简单粗暴
```go
type MyHashMap struct {
	Val [][2]int
}


func Constructor() MyHashMap {
	return MyHashMap{[][2]int{}}
}


func (this *MyHashMap) Put(key int, value int)  {
	if i, ok := this.isExist(key); ok{
		this.Val[i][1] = value
	} else{
		this.Val = append(this.Val,[2]int{key, value})
	}
}


func (this *MyHashMap) isExist(key int) (index int, here bool) {
	here = false
	index = -1
	for i, v := range this.Val{
		if v[0] == key{
			here = true
			index = i
			break
		}
	}
	return
}


func (this *MyHashMap) Get(key int) int {
	for _, v := range this.Val{
		if v[0] == key{
			return v[1]
		}
	}
	return -1
}


func (this *MyHashMap) Remove(key int)  {
	if i, ok := this.isExist(key); ok{
		if len(this.Val)>i+1{
			this.Val = append(this.Val[:i], this.Val[i+1:]...)
		} else{
			this.Val = this.Val[:i]
		}
	}
}
```
思路2：仅hash表
```go
type MyHashMap struct {
    Val map[int]int
}


func Constructor() MyHashMap {
    return MyHashMap{}
}


func (this *MyHashMap) Put(key int, value int)  {
    if this.Val == nil{
        this.Val = make(map[int]int)
    }
    this.Val[key] = value
}


func (this *MyHashMap) Get(key int) int {
    if _, ok := this.Val[key]; ok{
        return this.Val[key]
    }
    return -1
}


func (this *MyHashMap) Remove(key int)  {
    if _, ok := this.Val[key]; !ok{ // 不存在则跳出
        return
    }
    delete(this.Val,key)
}
```