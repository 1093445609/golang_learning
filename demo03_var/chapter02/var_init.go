package main

/*变量赋值*/
import "fmt"

var (
	hp  int = 100
	hp2     = 101
)

func main() {
	hp3 := 520
	fmt.Println(hp)
	fmt.Println(hp2)
	fmt.Println(hp3)
	hp4 := int(hp2 - hp) //101-100 返回一个int类型
	fmt.Println(hp4)     //1
}
