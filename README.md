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



#### HTTP中的connect

假设浏览器与服务器之间的 HTTPS 通信都是加密的，浏览器通过代理服务器发起 HTTPS 请求时，由于请求的站点地址和端口号都是加密保存在 HTTPS 请求报文头中的，代理服务器如何知道往哪里发送请求呢？为了解决这个问题，浏览器通过 HTTP 明文形式向代理服务器发送一个 CONNECT 请求告诉代理服务器目标地址和端口，代理服务器接收到这个请求后，会在对应端口与目标站点建立一个 TCP 连接，连接建立成功后返回 HTTP 200 状态码告诉浏览器与该站点的加密通道已经完成。接下来代理服务器仅需透传浏览器和服务器之间的加密数据包即可，代理服务器无需解析 HTTPS 报文。

举一个简单例子：

1. 浏览器向代理服务器发送 CONNECT 请求。

```bash
CONNECT geektutu.com:443 HTTP/1.0
```

1. 代理服务器返回 HTTP 200 状态码表示连接已经建立。

```bash
HTTP/1.0 200 Connection Established
```

1. 之后浏览器和服务器开始 HTTPS 握手并交换加密数据，代理服务器只负责传输彼此的数据包，并不能读取具体数据内容（代理服务器也可以选择安装可信根证书解密 HTTPS 报文）。

事实上，这个过程其实是通过代理服务器将 HTTP 协议转换为 HTTPS 协议的过程。对 RPC 服务端来，需要做的是将 HTTP 协议转换为 RPC 协议，对客户端来说，需要新增通过 HTTP CONNECT 请求创建连接的逻辑。



#### 负载均衡

假设有多个服务实例，每个实例提供相同的功能，为了提高整个系统的吞吐量，每个实例部署在不同的机器上。客户端可以选择任意一个实例进行调用，获取想要的结果。那如何选择呢？取决了负载均衡的策略。对于 RPC 框架来说，我们可以很容易地想到这么几种策略：

- **随机选择策略** - 从服务列表中随机选择一个。
- **轮询算法(Round Robin)** - 依次调度不同的服务器，每次调度执行 i = (i + 1) mode n。
- **加权轮询(Weight Round Robin)** - 在轮询算法的基础上，为每个服务实例设置一个权重，高性能的机器赋予更高的权重，也可以根据服务实例的当前的负载情况做动态的调整，例如考虑最近5分钟部署服务器的 CPU、内存消耗情况。
- **哈希/一致性哈希策略** - 依据请求的某些特征，计算一个 hash 值，根据 hash 值将请求发送到对应的机器。一致性 hash 还可以解决服务实例动态添加情况下，调度抖动的问题。一致性哈希的一个典型应用场景是分布式缓存服务。







## 错误及解决



## 没搞懂的问题



