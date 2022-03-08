# 说明
本项目是基于golang刷leetcode题目，[题库](https://leetcode-cn.com/problemset/all/) \
初步刷题计划：[数据结构](https://leetcode-cn.com/study-plan/data-structures/?progress=v04mu2t) \
学习笔记：[数据结构与算法课](notes/icource.md)

目录:
- [排序算法](notes/sort.md)
- [入门](notes/getting_started.md)
- [基础](notes/base.md)
- [进阶](notes/advance.md)
- [精选200](notes/chosen.md)
- [随想录题集](notes/random.md)

![](img/structure.png) \

## 知识点

### 切片的操作
批量插入
```go
s1 := []int{0,1,2,3}
s2 := s1[0:2]
s2 = append(s2, 5,6,7) // s1: [0,1,2,3]
```
逐个扩容
```go
s1 := []int{0,1,2,3}
s2 := s1[0:2]
s2 = append(s2, 5)
s2 = append(s2, 6)
s2 = append(s2, 7) // s1: [0,1,5,6]
```

### 字符串
通过下标取值时，是byte类型（等于uint8），通过range直接赋值是int32 / rune类型（rune = int32）
```go
	a := "abc"
	var b byte
	var i32 int32
	var r rune
	for i,v := range a{
		i32 = v
		b = a[i]
		r = v
	}
```

### 反转字符串
```go
func reverserStr(x string) string {
	res := strings.Builder{}
	for i := len(x) - 1; i >= 0; i-- {
		res.WriteByte(x[i])
	}
	return res.String()
}
```


