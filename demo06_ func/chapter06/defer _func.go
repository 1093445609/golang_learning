package main

import (
	"fmt"
	"sync"
)

var (
	// 一个演示用的映射
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)

/*多个延迟执行语句的处理顺序*/
func main() {
	fmt.Println("defer begin")
	// 将defer放入延迟调用栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	// 最后一个放入, 位于栈顶, 最先调用
	defer fmt.Println(3)
	fmt.Println("defer end")
	/*
		defer 先进后出
		输出
		defer begin
		defer end
		3
		2
		1
	*/

}

//使用延迟执行语句在函数退出时释放资源
//1) 使用延迟并发解锁

func readValue(key string) int {
	//上锁
	valueByKeyGuard.Lock()

	// defer后面的语句不会马上调用, 而是延迟到函数结束时调用
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}
