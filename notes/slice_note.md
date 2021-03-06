Go 中使用 append 多次操作并赋值 slice，为什么原始值可能被改变？
在使用 append 操作 slice 赋值给新 slice 时，可能会遇到中间赋值的 slice 值被改变的非预期情况，比如说如下的代码：
```go
func main(){
    s := []int{5}
    s = append(s, 7)
    s = append(s, 9)
    x := append(s, 11)
    y := append(s, 12)
   fmt.Println(s, x, y)
}
```
实际输出的结果是 [5 7 9] [5 7 9 12] [5 7 9 12]，并非预期的 [5 7 9] [5 7 9 11] [5 7 9 12]

比较敏感的同学应该一下会猜到跟内存地址有关系，实际上这种情况也是对 slice 和 append 的特性不够了解造成的。

首先我们要知道 slice 数据结构包含两个内部变量 len（已有数据长度）和 cap（总共可容纳数据量）。当使用 append 操作 slice 时，如果 cap > len，即容量够用，便会继续使用当前的空间，那么内存地址不会变化。如果 cap == len，空间不够用了，就会对 slice 进行自动扩容，cap 变为两倍，此时会申请新空间，也就是返回了新的 slice，内存地址也就发生了变化。

在上述例子中，初始的 slice 大小为 1, 在第二、三步时都需要扩容。而第四、五步都是以 s 为基础，容量够用，因此没有扩容申请新空间，因此便将同一个地址赋给了 x 和 y。

<table>
<thead>
<tr>
<th>动作</th>
<th>地址</th>
<th>数据</th>
<th>cap</th>
</tr>
</thead>
<tbody>
<tr>
<td>s := []int{5}</td>
<td>0x1111</td>
<td>5</td>
<td>1</td>
</tr>
<tr>
<td>s = append(s, 7)</td>
<td>0x2222</td>
<td>5,7</td>
<td>2</td>
</tr>
<tr>
<td>s = append(s, 9)</td>
<td>0x3333</td>
<td>5,9</td>
<td>4</td>
</tr>
<tr>
<td>x := append(s, 11)</td>
<td>0x2222</td>
<td>5,7，9，11</td>
<td>2</td>
</tr>
<tr>
<td>y := append(s, 12)</td>
<td>0x2222</td>
<td>5,7，9，12</td>
<td>2</td>
</tr>
</tbody>
</table>

因为四五步操作的是同一个地址上的数据，因此 x 的值自然就被 y 的值覆盖了。

而如果是用下面的代码，s 的大小初始为 3，那么每次操作都需要扩容，都会返回新空间，那么结果就符合我们的预期了。
```go
func main() {
	s := []int{5, 7, 9}
	x := append(s, 11)
	y := append(s, 12)
	fmt.Println(s, x, y)
}
```

如何避免或解决这一问题呢

其实解决办法也很简单，就是当要赋值给新的 slice 的时候，都重新申请空间并使用 copy 复制就可以了。类似如下代码：
```go
func main(){
	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)
	x := make([]int, len(s))
	copy(x, s)
	x = append(x, 11)
	y := append(s, 12)
	fmt.Println(s, x, y)
}
```
最后可能有人会有疑问：既然 x 和 y 都是使用的和 s 相同的内存，那为什么打印出的 s 和 x、y 结果不一样呢？

实际上虽然指向同一块内存，但是他们的 len 是不同的，所以打印出来的结果是不一样的