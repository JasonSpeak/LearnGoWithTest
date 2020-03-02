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