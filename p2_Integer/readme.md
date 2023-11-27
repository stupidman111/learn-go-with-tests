# 总结

## golang中示例的编写
* 示例同样位于`xxx_test.go`文件中
* 示例以ExampleXxx()为方法名
* 在代码注释中使用 `// Output: 预期结果值`来演示代码，如果实际输出预期不同则会报错
* 使用`go test -v`来运行
