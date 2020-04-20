package main

import (
	"fmt"
)

//Go语言类型分支（switch判断空接口中变量的类型）
func main() {

}

//伪代码
//传入接口中变量
func (areaIntf interface{}) {
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}
