package main

import (
	"fmt"
	"time"
)

//Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。

//goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

func main() {
	go say("world")
	say("hello")
}
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
