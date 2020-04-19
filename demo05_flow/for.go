package main

import "fmt"

//Go语言for（循环结构）
func main() {

	//只支持 for 关键字，而不支持 while 和 do-while 结构
	//支持 continue 和 break 来控制循环，

JLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop //跳出外部循环
			}
			fmt.Println(i)
		}
	}

	//for 中的结束语句——每次循环结束时执行的语句
	//在结束每次循环前执行的语句，如果循环被 break、goto、return、panic 等语句强制退出，结束语句不会被执行。
	//continue 语句可以结束当前循环，开始下一次的循环迭代过程，仅限在 for 循环内使用
}
