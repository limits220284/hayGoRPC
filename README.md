# hayGoRPC

- 记录写该项目遇到的问题



## 知识点

#### 对 `net/rpc` 而言，一个函数需要能够被远程调用，需要满足如下五个条件：

- the method’s type is exported.
  - **方法的类型是可导出的**
- the method is exported.
  - **方法本身是可导出的**
- the method has two arguments, both exported (or builtin) types.
  - **方法有两个参数，都是可导出的(或内建的)类型**
- the method’s second argument is a pointer.
  - **方法的第二个参数是指针类型**
- the method has return type error.
  - **方法的返回值是error**

更直观一些：

```go
func (t *T) MethodName(argType T1, replyType *T2) error
```

通过反射，我们能够非常容易地获取某个结构体的所有方法，并且能够通过方法，获取到该方法所有的参数类型与返回值。例如：

```go
func main() {
	var wg sync.WaitGroup
	typ := reflect.TypeOf(&wg)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		argv := make([]string, 0, method.Type.NumIn())
		returns := make([]string, 0, method.Type.NumOut())
		// j 从 1 开始，第 0 个入参是 wg 自己。
		for j := 1; j < method.Type.NumIn(); j++ {
			argv = append(argv, method.Type.In(j).Name())
		}
		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}
		log.Printf("func (w *%s) %s(%s) %s",
			typ.Elem().Name(),
			method.Name,
			strings.Join(argv, ","),
			strings.Join(returns, ","))
    }
}
```

运行的结果是：

```go
func (w *WaitGroup) Add(int)
func (w *WaitGroup) Done()
func (w *WaitGroup) Wait()
```













#### 在 Go 语言中，`...`（三个点）有两个主要的用法，它们分别是可变参数和切片的展开。

- **可变参数（Variadic Parameters）：** 在函数声明中，`...` 可以用于指定可变数量的参数。这样的参数被视为一个切片。例如：

```go
func exampleFunction(args ...int) {
    // args 是一个 int 切片
    // 可以通过 args[index] 访问每个参数
}
```

调用这个函数时，可以传递任意数量的整数：

```go
exampleFunction(1, 2, 3)
exampleFunction(4, 5, 6, 7, 8)
```

- **切片的展开（Slice Unpacking）：** 在函数调用时，`...` 可以用于展开切片，将切片的元素作为单独的参数传递给函数。例如：

```go
nums := []int{1, 2, 3}
result := sum(nums...)
```

这里的 `sum` 函数可以接受多个整数作为参数，通过 `nums...` 将切片的元素展开传递给 `sum` 函数。

#### cap

在 Go 语言中，`cap` 是一个内建函数，用于获取切片（slice）、数组（array）或通道（channel）的容量。容量是数据结构可以容纳的元素个数，而长度（length）是当前实际包含的元素个数。



#### var _ io.Closer = (*Client)(nil)

这段代码是在 Go 语言中进行接口断言（interface assertion）的一种常见用法。让我来解释每个部分的含义：

1. **`io.Closer`：** 这是一个接口类型，它包含了 `Close` 方法，用于释放资源或执行清理工作。
2. **`var _ io.Closer = (*Client)(nil)`：**
   - `var` 关键字用于声明一个变量。
   - `_`（下划线）是一个特殊的标识符，用于占位，表示不关心这个变量的值。
   - `io.Closer` 表明声明的变量的类型是 `io.Closer` 接口。
   - `= (*Client)(nil)` 表示将一个 `nil` 值的 `*Client` 类型的指针赋给这个变量。

这段代码的目的是确保 `*Client` 类型实现了 `io.Closer` 接口。如果 `*Client` 类型正确实现了 `Close` 方法，那么它就被认为是 `io.Closer` 接口的实现。这种用法通常在编译时进行接口实现的检查，以确保代码的正确性。



## 错误及解决



## 没搞懂的问题



