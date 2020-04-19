package main //声明main包
//其中 package 是声明包名的关键字，name 为包的名字。
//
//Go语言的包与文件夹是一一对应的，它具有以下几点特性：
//一个目录下的同级文件属于同一个包。
//包名可以与其目录名不同。
//main 包是Go语言程序的入口包，一个Go语言程序必须有且仅有一个 main 包。如果一个程序没有 main 包，那么编译时将会出错，无法生成可执行文件。

import (
	"fmt" //导入fmt,在包声明之后，是import，用于导入程序中所依赖的包 ,相当于java的导包.
)

func main() { //main函数为入口函数,只能声明在main包下而且有且仅有一个.
	fmt.Print("Hello,Go!") //输出hello
}
