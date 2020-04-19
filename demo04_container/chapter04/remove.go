package main

import "fmt"

/*从切片中删除元素,使用切片本身的特性来删除元素*/
func main() {
	a := make([]int, 3)

	a = []int{1, 2, 3}
	//从开头位置删除
	//1. a[n:]  删除开头n个元素
	fmt.Println(a[1:])

	//2.也可以不移动数据指针，但是将后面的数据向开头移动，可以用 append 原地完成
	a = append(a[:0], a[1:]...)
	fmt.Println(a)
	//a = append(a[:0], a[N:]...) // 删除开头N个元素   append到从哪开始的

	//3. copy() 函数来删除开头的元素

	a = a[:copy(a, a[1:])] // 删除开头1个元素
	//a = a[:copy(a, a[N:])] // 删除开头N个元素
	fmt.Println(a)

	//从中间位置删除
	//	同样可以用 append 或 copy 原地完成
	a = []int{1, 2, 3, 4, 5, 6}
	a = append(a[:1], a[2:]...) // 删除中间1个元素  删除的是第二个元素
	fmt.Println(a)
	//a = append(a[:i], a[i+N:]...) // 删除中间N个元素
	//	a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
	//a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素

	//	从尾部删除

	a = []int{1, 2, 3}
	a = a[:len(a)-1] // 删除尾部1个元素
	//a = a[:len(a)-N] // 删除尾部N个元素

}
