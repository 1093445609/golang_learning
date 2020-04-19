package main

import "fmt"

/*使用 goto 集中处理错误*/
//伪代码
func main() {
	err := firstCheckError()
	if err != nil {
		//判断发生了错误 执行goto
		goto onExit
	}
	err = secondCheckError()
	if err != nil {
		//继续检查第二个
		//发生错误执行goto
		goto onExit
	}
	//没有错误 执行内容
	fmt.Println("done")
	return
	//发生错误时执行下面
onExit:
	fmt.Println(err)
	exitProcess()
}
