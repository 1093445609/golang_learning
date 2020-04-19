package main

import (
	"fmt"
)

/*迭代集合*/
func main() {
	a := []int{10, 11, 12, 13}
	//方式1  k,v  下标+value
	for index, v := range a {
		fmt.Println(index, ":", v)
	}

	//方式2  不需要下标
	for _, v := range a {
		fmt.Println(":", v)
	}

	//方式3 普通for,index为下标
	for index := 0; index < len(a); index++ {
		fmt.Println(a[index])
	}

}
