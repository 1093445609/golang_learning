package main

import "fmt"

//手机接口,只有一个方法call
type Phone interface {
	call()
}

//结构体(习惯叫它类)
type NokiaPhone struct {
}

//类诺基亚实现接口 重新call方法
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

//结构体类
type IPhone struct {
}

//苹果类实现接口 重写call
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}
