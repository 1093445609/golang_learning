package main

/*copy()：切片复制（切片拷贝）*/
import "fmt"

func main() {
	//1.copy( toSlice, fromSlice []T) int

	a := []int{1, 2, 3}
	b := []int{4, 5, 6, 7, 8, 9}
	copy(a, b)         //将b复制到 a  截取b 前三的值copy到 a
	fmt.Println(a)     //因为a 长度只有3  输出 4 56
	a = []int{1, 2, 3} //重新给 a 赋值
	copy(b, a)
	fmt.Println(b) //123 789
}
