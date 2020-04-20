package main

import "fmt"

/*匿名函数*/

//func(参数列表)(返回参数列表){
//    函数体
//}
func main() {
	//没有名字的函数,后面紧跟着"(100)" 表示调用该函数,参数为100
	func(data int) {
		fmt.Println("hello", data)
	}(100)

	// 将匿名函数体保存到f()中
	f := func(data int) {
		fmt.Println("hello", data)
	}
	// 使用f()调用
	f(100)

}
