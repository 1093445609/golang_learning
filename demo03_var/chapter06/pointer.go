package main

/*go指针学习*/
import "fmt"

func main() {
	a := 110
	b := &a
	fmt.Println(b)  //b的值为a的内存地址 0xc00000a0d8
	fmt.Println(*b) //*b 表示取出内存地址为b里面储存的值 打印 110
	*b = 120
	fmt.Println(a)    //操作修改 *b会修改掉a的值         打印120
	var c func() bool //对象没有初始化,有默认值的不会发生空指针哦
	d := &c
	fmt.Println("打印c地址", d)
	fmt.Println("打印c值", *d)
	if nil == c {
		fmt.Println("发生空指针")
	}
}
