package main

import "fmt"

//全局变量 不重名
var f2 = 1

func main() {
	if f2 < 0 {
		fmt.Print("小于0")
	} else if f2 == 0 {
		fmt.Print("等于0")
	} else {
		fmt.Println("大于0")
	}
}
