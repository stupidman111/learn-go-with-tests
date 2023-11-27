# map总结

## 在golang中map的键必须是可比较类型的
> 在golang中map的键必须是可比较类型的， 如果不能判断两个键是否相等，就无法确保得到正确的值。
* 如果自定义类型是基本类型（如int、string、float64等），则Go会使用默认的相等运算符来判断两个键是否相等。
* 如果自定义类型是结构体类型，Go会比较结构体的每个字段，判断它们是否相等。只有当结构体的所有字段都相等时，两个结构体才被认为是相等的。
* 如果自定义类型是切片、数组、函数等类型，由于这些类型是不可比较的，你将无法将它们用作map的键。

在使用自定义类型作为map的键时，确保你的自定义类型是可比较的。可比较的类型是可以使用 == 运算符进行比较的类型。

如果你需要自定义类型的相等性判断行为，可以为该类型实现 Equal 方法的重载。这样你就能够自定义两个实例是否相等的判断逻辑。例如：
```go
type MyKey struct {
    // fields
}

func (mk MyKey) Equal(other MyKey) bool {
    // 自定义相等性判断逻辑
    // 返回 true 表示相等，false 表示不相等
}

// 在使用 MyKey 作为 map 的键时，Go 会使用 Equal 方法来判断两个键是否相等。
```

## map是引用类型
map不使用指针传递你就可以修改它们，因为map是引用类型

这意味着它拥有对底层数据结构的引用，就像指针一样。它底层的数据结构是 hash table 或 hash map

Map 作为引用类型是非常好的，因为无论 map 有多大，都只会有一个副本。

引用类型引入了 maps 可以是 nil 值。如果你尝试使用一个 nil 的 map，你会得到一个 nil 指针异常，这将导致程序终止运行。

```go
//未初始化的map:
var myMap map[string]int
myMap["key"] = 1 // 尝试对未初始化的map进行操作，会导致panic

//使用make函数创建map:
myMap := make(map[string]int)
myMap = nil
myMap["key"] = 1 // 尝试对nil map进行操作，会导致panic

//通过nil指针间接创建map:
var pointerToMap *map[string]int
*pointerToMap = make(map[string]int)
(*pointerToMap)["key"] = 1 // 尝试对通过nil指针间接创建的map进行操作，会导致panic

````

## Error
* Error类型可以使用`.Error` 方法转换为字符串
* 可以通过`errors.New("msg")`来创建一个自定义error
