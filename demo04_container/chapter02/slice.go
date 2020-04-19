package main

import "fmt"

var array = [6]int{1, 2, 3, 4, 5, 6}

func main() {
	//1.从数组或切片生成新的切片  [开始下标:结束下标但不包括该下标]
	a := array[0:3]
	for _, v := range a {
		fmt.Println(v) // 1  2 3
	}

	b := array[3:] //省却   如果都不写那么就是整个拿过来
	for _, v := range b {
		fmt.Println(v) // 4 5 6
	}
	//2.使用 make() 函数构造切片 ,make( []Type, size, cap )
	//Type 是指切片的元素类型，size 指的是为这个类型分配多少个元素，
	//cap 为预分配的元素数量，这个值设定后不影响 size，
	//只是能提前分配空间，降低多次分配空间造成的性能问题。
	c := make([]int, 2)     //数组长度2
	d := make([]int, 2, 10) //预分配的cap
	fmt.Println(c, d)
	//Go语言的内建函数 append() 可以为切片动态添加元素,这个有很多方法可以添加元素
	//注意append 是加在数组扩容后的下标开始
	d = append(d, 1, 2, 3, 4)
	fmt.Println(d)      //输出0 0 1 2 3 4 初始的 0  0 不会被覆盖哦
	fmt.Println(len(d)) //扩容因子 *2
}
