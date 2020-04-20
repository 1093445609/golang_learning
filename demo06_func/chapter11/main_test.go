package main

//单元测试
//文件名 以_test.go结尾
import "testing"

//cmd命令 go test -v
func TestGetArea(t *testing.T) {
	area := GetArea(40, 50)

	if area != 2000 {
		t.Error("测试失败")
	}
}

//压力测试
//go test -bench="BenchmarkGetArea$" --run=none
func BenchmarkGetArea(t *testing.B) {
	for i := 0; i < t.N; i++ {
		GetArea(40, 50)
	}
}

//覆盖率测试
//go test -cover
//最好是100%
//覆盖率测试能知道测试程序总共覆盖了多少业务代码（也就是 demo_test.go 中测试了多少 demo.go 中的代码），可以的话最好是覆盖100%。
