## Learn GO with [GitBook](https://studygolang.gitbook.io/learn-go-with-tests/)
### Hello World
#### 学习编写测试
- 编写一个测试
- 让编译通过
- 运行测试，查看失败原因并检查错误消息是很有意义的
- 编写足够的代码以使测试通过
- 重构
- 通过编写自动测试代码来保证测试的覆盖率和准确性，不需要通过运行软件来查看代码的正确性
- 学习TDD（Test drive develop）的过程，通过将问题分解成更小的可测试的组件，编写软件将会更加轻松
- 在测试中完善开发需求
### 学习编写注释
- 在方法中适当编写注释可以方便得生成文档，方便后续维护
- 使用代码示例，代码示例将作为测试用例的一部分

### 数组与迭代
- 数组前指定长度的为数组，不指定长度的为切片

```golang
    for _, num := range array {
		sum += num
	}
```
- 使用`range`迭代
    range迭代每次会返回数组的索引与值，这里使用空白符`_`替代返回的索引

### 结构体 方法 接口
#### 定义结构体
```golang
type Rectangle struct {
	width  float64
	height float64
}
```

#### 为结构体添加方法
```golang 
//Area returns the area of a rectangle
func (rect Rectangle) Area() float64 {
	if rect.width <= 0.0 || rect.height <= 0.0 {
		return 0.0
	}
	return rect.height * rect.width
}
```
- 即可使用`rectangle.Area()`来调用Area()方法

#### 接口
- golang的接口继承是隐式的，只要传入的类型与接口的定义相匹配，则编译正确

例如：
```golang 
//Shape defines the shape interface
type Shape interface {
	Area() float64
}
```
- 定义了一个`IShape`接口，接口中定义了一个Area函数，则上面的`Rectangle`结构体与该接口匹配，且不需要在`Rectangle`中显示声明
- 通过声明一个接口，辅助函数能从具体类型解耦而只关心方法本身需要做的工作。


#### 表格驱动测试
```golang
func TestArea(t *testing.T) {

    areaTests := []struct {
        shape Shape
        want  float64
    }{
        {Rectangle{12, 6}, 72.0},
        {Circle{10}, 314.1592653589793},
    }

    for _, tt := range areaTests {
        got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got %.2f want %.2f", got, tt.want)
        }
    }

}
```
- 如上例代码，创建一个匿名结构体，含有两个域：`shape`和`want`,后续就可以方便的向其中添加测试用例，并且不会破坏之前的测试

### 指针与错误
#### 结构体数据的私有性
- 在 Go 中，如果一个符号（例如变量、类型、函数等）是以小写符号开头，那么它在 定义它的包之外 就是私有的。

#### 指针
- 当你传值给函数或方法时，Go 会复制这些值。因此，如果你写的函数需要更改状态，你就需要用指针指向你想要更改的值
- Go 取值的副本在大多数时候是有效的，但是有时候你不希望你的系统只使用副本，在这种情况下你需要传递一个引用。例如，非常庞大的数据或者你只想有一个实例（比如数据库连接池）

#### nil
- 指针可以是 nil
- 当函数返回一个的指针，你需要确保检查过它是否为 nil，否则你可能会抛出一个执行异常，编译器在这里不能帮到你
- nil 非常适合描述一个可能丢失的值

#### 错误
- 错误是在调用函数或方法时表示失败的
- 通过测试我们得出结论，在错误中检查字符串会导致测试不稳定。因此，我们用一个有意义的值重构了，这样就更容易测试代码，同时对于我们 API 的用户来说也更简单。
- 错误处理的故事远远还没有结束，你可以做更复杂的事情，这里只是抛砖引玉。后面的部分将介绍更多的策略。
- 不要只是检查错误，要优雅地处理错误

#### 从现有的类型中创建新的类型
- 用于为值添加更多的领域内特定的含义
- 可以让你实现接口

### Maps
#### 用法
- map用法为`map[type1]type2`
- 例如：`map[string]string`，定义了一个key为string类型，value为string类型的map

#### 特性
- map查找可以返回两个值,若要查询的key-vlaue存在，则第一个值为查到的value，第二个值为`bool`值`true`;否则第一个值为空，第二个值为`false`
- 例如:
```golang
dictionary map[string]string
value,result := dictionary[key]
```
- map是引用类型，所以不需要指针传递就可以修改原数据
- 由于map是引用类型，所以map可以是`nil`值，如果使用一个nil的map，就会得到nil指针异常，导致程序终止
- 所以永远不要初始化一个空的map变量,例如`var m map[string]string`
- 可以通过下列方法初始化空map
```golang
dictionary = map[string]string{}

// OR

dictionary = make(map[string]string)
```
- 这两种方法都可以创建一个空的 hash map 并指向 dictionary。这确保永远不会获得 nil 指针异常

### Mocking
- 反引号语法是创建string的另一种形式，例如
```golang
want := `3 
2
1
GO!`
```

### 并发
- goroutines 是 Go 的基本并发单元，它让我们可以同时运行多项任务
- anonymous functions（匿名函数），我们用它来启动每个并发进程
- channels，用来组织和控制不同进程之间的交流，使我们能够避免 race condition（竞争条件）(锁) 的问题。
- the race detector（竞争探测器） 帮助我们调试并发代码的问题。

#### channels
- channel 内部就是一个带锁的队列
- https://www.cnblogs.com/oxspirt/p/7124291.html

#### Select 
- 可帮助你同时在多个 channel 上等待。
- 有时你想在你的某个「案例」中使用 time.After 来防止你的系统被永久阻塞

#### httptest标准库
- 一种方便地创建测试服务器的方法，这样你就可以进行可靠且可控的测试。
- 使用和 net/http 相同的接口作为「真实的」服务器会和真实环境保持一致，并且只需更少的学习。


### json 路由 嵌入
#### 嵌入
- Go 没有提供典型的，类型驱动的子类化概念，但它具有通过在结构或接口中嵌入类型来“借用”一部分实现的能力。
- 例如：
```golang
type PlayerServer struct {
	store PlayerStore
	http.Handler
}
```
- 通过在`PlayerServer`结构体中**嵌入**`http.handler`类型，这样，`PlayerServer`就可以拥有`http.handler`所有的公开方法以及属性
- 你必须小心使用嵌入类型，因为你将公开所有嵌入类型的公共方法和字段。在我们的例子中它是可以的，因为我们只是嵌入了 `http.Handler` 这个 接口。
- 如果我们懒一点，嵌入了 `http.ServeMux`（混合类型），它仍然可以工作 但 `PlayerServer` 的用户就可以给我们的服务器添加新路由了，因为 `Handle(path, handler)`会公开。
- 路由。标准库为你提供了易于使用的类型来处理路由。它完全支持 http.Handler 接口，因为你可以将路由分配给 Handler，而路由本身也是 Handler。它没有你可能期望的某些特性，例如路径变量（例如 /users/{id}）。你可以自己轻易地解析这些信息，但如果它成了负担，你可能会考虑查看其它路由库。大多数流行的库都坚持标准库的实现 http.Handler 的理念。
- 类型嵌入。我们对这项技术略有提及，但你可以从 Effective Go 了解更多信息。如果你应该从中得到一个收获，那就是它极其有用，但是 总是考虑你的公开 API，只有适合被公开的才公开。
- JSON 反序列化和序列化。标准库使得序列化和反序列化数据变得非常简单。它也是开放的配置，你可以根据需要自定义这些数据转换的工作方式。