package main

import (
	"errors"
	"fmt"
	"math"
)

//Go语言error接口：返回错误信息

//type error interface {
//    Error() string
//}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}
func main() {
	result, err := Sqrt(-13)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
