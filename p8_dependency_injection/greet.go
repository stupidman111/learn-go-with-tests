package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func GreetA(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func GreetB(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

//os.Stdout实际是一个os.*File，他实现了io.Writer接口的Write方法
//bytes.Buffer结构体也实现了io.Writer接口的Write方法
//但os.Stuout跟bytes.Buffer没有任何关系
func main() {
	//GreetA(os.Stdout, "zy")
	GreetB(os.Stdout, "zyy")
}