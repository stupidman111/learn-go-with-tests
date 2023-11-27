package main

import (
	"testing"
)

/*
	测试文件必须名为xxx_test.go
	测试函数名必须以Test开头
	测试函数参数必须是*testing.T
*/
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}