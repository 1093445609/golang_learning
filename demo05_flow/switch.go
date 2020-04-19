package main

import "fmt"

func main() {
	var a = "hello"
	//不需要 break跳出循环
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}

	//case 可以多个值 可以使用表达式
	var a2 = "mum"
	switch a2 {
	case "mum", "daddy":
		fmt.Println("family")
	}

	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}
	//fallthrough 匹配到了还会继续往下执行
	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}
}
