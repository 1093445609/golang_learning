package main

import (
	"bufio"
	"bytes"
	"fmt"
)

//Go语言数据I/O对象及操作

//ReadWriter 对象
/*ReadWriter 对象可以对数据 I/O 接口 io.ReadWriter 进行输入输出缓冲操作，ReadWriter 结构定义如下：
type ReadWriter struct {
	*Reader
	*Writer
}*/

//Reader 对象

//Reader 对象可以对数据 I/O 接口 io.Reader 进行输入缓冲操作

func main() {
	//Read() 方法的功能是读取数据，并存放到字节切片 p 中。
	//ReadByte() 方法的功能是读取并返回一个字节，如果没有字节可读，则返回错误信息。
	data := []byte("C语言中文网")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [128]byte
	n, err := r.Read(buf[:])
	//ReadBytes() 方法的功能是读取数据直到遇到第一个分隔符“delim”，并返回读取的字节序列（包括“delim”）。
	//line, err := r.ReadBytes(delim)
	fmt.Println(string(buf[:n]), n, err)

	//ReadLine() 是一个低级的用于读取一行数据的方法，大多数调用者应该使用 ReadBytes('\n') 或者 ReadString('\n')。

	//ReadRune() 方法的功能是读取一个 UTF-8 编码的字符，并返回其 Unicode 编码和字节数。

	// ReadSlice() 方法
	//ReadString() 方法
	//....
}
