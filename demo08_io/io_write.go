package main

import (
	"bufio"
	"bytes"
	"fmt"
)

//Writer 对象
//Writer 对象可以对数据 I/O 接口 io.Writer 进行输出缓冲操作

func main() {
	wr := bytes.NewBuffer(nil)
	w := bufio.NewWriter(wr)
	p := []byte("C语言中文网")
	fmt.Println("写入前未使用的缓冲区为：", w.Available())
	w.Write(p)
	fmt.Printf("写入%q后，未使用的缓冲区为：%d\n", string(p), w.Available())
}

//操作 Writer 对象
//操作 Writer 对象的方法共有 7 个，分别是 Available()、Buffered()、Flush()、Write()、WriteByte()、WriteRune() 和 WriteString() 方法
