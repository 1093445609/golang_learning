package main

import "fmt"

//匿名函数用作回调函数
func main() {
	//传入 两个参数 一个数组 一个是匿名函数
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})

}

// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}
