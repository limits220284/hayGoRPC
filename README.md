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



#### 在 Go 语言中，`...`（三个点）有两个主要的用法，它们分别是可变参数和切片的展开。

- **可变参数（Variadic Parameters）：** 在函数声明中，`...` 可以用于指定可变数量的参数。这样的参数被视为一个切片。例如：

```go
goCopy codefunc exampleFunction(args ...int) {
    // args 是一个 int 切片
    // 可以通过 args[index] 访问每个参数
}
```

调用这个函数时，可以传递任意数量的整数：

```go
goCopy codeexampleFunction(1, 2, 3)
exampleFunction(4, 5, 6, 7, 8)
```

- **切片的展开（Slice Unpacking）：** 在函数调用时，`...` 可以用于展开切片，将切片的元素作为单独的参数传递给函数。例如：

```go
goCopy codenums := []int{1, 2, 3}
result := sum(nums...)
```

这里的 `sum` 函数可以接受多个整数作为参数，通过 `nums...` 将切片的元素展开传递给 `sum` 函数。

#### cap

在 Go 语言中，`cap` 是一个内建函数，用于获取切片（slice）、数组（array）或通道（channel）的容量。容量是数据结构可以容纳的元素个数，而长度（length）是当前实际包含的元素个数。



## 错误及解决



## 没搞懂的问题



