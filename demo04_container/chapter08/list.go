package main

import (
	"container/list"
	"fmt"
)

func main() {
	//初始化列表
	//方式1
	lists := list.New()
	//方式2
	var list2 list.List

	//插入元素
	lists.PushBack("a")
	lists.PushBack(1)

	//后插
	list2.PushBack(1)
	list2.PushBack(2)
	list2.PushBack(3)
	//前插
	list2.PushFront(10)

	// 尾部添加后保存元素句柄	可以更好操作
	element := list2.PushBack("fist")

	list2.InsertAfter("蔡徐坤", element)

	list2.InsertBefore("打篮球", element)
	list2.Remove(element)

	//遍历
	//在Go语言中，布尔类型的零值（初始值）为 false，数值类型的零值为 0，字符串类型的零值为空字符串""，
	//而指针、切片、映射、通道、函数和接口的零值则是 nil。
	for i := list2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
