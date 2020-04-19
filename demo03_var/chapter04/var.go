package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = 100
	var st1 = "字符串1"
	//var st2 = "字符串"+a 错误不能和int拼接,转换为同一类型再 +
	var st2 = "字符串" + strconv.Itoa(a) //字符串100
	fmt.Println(st1)
	fmt.Println(st2)
}
