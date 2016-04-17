// ch5 project main.go
package main

import (
	"fmt"
	"reflect"
)

type S struct {
	i int
}

func (p *S) Get() int {
	return p.i
}
func (p *S) Put(v int) {
	p.i = v
}

type R struct {
	i int
}

func (p *R) Get() int {
	return p.i
}
func (p *R) Put(v int) {
	p.i = v
}

//对于接口I，S 是合法的实现，因为它定义了I 所需的两个方法。注意，即便是没有明
//确定义S 实现了I，这也是正确的。
//Go 程序可以利用这个特点来实现接口的另一个含义，就是接口值:
//func f(p I) { ..0
//fmt.Println(p.Get()) ..1
//p.Put(1) ..2
//}
//..0 定义一个函数接受一个接口类型作为参数；
//..1 p 实现了接口I，必须有Get() 方法；
//..2 Put() 方法是类似的。
type I interface {
	Get() int
	Put(int)
}

//获取s 的地址，而不是S 的值的原因，是因为在s 的指针上定义了方法，参阅上面的
//代码5.1。这并不是必须的——可以定义让方法接受值——但是这样的话Put 方法就不
//会像期望的那样工作了。
func f(p I) {
	p.Put(1)

	fmt.Println(p.Get())

}

/*
函数f 现在可以接受类型为R 或S 的变量。假设需要在函数f 中知道实际的类型。在
Go 中可以使用type switch 得到。
类型判断。在switch 语句中使用(type)。保存类型到变量t；
..1 p 的实际类型是S 的指针；
..2 p 的实际类型是R 的指针；
..3 p 的实际类型是S；
..4 p 的实际类型是R；
..5 实现了I 的其他类型。
*/
//func f2(p I) {
//	switch t := p.(type) {
//	case *S:
//	case *R:
//	//case S:
//	//case R:
//	default:
//	}
//}

//在switch 之外使用(type) 是非法的。类型判断不是唯一的运行时得到类型的方法。
//为了在运行时得到类型，同样可以使用“comma, ok” 来判断一个接口类型是否实现了
//某个特定接口：
//i f t, ok := something.(I) ; ok {
//// 对于某些实现了接口I 的
//// t 是其所拥有的类型
//}
//t := something.(I)

/*
，Go 确实有纯粹动态的方面，如可将一个接口类型转
换到另一个。通常情况下，转换的检查是在运行时进行的。如果是非法转换——当在
已有接口值中存储的类型值不匹配将要转换到的接口——程序会抛出运行时错误。
在Go 中的接口有着与许多其他编程语言类似的思路：C++ 中的纯抽象虚基类，Haskell
中的typeclasses 或者Python 中的duck typing。然而没有其他任何一个语言联合了接
口值、静态类型检查、运行时动态转换，以及无须明确定义类型适配一个接口。
*/
/*
由于每个类型都能匹配到空接口：interface{}。我们可以创建一个接受空接口作为
参数的普通函数：
在这个函数中的return something.(I).Get() 是有一点窍门的。值something 具有
类型interface{}，这意味着方法没有任何约束：它能包含任何类型。.(I) 是类型
断言，用于转换something 到I 类型的接口。如果有这个类型，则可以调用Get() 函
数。因此，如果创建一个*S 类型的新变量，也可以调用g()，因为*S 同样实现了空
接口。
*/
func g(sth interface{}) int {
	return sth.(I).Get()
}

//方法
type Foo int

func (self Foo) Emit() {
	fmt.Printf("%v", self)
}

type Emitter interface {
	Emit()
}

/*
接口定义为一个方法的集合。方法包含实际的代码。换句话说，一个接口就是定义，
而方法就是实现。因此，接收者不能定义为接口类型，这样做的话会引起invalid
receiver type ... 的编译器错误。来自语言说明书[10] 的权威内容：
接收者类型必须是T 或*T，这里的T 是类型名。T 叫做接收者基础类型或
简称基础类型。基础类型一定不能使指针或接口类型，并且定义在与方法
相同的包中。
*/

//根据规则，单方法接口命名为方法名加上-er 后缀：Reader，Writer，Formatter 等。

//func sort(i []interface{}) {
//	switch i.(type) {
//	case string:
//	//...
//	case int:
//		//...
//	}
//	return /* */
//}

//这是因为Go 不能简单的将其转换为接口的slice。转换到接口是容易的，但是转换到
//slice 的开销就高了。
//关于这个话题完 简单来说 ：Go 不能（隐式）转换为 slice。
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

//定义用于排序slice 的新类型。注意定义的是slice 类型；
type Xi []int
type Xs []string

func (p Xi) Len() int {
	return len(p)
}
func (p Xi) Less(i int, j int) bool {
	return p[j] < p[i]
}
func (p Xi) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p Xs) Len() int {
	return len(p)
}
func (p Xs) Less(i int, j int) bool {
	return p[j] < p[i]
}
func (p Xs) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func Sort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}

//reflect
type Person struct {
	Name string "namestr" //"namestr" 是标签
	age  int
}

//func ShowTag(i interface{}){
//	switch t := reflect.TypeOf(i);t.Kind(){
//		case reflect.Ptr:
//		tag := t.Elem().Field(0).Tag
//	}
//}

func show(i interface{}) {
	switch h := i.(type) {
	case *Person:
		//h := reflect.TypeOf(i)
		v := reflect.ValueOf(i)
		//tag := h.Elem().Field(0).Tag
		name := v.Elem().Field(0).String()
	}
}

func main() {
	var p Person
	show(p)
	//var s S
	//f(&s)
	//fmt.Println("Hello World!")
	s := new(S)
	fmt.Println(g(s)) //0
	var i Foo
	i = 10
	i.Emit()
	ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	strings := Xs{"nut", "ape", "elephant", "zoo", "go"}
	Sort(ints)
	fmt.Printf("%v", ints)
	Sort(strings)
	fmt.Printf("%v", strings)
}
