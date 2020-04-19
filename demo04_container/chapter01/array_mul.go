package main

/*多维数组*/
import (
	"fmt"
)

func main() {
	var array = [][]int{{1, 2}, {3, 4}, {5, 6}} //三个数组 每个数组有两个数组 也就是和下面一样
	var array2 = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(len(array))
	fmt.Println(len(array2))
	//给多维数组赋值
	var array3 [3][2]int

	array3[0][0] = 1
	array3[0][1] = 2

	array3[1][0] = 3
	array3[1][1] = 4

	array3[2][0] = 1
	array3[2][1] = 2
	//多维数组循环
	//实际上只要再次循环value即可
	for key, value := range array3 {
		fmt.Println(key, ":")
		for key2, v := range value {
			fmt.Print(key2, ":", v)
			fmt.Println("")
		}
	}

}
