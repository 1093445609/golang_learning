package main

import "fmt"

//Go语言链表操作
type Node struct {
	data int
	next *Node
}

func Shownode(p *Node) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.next //移动指针
	}
}
func main() {
	var head = new(Node)
	head.data = 1
	var node1 = new(Node)
	node1.data = 2
	head.next = node1
	var node2 = new(Node)
	node2.data = 3
	node1.next = node2
	Shownode(head)
}

//头插 尾插法 省略

//双向链表省略
