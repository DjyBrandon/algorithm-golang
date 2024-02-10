package main

import (
	"fmt"
	"math/rand"
)

func rotate(arr []int, length, n int) []int {
	for i := 0; i < length; i++ {
		temp := arr[0]
		for i := 0; i < n-1; i++ {
			arr[i] = arr[i+1]
		}
		arr[n-1] = temp
	}
	return arr
}

func insert(arr []int, n, k int) []int {
	// 随机选择插入的位置
	index := rand.Intn((n-1)-k) + k
	// 创建足够大的目标数组
	arr1 := make([]int, n)
	// 复制插入位置之前的元素到目标数组
	copy(arr1[:index], arr[k:index+1])
	// 复制要移动的元素到目标数组的中间位置
	copy(arr1[index-k+1:index+1], arr[:k])
	// 复制插入位置之后的元素到目标数组
	copy(arr1[index+1:], arr[index+1:])
	return arr1
}

func main() {
	arr := []int{1, 2, 3, 4}

	// 输出初始数组
	fmt.Println("初始数组：", arr)

	// 1.打乱四张纸牌
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println("1.打乱四张纸牌：", arr)

	// 2.撕开纸牌
	arr = append(arr, arr...)
	fmt.Println("2.撕开纸牌（最终选出的两张牌一定 mod4 同余）：", arr)

	// 3.根据名字依次移动 N 项到末尾
	n1 := len(arr)
	// 获取名字长度
	var nameLength int
	fmt.Print("3.(1) 请输入名字长度: ")
	_, _ = fmt.Scan(&nameLength)
	arr = rotate(arr, nameLength, n1)
	fmt.Println("3.(2) 根据名字依次移动 N 项到末尾：", arr)

	// 4.取前 3 张卡片插入中间随机位置
	arr = insert(arr, n1, 3)
	fmt.Println("4.取前 3 张卡片插入中间随机位置（3<index<7）：", arr)

	// 5.取出星星值 C 藏起来，剩余列表
	star := arr[0]
	fmt.Println("5.(1) 藏起来的第一张纸牌为：", star)
	arr = append(arr[:0], arr[1:]...)
	fmt.Println("5.(2) 剩余列表为：", arr)

	// 6.取前 N（南方人 1，北方人 2，不确定 3）张牌插入中间随机位置
	// 获取地域长度
	var regionLength int
	fmt.Print("6.(1) 请选择地域 N 序号（南方人1，北方人2，不确定3）: ")
	_, _ = fmt.Scan(&regionLength)
	arr = insert(arr, len(arr), regionLength)
	fmt.Println("6.(2) 按地域取前 N 张牌插入中间随机位置（k<index<6）：", arr)

	// 7.丢弃前 N（男生 1，女生 2）张卡片后
	// 获取性别长度
	var sexLength int
	fmt.Print("7.(1) 请输入性别 N 序号（男生1，女生2）: ")
	_, _ = fmt.Scan(&sexLength)
	arr = append(arr[:0], arr[sexLength:]...)
	fmt.Println("7.(2) 按性别丢弃前 N 张卡片后：", arr)

	// 8.依次移动首张到末张 × 7
	n2 := len(arr)
	arr = rotate(arr, 7, n2)
	fmt.Println("8.依次移动首张到末张×7：", arr)

	// 9.好运留下来，烦恼丢出去 × 4
	for i := 0; i < 6-sexLength; i++ {
		arr = rotate(arr, 1, len(arr))
		arr = append(arr[:0], arr[1:]...)
	}
	fmt.Println("9.好运留下来，烦恼丢出去（女×4，男×5）：", arr)

	// 10.比较藏起来的牌与剩余最后一张牌
	fmt.Println("==================================")
	fmt.Println("10.藏起来的牌与剩余最后一张牌：", star, arr[0])
}
