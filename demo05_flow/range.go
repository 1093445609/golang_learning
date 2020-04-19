package main

//for range 结构是Go语言特有的一种的迭代结构，在许多情况下都非常有用，
//for range 可以遍历数组、切片、字符串、map 及通道（channel）
import "fmt"

func main() {
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d  value:%d\n", key, value)
	}

	var str = "hello 你好"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}

	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}

}
