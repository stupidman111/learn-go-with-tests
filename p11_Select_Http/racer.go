package main

import (
	"fmt"
	"net/http"
	"time"
)

func RacerA(a, b string) (winner string) {

	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}

	return b
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func RacerB(a, b string) (winner string, err error) {
	//select用于同时监听多个管道，哪个管道先解除阻塞，select就会执行该case中的代码
	select {
	case <- ping(a):
		return a, nil
	case <- ping(b):
		return b, nil
	}
}

var tenSecondTimeout = 10 * time.Second
//使用RacerC则默认超时时间是10秒
func RacerC(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

//当想自定义超时时间时就使用`可配置的Racer函数`
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <- ping(a):
		return a, nil
	case <- ping(b):
		return b, nil
	case <- time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}