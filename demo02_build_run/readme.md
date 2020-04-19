 假如你没有使用idea等ide来编写go
 那么,你需要先编译.go文件 为二进制文件 即生成.exe文件
 再进行执行.exe
 cmd命令步骤:
	1. go build 编译一个.go文件
	2.执行一个 .exe文件

或者:
	go run命令将编译和执行指令合二为一，会在编译之后立即执行Go语言程序，但是不会生成可执行文件。
cmd命令:
`	go run Demo.go`
    ~~~~~~~~~~~~~~~
	
go build
	
1.cmd进入项目路径

    `F:\>cd idea代码存放地址\godemo`

    `F:\Idea代码存放地址\godemo>cd demo01_hello_go`   

2.go build 需要编译的go文件

    `F:\Idea代码存放地址\godemo\demo01_hello_go>go build HelloWorld.go`
  
3.成功了会在.go同级文件夹下生成 .exe二进制文件

4.运行 二进制文件
    `F:\Idea代码存放地址\godemo\demo01_hello_go>  HelloWorld.exe`
结果:
    `Hello,Go!`

go run

1.这次我们cmd进入demo2

    `F:\Idea代码存放地址\godemo>cd demo02_build_run`
    
2. go run 会编译二进制并执行文件 但是不保存二进制文件!
    `F:\Idea代码存放地址\godemo\demo02_build_run> go run Demo.go`
执行结果:
    `hello,goBuild,goRun!`
~~~~