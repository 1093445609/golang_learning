package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var mapList map[string]int = map[string]int{"key1": 1, "key2": 2}
	fmt.Println("key1:", strconv.Itoa(mapList["key1"])) //不需要转
	fmt.Println(mapList)
	fmt.Println(mapList["key2"])

	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960
	//遍历 k-v 没有下标
	for k, v := range scene {
		fmt.Println(k, v)
	}

	//遍历 k-v 没有key
	for _, v := range scene {
		fmt.Println(v)
	}

	//map元素的删除和清空

	delete(scene, "brazil")
	for k, v := range scene {
		fmt.Println(k, v)
	}

	//清空 map 中的所有元素
	//有意思的是，Go语言中并没有为 map 提供任何清空所有元素的函数、方法，清空 map 的唯一办法就是重新 make 一个新的 map，不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。

	//Go语言中的 map 在并发情况下，只读是线程安全的，同时读写是线程不安全的。

	//Go语言sync.Map（在并发环境中使用的map）

	var scene2 sync.Map
	// 将键值对保存到sync.Map
	scene2.Store("greece", 97)
	scene2.Store("london", 100)
	scene2.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene2.Load("london"))
	// 根据键删除对应的键值对
	scene2.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene2.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

}
