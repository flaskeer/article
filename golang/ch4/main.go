// ch4 project main.go
package main

/*
Go 有指针。然而却没有指针运算，因此它们更象是引用而不是你所知道的来自于C
的指针。指针非常有用。在Go 中调用函数的时候，得记得变量是值传递的。因此，为
了修改一个传递入函数的值的效率和可能性，有了指针。
通过类型作为前缀来定义一个指针’*’：var p *int。现在p 是一个指向整数值的指针。
所有新定义的变量都被赋值为其类型的零值，而指针也一样。一个新定义的或者没有
任何指向的指针，有值nil。在其他语言中，这经常被叫做空（NULL）指针，在Go 中
就是nil。让指针指向某些内容，可以使用取址操作符（&），

Go 有两个内存分配原语，new 和make。它们应用于不同的类型，做不同的工作，可能
有些迷惑人，但是规则很简单
*/
import (
	"fmt"
)

var p *int
var i int

/*
用new 分配内存
内建函数new 本质上说跟其他语言中的同名函数功能一样：new(T) 分配了零值填充
的T 类型的内存空间，并且返回其地址，一个*T 类型的值。用Go 的术语说，它返回
了一个指针，指向新分配的类型T 的零值。记住这点非常重要。
这意味着使用者可以用new 创建一个数据结构的实例并且可以直接工作。如
bytes.Buffer 的文档所述“Buffer 的零值是一个准备好了的空缓冲。” 类似的，
sync.Mutex 也没有明确的构造函数或Init 方法。取而代之， sync.Mutex 的零值
被定义为非锁定的互斥量。

type SyncedBuffer s t r u c t {
lock sync.Mutex
buffer bytes.Buffer
}
SyncedBuffer 的值在分配内存或定义之后立刻就可以使用。在这个片段中，p 和v 都
可以在没有任何更进一步处理的情况下工作。
p := new(SyncedBuffer)   Type *SyncedBuffer，已经可以使用
var v SyncedBuffer   Type SyncedBuffer，同上
*/

/*
用make 分配内存
回到内存分配。内建函数make(T, args) 与new(T) 有着不同的功能。它只能创建
slice，map 和channel，并且返回一个有初始值（非零）的T 类型，而不是*T。本质
来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。
例如，一个slice，是一个包含指向数据（内部array）的指针，长度和容量的三项描述
符；在这些项目被初始化之前，slice 为nil。对于slice，map 和channel，make 初始
化了内部的数据结构，填充适当的值。
例如，make([]int, 10, 100) 分配了100 个整数的数组，然后用长度10 和容量100
创建了slice 结构指向数组的前10 个元素。区别是，new([]int) 返回指向新分配的内
存的指针，而零值填充的slice 结构是指向nil 的slice 值。
*/

//var p *[]int = new([]int)  //分配slice内存结构 很少使用
//var v []int = make([]int,100) //指向一个新分配的有100个整数的数组
//var p *[]int = new([]int)  //不必要

//v := make([]int,100)  //更常见

/*
务必记得make 仅适用于map，slice 和channel，并且返回的不是指针。应当用new 获
得特定的指针。

• new(T) 返回*T 指向一个零值T
• make(T) 返回初始化后的T
当然make 仅适用于slice，map 和channel。
*/

//func NewFile(fd int, name string) *File {
//	if fd < 0 {
//		return nil
//	}
//	f := File{fd, name, nil, 0}
//	return &f
//}
const (
	Enone  = 1
	Einval = 2
)

type NameAge struct {
	name string
	age  int
}

//struct{
//	x,y int
//	A *[]int
//	F func()
//}

//func doSomething(n1 *NameAge,n2 int) { /* */}
func (n1 *NameAge) doSomething(n2 int) {
	fmt.Println(n2)
}

//注意首字母大写的字段可以被导出，也就是说，在其他包中可以进行读写。字段名以
//小写字母开头是当前包的私有的。包的函数定义是类似的
//struct{
//	T1
//	*T2
//	P.T3
//	x,y int
//}

/*
如果x 可获取地址，并且&x 的方法中包含了m，x.m() 是(&x).m() 更短
的写法。
根据上面所述，这意味着下面的情况不是错误：
var n NameAge  不是指针
n.doSomething(2)
这里Go 会查找NameAge 类型的变量n 的方法列表，没有找到就会再查找*NameAge
类型的方法列表，并且将其转化为(&n).doSomething(2)。
*/

//Mutex 数据类型有两个方法，Lock 和Unlock。
//type Mutex struct { /* Mutex 字段*/}
//func (m *Mutex) Lock() { /* Lock 实现*/ }
//func (m *Mutex) Unlock() { /* Unlock 实现*/ }
//现在用两种不同的风格创建了两个数据类型。
//type NewMutex Mutex;
//type PrintableMutex struct {Mutex }.
//现在NewMutux 等同于Mutex，但是它没有任何Mutex 的方法。换句话说，它的方法
//是空的。
//但是PrintableMutex 已经从Mutex 继承了方法集合。如同[10] 所说：
//PrintableMutex 的方法集合包含了Lock 和Unlock 方法，被绑定到其
//匿名字段Mutex

type foo struct {
	int
}
type bar foo

var b bar = bar{1}

//var f foo = b //报错  cannot use b (type bar) as type foo in assignment(不能使用b（类型bar）
//作为类型foo 赋值)
//var f foo = foo(b)

type Person struct{
	name string
	age int
}
var p1 Person
p2 := new(Person)

//x指向了t指向的内容，也就是实际上的参数指向的内容
func Set(t *T){
	x = t
}
//x指向了一个新的（堆上分配的）变量t。其包含了实际参数值的副本
//有了额外的变量存储相关值的副本
func Set(t T){
	x = &t
}

func main() {
	var n *NameAge
	n.doSomething(3)
	//转换
	mystring := "hello this is string"
	byteslice := []byte(mystring)
	runeslice := []rune(mystring)
	fmt.Println("%v---%v", byteslice, runeslice)
	/*
		从字节或者整形的slice 到string。
		b := []byte {'h','e','l','l','o'} // 复合声明
		s := s t r i n g (b)
		i := []rune {257,1024,65}
		r := s t r i n g (i)
	*/
	arr := [...]string{Enone: "no error", Einval: "invalid argument"}
	sl := []string{Enone: "no error", Einval: "invalid argument"}
	ma := map[int]string{Enone: "no error", Einval: "invalid argument"}
	p = &i
	*p = 8
	fmt.Printf("%v ++ %v", *p, i)
	fmt.Printf("%v==%T", arr, arr)
	fmt.Println()
	fmt.Printf("%v ==%T", sl, sl)
	fmt.Println()
	fmt.Printf("%v==%T", ma, ma)
	fmt.Println()
	a := new(NameAge)
	//结构中的项目被称为field。
	a.name = "Pete"
	a.age = 354
	fmt.Printf("%v\n ++ field : %d ---%s", a, a.age, a.name)
}
