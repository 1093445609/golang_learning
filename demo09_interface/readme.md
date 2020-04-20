~~~~
Go语言不是一种 “传统” 的面向对象编程语言：它里面没有类和继承的概念。

但是Go语言里有非常灵活的接口概念，通过它可以实现很多面向对象的特性。很多面向对象的语言都有相似的接口概念，
但Go语言中接口类型的独特之处在于它是满足隐式实现的。也就是说，我们没有必要对于给定的具体类型定义所有满足的接口类型；
简单地拥有一些必需的方法就足够了。
~~~~
~~~~
接口声明的格式
每个接口类型由数个方法组成。
`type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}`

~~~~
~~~~
接口被实现的条件一：接口的方法与实现接口的类型方法格式一致

只要实现接口类型中的方法的名称、参数列表、返回参数列表中的任意一项与接口要实现的方法不一致，
那么接口的这个方法就不会被实现。

package main
import (
    "fmt"
)
// 定义一个数据写入器
type DataWriter interface {
    WriteData(data interface{}) error
}
// 定义文件结构，用于实现DataWriter
type file struct {
}
// 实现DataWriter接口的WriteData方法
func (d *file) WriteData(data interface{}) error {
    // 模拟写入数据
    fmt.Println("WriteData:", data)
    return nil
}
func main() {
    // 实例化file
    f := new(file)
    // 声明一个DataWriter的接口
    var writer DataWriter
    // 将接口赋值f，也就是*file类型
    writer = f
    // 使用DataWriter接口进行数据写入
    writer.WriteData("data")
}
~~~~

~~~~
接口被实现的条件二：接口中所有方法均被实现
当一个接口中有多个方法时，只有这些方法都被实现了，接口才能被正确编译并使用。
~~~~