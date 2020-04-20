package main

import "fmt"

func main() {
	a := gInt()
	b := gInt2(2)
	c, d := gInt3(3) //不需要返回的值可以使用 "_"省缺

	fmt.Print(a)
	fmt.Print(b)
	fmt.Println(c, d)
}

//无参函数 返回int
func gInt() int {
	return 1
}

//带参函数 注意 go没有重载
func gInt2(a int) int {
	return a
}

//返回多个参数
func gInt3(a int) (int, string) {
	return a, "hello,func!"
}

//返回多个参数,返回值命名
func gInt4(a int) (b int, c string) {
	b = a
	return b, "hello,func!"
}
