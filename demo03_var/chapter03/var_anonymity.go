package main

/*匿名变量 _ */
import "fmt"

var a = []int{1, 2, 3, 8} //声明一个数组 a
func main() {
	for _, num := range a { //后面会练习range  你可以理解为java for each

		fmt.Println(num) //我们不需要使用该元素的序号，所以我们使用空白符"_"省略了 可以使用i来替代获取下标
	}
}
