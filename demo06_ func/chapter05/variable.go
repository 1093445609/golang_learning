package main

import "fmt"

func main() {
	//传入多个参数
	myfunc(520, 1314, 784)

	//任意类型参数
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)
}

//传入的参数个数不确定,语法糖...会接收多个参数
func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

//任意类型的可变参数
//传任意类型，可以指定类型为 interface{}
//用 interface{} 传递任意类型数据是Go语言的惯例用法，使用 interface{} 仍然是类型安全的

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		//判断传入参数的type
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}
