package main

import "fmt"

func main() {
	var array = [6]int{1, 3, 5, 7, 9}    //创建len=6的数组
	var array2 = []int{1, 3, 5, 7, 9}    //创建数组的长度为定义好的 5
	var array3 = [...]int{1, 3, 5, 7, 9} //...是接收右边定义的数量 和上面一样 5
	//普通for循环取值
	for a := 0; a < len(array); a++ {
		fmt.Println(a, ":", array[a])
	}
	fmt.Println("----------")
	//range循环取值
	for key, value := range array {
		fmt.Println(key, ":", value)
	}
	fmt.Println("----------")
	//range循环取值,不需要下标    可以使用"_"取代key
	for _, value := range array {
		fmt.Println(value)
	}
	fmt.Println(len(array2)) //5
	fmt.Println(len(array3)) //5
}
