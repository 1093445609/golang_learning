package main //包
import "fmt" //导入的"fmt",必须使用,否则会报错哦.

/*
 假如你没有使用idea等ide来编写go
 那么,你需要先编译.go文件 为二进制文件 即生成.exe文件
 再进行执行.exe
 cmd命令步骤:
	1. go build 编译一个.go文件
	2.执行一个 .exe文件

或者:
	go run命令将编译和执行指令合二为一，会在编译之后立即执行Go语言程序，但是不会生成可执行文件。
cmd命令:
	go run Demo.go

*/

func main() {
	//fmt.Print("hello,goBuild,goRun!")   注释后报错 红线
	fmt.Print("hello,goBuild,goRun!") //使用了就不会报错
}
