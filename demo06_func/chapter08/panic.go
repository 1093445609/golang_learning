package main

/*Go语言宕机（panic）——程序终止运行*/
func main() {

	//1.手动触发宕机
	panic("手动宕机!")

	//2.在运行依赖的必备资源缺失时主动触发宕机

	/*regexp 是Go语言的正则表达式包，正则表达式需要编译后才能使用，而且编译必须是成功的，表示正则表达式可用。

	编译正则表达式函数有两种，具体如下：
	1) func Compile(expr string) (*Regexp, error)
	编译正则表达式，发生错误时返回编译错误同时返回 Regexp 为 nil，该函数适用于在编译错误时获得编译错误进行处理，同时继续后续执行的环境。
	2) func MustCompile(str string) *Regexp
	当编译正则表达式发生错误时，使用 panic 触发宕机，该函数适用于直接使用正则表达式而无须处理正则表达式错误的情况。*/

	//3.在宕机时触发延迟执行语句
	//	defer fmt.Println("宕机后要做的事情1")
	//	defer fmt.Println("宕机后要做的事情2")
	//panic("发生宕机啦")

}
