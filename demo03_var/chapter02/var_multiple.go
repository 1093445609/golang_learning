package main

/*多重赋值*/

import "fmt"

//使用 Go 的“多重赋值”特性，可以轻松完成变量交换的任务
func main() {
	a := 100
	b := 520

	a, b = b, a
	fmt.Println("a:", a) //520
	fmt.Println("b:", b) //100

}
