// ch2 project main.go

/*
func (p mytype) funcname(q i n t ) (r,s i n t ) { return 0,0 }
..0 关键字func 用于定义一个函数；
..1 函数可以绑定到特定的类型上。这叫做接收者。有接收者的函数被称作method。
第5 章将对其进行说明；
..2 funcname 是你函数的名字；
..3 int 类型的变量q 作为输入参数。参数用pass-by-value 方式传递，意味着它们会
被复制；
..4 变量r 和s 是这个函数的命名返回值。在Go 的函数中可以返回多个值。参阅
第28 页的“多值返回”。如果不想对返回的参数命名，只需要提供类型：(int,int)。
如果只有一个返回值，可以省略圆括号。如果函数是一个子过程，并且没有任何
返回值，也可以省略这些内容；
..5 这是函数体。注意return 是一个语句，所以包裹参数的括号是可选的。

在Go 中，定义在函数外的变量是全局的，那些定义在函数内部的变量，对于函数来说
是局部的。如果命名覆盖——一个局部变量与一个全局变量有相同的名字——在函数
执行的时候，局部变量将覆盖全局变量。

Go 一个非常特别的特性（对于编译语言而言）是函数和方法可以返回多个值（Python
和Perl 同样也可以）。这可以用于改进一大堆在C 程序中糟糕的惯例用法：修改参数
的方式，返回一个错误（例如遇到EOF 则返回-1）。在Go 中，Write 返回一个计数值
和一个错误：“是的，你写入了一些字节，但是由于设备异常，并不是全部都写入了。
”。os 包中的*File.Write 是这样声明的：
func (file *File) Write(b []byte) (n int , err e r r o r )
如同文档所述，它返回写入的字节数，并且当n != len(b) 时，返回非nil 的error。
这是Go 中常见的方式。
元组没有作为原生类型出现，所以多返回值可能是最佳的选择。你可以精确的返回希
望的值，而无须重载域空间到特定的错误信号上。

Go 函数的返回值或者结果参数可以指定一个名字，并且像原始的变量那样使用，就像
输入参数那样。如果对其命名，在函数开始时，它们会用其类型的零值初始化。如果
函数在不加参数的情况下执行了return 语句，结果参数会返回。用这个特性，允许
（再一次的）用较少的代码做更多的事a。
*/
package main

import (
	"fmt"
)

func Even(i int) bool {
	return i%2 == 0
}

var a int = 6

func p() {
	a = 5 //赋值  全局范围内可见
}

func q() {
	a := 5 //局部变量 函数内可见
	fmt.Printf("%d", a)
}

//没有返回值，只是简单的将输入返回
func identity(in int) int { return in }

//多返回值
//func ReadFull(r Reader,buf []byte) (n int,err error){
//	for len(buf) > 0 && err == nil{
//		var nr int
//		nr,err = r.Read(buf)
//		n += nr
//		buf = buf[nr:len(buf)]
//	}
//}

/*
Go 有了defer 语句。在defer 后指定的
函数会在函数退出前调用。
*/
//func ReadWrite() bool{
//	file.open("file")
//	defer file.close()
//	if failureX{
//		return false
//	}
//	if failureY{
//		return false
//	}
//	return true
//}

func rec(i int) {
	if i == 10 {
		return
	}
	rec(i + 1)
	fmt.Println(i)
}

//利用defer 甚至可以修改返回值  return 1
func f() (ret int) {
	defer func() {
		ret++

	}()
	return 0
}

func vargs(arg ...int) {
	for _, n := range arg {
		fmt.Printf("and the number is:%d\n", n)
	}
}

////由于函数也是值，所以可以很容易的传递到其他函数里，然后可以作为回调。
func printit(x int) {
	fmt.Printf("%v\n", x)
}

func callback(y int, f func(int)) {
	f(y)
}

/*
Panic
是一个内建函数，可以中断原有的控制流程，进入一个令人恐慌的流程中。当函
数F 调用panic，函数F 的执行被中断，并且F 中的延迟函数会正常执行，然
后F 返回到调用它的地方。在调用的地方，F 的行为就像调用了panic。这一过
程继续向上，直到程序崩溃时的所有goroutine 返回。
恐慌可以直接调用panic 产生。也可以由运行时错误产生，例如访问越界的数
组。
Recover
是一个内建的函数，可以让进入令人恐慌的流程中的goroutine 恢复过来。recover
仅在延迟函数中有效。
在正常的执行过程中，调用recover 会返回nil 并且没有其他任何效果。如果
当前的goroutine 陷入恐慌，调用recover 可以捕获到panic 的输入值，并且恢
复正常的执行。
*/
/*
定义一个新函数throwsPanic 接受一个函数作为参数（参看“函数作为值”）。函
数f 产生panic，就返回true，否则返回false；
定义了一个利用recover 的defer 函数。如果当前的goroutine 产生了panic，
这个defer 函数能够发现。当recover() 返回非nil 值，设置b 为true；
调用作为参数接收的函数。
返回b 的值。由于b 是命名返回值*/

func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
}

func main() {
	//延迟的函数是按照后进先出（LIFO）的顺序执行，所以上面的代码打印：4 3 2 1 0。
	for i := 0; i < 5; i++ {
		defer fmt.Printf("---%d----", i)
	}
	def := f()
	fmt.Println("val defer is :", def)
	q()
	a := identity(10)
	fmt.Println("Hello World!", a)
	rec(0)
	//arr := [5]int{1, 2, 3, 4, 5}

	vargs(1, 2, 3, 4, 5, 6, 7, 87, 8)
	//vargs(arr)   cannot use arr (type [5]int) as type int in argument to vargs
	//就像其他在Go 中的其他东西一样，函数也是值而已
	method := func() {
		println("hello 匿名函数")
	}
	method()
	fmt.Printf("%T\n", method) //func()

	var xs = map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}
	//k is: 1 :val is : 10
	//k is: 2 :val is : 20
	//k is: 3 :val is : 30
	for k, v := range xs {
		fmt.Println("k is:", k, ":val is :", v())
	}

}
