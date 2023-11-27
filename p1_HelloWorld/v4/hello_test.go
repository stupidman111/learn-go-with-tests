package main

import "testing"

//这里使用到了子测试，并且重构重复的代码
/* 
	t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）。
	通过这样做，当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部。
	这将帮助其他开发人员更容易地跟踪问题。

*/
func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Zyy")
		want := "Hello, Zyy"
		
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world when an empty string is supplid", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}