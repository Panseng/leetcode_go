# 知识点

## mod replace
如果引用完整的路径，而项目还没推送到远程，是没办法下载的，因此需要用 replace 将其指定读取本地的模块路径，这样子就可以解决本地模块读取的问题。
```mod
module github.com/EDDYCJY/go-gin-example
go 1.13
require (...)
replace (
    github.com/EDDYCJY/go-gin-example/pkg/setting => ~/go-application/go-gin-example/pkg/setting
    github.com/EDDYCJY/go-gin-example/routers 	  => ~/go-application/go-gin-example/routers
)
```

## fmt
`Sprintf` 则格式化并返回一个字 符串而不带任何输出。举例： \
`s := fmt.Sprintf("是字符串 %s ","string")`

 `Fprintf` 用于格式化并输出：
`fmt.Fprintf(os.Stderr, "格式化 %s\n", "error")`

## 切片的操作
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

## 字符串
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

## 反转字符串
```go
func reverserStr(x string) string {
	res := strings.Builder{}
	for i := len(x) - 1; i >= 0; i-- {
		res.WriteByte(x[i])
	}
	return res.String()
}
```

## fmt占位符
### 通用verbs
| | |
| --- | --- |
%v | 值的默认格式
%+v | 类似%v,但输出结构体时会添加字段名
%#v | Go语法表示值
%T | Go语法表示类型
%% | 百分号表示
```go
	type Sample struct {
		Title string
		name  string
		Age   int
	}
	s := Sample{"测试", "wentao", 26}
	fmt.Printf("%v\n", s)       // {测试 wentao 26}
	fmt.Printf("%+v\n", s)      // {Title:测试 name:wentao Age:26}
	fmt.Printf("%#v\n", s)      // main.Sample{Title:"测试", name:"wentao", Age:26}
	fmt.Printf("%T\n", s)       //  main.Sample
	fmt.Printf("%v%%\n", s.Age) //  26%
```
### bool值
```go
fmt.Printf("%t\n", true) // true
```
###  string与[]byte
| | |
| --- | --- |
%s | 输出字符串表示（string类型或[]byte)
%q | 双引号围绕的字符串，由Go语法安全地转义
%x | 十六进制，小写字母，每字节两个字符 （使用a-f）
%X | 十六进制，大写字母，每字节两个字符 （使用A-F）
```go
	fmt.Printf("%s\n", []byte("go开发"))      //go开发
	fmt.Printf("%s\n", "go开发")          //go开发
	fmt.Printf("%q\n", []byte("go开发"))          //"go开发"
	fmt.Printf("%q\n", "go开发")          //"go开发"
	fmt.Printf("%x\n", "go开发")          //676fe5bc80e58f91
	fmt.Printf("%X\n", "go开发")          //676FE5BC80E58F91
```
### 整数
| | |
| --- | --- |
%b | 表示二进制
%c | 该值对应的unicode吗值
%d | 表示十进制
%o | 表示八进制
%q | 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
%x | 表示为十六进制，使用a-f
%X | 表示为十六进制，使用A-F
%U | 表示为Unicode格式：U+1234，等价于"U+%04X"
```go
	fmt.Printf("%b\n", 11)     // 1011
	fmt.Printf("%c\n", 0x4E2D) // 中
	fmt.Printf("%d\n", 0x4E2D) // 20013
	fmt.Printf("%o\n", 0x4E2D) // 47055
	fmt.Printf("%q\n", 0x4E2D) // '中'
	fmt.Printf("%x\n", 0x4E2D) // 4e2d
	fmt.Printf("%X\n", 0x4E2D) // 4E2D
	fmt.Printf("%U\n", 0x4E2D) // U+4E2D
```
### 浮点数与复数
| | |
| --- | --- |
%b | 无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat
%e | 科学计数法，例如 -1234.456e+78
%E | 科学计数法，例如 -1234.456E+78
%f | 有小数点而无指数，例如 123.456
%F | 等价于%f
%g | 根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
%G | 根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
```go
	fmt.Printf("%b\n", 10.45) //5882827013252710p-49
	fmt.Printf("%e\n", 10.45) //1.045000E+01
	fmt.Printf("%E\n", 10.45) //1.045000E+01
	fmt.Printf("%f\n", 10.45) //10.450000
	fmt.Printf("%F\n", 10.45) //10.450000
	fmt.Printf("%g\n", 10.45) //10.45
	fmt.Printf("%G\n", 10.45) //10.45
```