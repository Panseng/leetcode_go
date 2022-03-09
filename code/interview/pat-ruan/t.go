package main

import (
	"fmt"
	"math/rand"
)

/*
机试题目注意: 请不要轻视这题目，完成得高质量，绝对不容易。
请独立完成，其他注意事项见下方 FAQ。
请将完成的笔试题以“姓名-后端开发-完成所需时间“格式命名
要求：以下题目答题是 30 分钟,旨在考察面试者多方面的能力。请尽可能多答，
算法题可以任何编程语言的代码（欢迎加关键注释），也可以写出大概想法（得分会减少）。
题目：我们做了一个活动，根据用户的积分来抽奖，用户的积分都保存在一个数组里面
arr = [20, 34, 160, 2…]，数组下标就是用户的 ID，则这里：
ID 为 0 的用户的积分是 arr[0] 等于 20 分。
ID 为 1 的用户的积分是 arr[1] 等于 34 分。
请你设计一个抽奖算法，随机抽出一位中奖用户，要求积分越高中奖概率越高。
返回值是中奖用户的 ID
PS: 1<= arr.length <= 50000 且 1<= arr[i] <= 50000
代码写出算法，
并分析其时间复杂度，
为其编写尽量多 unit test。
FAQ：
我可以上网吗？－－ 可以，make yourself comfortable。
我可以问别人吗？ －－ 请独立完成，if you lie , we’ll know sooner or later。
我超过 30 分钟怎么办？请尽量按时提交。如果超过 30 分钟，请标注下完成用时。
我做不完怎么办？没关系请尽量按点顺序完成
完成后，可以发到邮箱：bhruan@riches.ai
 */
// 由最大值构成概率区间，确保分数高的被命中概率大
// 被命中的再次随机选中最终幸运客户
// 时间复杂度：O(n)，2次遍历的时间
// 空间复杂度：O(n)，最糟糕的情况是所有用户均被命中
func lottery(userInt []int) int {
	bigest := 0
	for _, v := range userInt{ // 找最大值即可代替区间
		if v > bigest{
			bigest = v
		}
	}
	luckyNum := rand.Intn(bigest+1)
	luckys := []int{}
	for i, v := range userInt{
		if luckyNum <= v{
			luckys = append(luckys, i)
		}
	}
	luckyNum = rand.Intn(len(luckys))
	return luckys[luckyNum]
}

func main()  {
	arr := []int{120,34,160,2}
	res := []int{}
	for i := 10; i >= 0; i--{
		res = append(res, lottery(arr))
	}
	fmt.Println(res)

	arr = []int{0,0,0,2}
	res = []int{}
	for i := 10; i >= 0; i--{
		res = append(res, lottery(arr))
	}
	fmt.Println(res)

	arr = []int{0,0,0,2,0,0,0,0,0,0,1,1,1,0,0,1,5,2,5,6}
	res = []int{}
	for i := 10; i >= 0; i--{
		res = append(res, lottery(arr))
	}
	fmt.Println(res)

	arr = []int{0,0,0,10,0,0,0,0,0,0,1,1,1,0,0,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}
	res = []int{}
	for i := 10; i >= 0; i--{
		res = append(res, lottery(arr))
	}
	fmt.Println(res)

	arr = []int{0,0,0,15,0,0,0,0,0,0,1,1,1,0,0,1,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,5,21,3,0,0,0, 5,6,8,9}
	res = []int{}
	for i := 10; i >= 0; i--{
		res = append(res, lottery(arr))
	}
	fmt.Println(res)
}

