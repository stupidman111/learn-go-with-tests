package main

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}


func TestCheckWebsites(t *testing.T) {
	websites := []string {
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, websites)

	//先检查长度
	want := len(websites)
	got := len(actualResults)

	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}

	//再检查内容
	expectedResult := map[string]bool{
		"http://google.com": true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds": false,
	}

	if !reflect.DeepEqual(actualResults, expectedResult) {
		t.Fatalf("Wanted %v, got %v", expectedResult, actualResults)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}