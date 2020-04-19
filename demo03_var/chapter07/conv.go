package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 123
	var a2 int32 = 456
	var a3 int64 = 789
	var a4 float64 = 5201314.2
	//var b  = "hello go"
	var c string
	var d float64
	//var e  float32
	c = strconv.Itoa(a) //int转string
	fmt.Println(c)
	a, _ = strconv.Atoi(c) //string 转int
	fmt.Println(a)
	c = strconv.FormatInt(a3, 10) //int64转10进制的数字 然后转成string
	fmt.Println(c)
	c = strconv.FormatInt(int64(a2), 10) // int32转为string          先转换64 再操作
	fmt.Println(c)
	c = strconv.FormatFloat(a4, 'f', 2, 64) // float64转为string  参数 1 f64 2 byte 'f' 3 保留小数 4.位数
	d, _ = strconv.ParseFloat(c, 64)        //string 转为float64 但是是科学计数法..
	fmt.Println(d)
	fmt.Println(c)
	c = strconv.FormatBool(true) //bool转string
	fmt.Println(c)

}
