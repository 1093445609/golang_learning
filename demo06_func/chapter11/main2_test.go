package main

//压力测试
//go test -bench="BenchmarkGetArea2$" --run=none
import "testing"

func BenchmarkGetArea2(t *testing.B) {
	for i := 0; i < 999999999; i++ {
		GetArea(40, 50)
	}
}
