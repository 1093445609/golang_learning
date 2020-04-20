package main

import "fmt"

//结构体
type a struct {
	id   int
	name string
	age  int
	//成员变量你可以嵌套一个结构体
}

func main() {
	//构造一个a实例
	user := a{id: 1, name: "蔡徐坤", age: 30}
	fmt.Print(user)
	//访问a 成员变量
	fmt.Println(user.id)

	var b a
	b.id = 2
	b.age = 35
	b.name = "柳岩"

	fmt.Println(b)

	//结构体指针

	//var c=	&user.id  //结构体的成员有地址
	var c = &user //结构体地址打印不出来? 暂时不清除为啥
	fmt.Println(c)
	d := *c
	fmt.Println(d)
}
