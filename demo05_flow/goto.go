package main

import "fmt"

/*使用 goto 退出多层循环*/
func main() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 跳转到标签
				goto breakHere
			}
		}
	}
	// 手动返回, 避免执行进入标签
	return
	// 标签     下面的代码必须满足条件 才执行
breakHere:
	fmt.Println("done")
}
