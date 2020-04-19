package main

import "fmt"

func main() {
	var a int16 = 110
	var b int32 = 520
	var c int64 = 250
	var result int32
	var bolean bool = false
	//result=a  错误,未作显式转换 go没有隐式类型转换
	result = int32(a)   //int16转int32,精度不丢失
	fmt.Println(result) //110
	result = b
	fmt.Println(result) //520
	//result=bolean   类型不同 报错 cannot use bolean (type bool) as type int32 in assignment
	//fmt.Println(result)
	fmt.Println(bolean)
	result = int32(c)   //int64转int32 可能发生丢失精度问题
	fmt.Println(result) //250

	//result=int32(a+b)  不同类型运算未作转换也会报错 没有隐式转换   invalid operation
	result = int32(int32(a) + b)
	fmt.Println(result) //520+110=630
}
