package main

import "fmt"

var f = 1

func main() {
	if f < 0 {
		fmt.Print("小于0")
	} else {
		fmt.Println("大于等于0")
	}
}
