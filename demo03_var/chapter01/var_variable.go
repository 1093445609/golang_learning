package main

/*变量的声明*/
import "fmt"

//1.声明全局变量
var var1 int = 123        //变量声明一般格式,var name type
var var2 = 456            //其中type可以省略,会自动判断类型
var var3, var4 int = 3, 4 //多个变量
//2.批量声明变量 var,具体变量类型就不介绍了,数组是在类型前+[] ,struct结构体,tunc函数
var (
	a int
	b string
	c []float32
	d func() bool
	e struct {
		x int
	}
)

func main() { //主要格式"{"不能单独换行 这是乌龟的屁股: 规定
	fmt.Println("var1:", var1) //123
	fmt.Println("var2:", var2) //456
	fmt.Println("a:", a)       //int a没有赋值 默认为0  和Java差不多
	fmt.Println("b:", b)       //string b 默认为""空字符串 所以没有打印结果
	fmt.Println(":", c)        //[] float32 数组未分配  结果为[]
	fmt.Println("d:", d)       //打印为<nil> 声明了一个变量但没有进行赋值，变量就会有一个类型的默认零值 nil意为"无"

	//3.简短声明局部变量,切记这种方式是作用在局部哦,名字 := 表达式

	f := 789       //局部变量必须使用,否则报错
	fmt.Println(f) //注释会报错,个人理解:这是乌龟的屁股,规定.你的局部b变量如果没有使用那么毫无意义

	//4.简短声明多个
	g, h := 110, "i am string" //INT STRING

	fmt.Println("g:", g)
	fmt.Println(h)
}
