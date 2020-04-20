package main

import (
	"fmt"
	"time"
)

//Go语言计算函数执行时间

func test() {
	start := time.Now() // 获取当前时间
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum++
	}
	//从开始时间计算 花了多少ms
	//elapsed := time.Since(start)
	elapsed := time.Now().Sub(start) //方式2
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func main() {
	test()
}
