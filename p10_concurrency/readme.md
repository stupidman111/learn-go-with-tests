# 总结

## goroutine

## 匿名函数

## channel

`channel <- value`：值传入channel

`value := <- channel`：从value中读取值

## 类型别名
使用type由一个类型构建另一个类型后
这两个类型是不相同的两个类型（底层也是）

## 使用go test -race来判断测试代码中是否存在竞争
当不确定代码是否并发安全时
可以使用`go test -race`来判断测试代码是否存在竞争