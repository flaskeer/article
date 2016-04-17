// ch6 project main.go
package main

/*
关于并发
叫做goroutine 是因为已有的短语——线程、协程、进程等等——传递了不
准确的含义。goroutine 有简单的模型：它是与其他goroutine 并行执行的，
有着相同地址空间的函数。。它是轻量的，仅比分配栈空间多一点点􂜷􃧆。
而初始时栈是很小的，所以它们也是􁡸价的，并且随着需要在堆空间上分
配（和释放）。
goroutine 是一个普通的函数，只是需要使用关键字go 作为开头。

如果不等待goroutine 的执行（例如，移除第17 行），程序立刻终止，而任何正在执行
的goroutine 都会停止。为了修复这个，需要一些能够同goroutine 通讯的机制。这一
机制通过channels 的形式使用。channel 可以与Unix sehll 中的双向管道做类比：可以
通过它发送或者接收值。这些值只能是特定的类型：channel 类型。定义一个channel
时，也需要定义发送到channel 的值的类型。注意，必须使用make 创建channel：
*/

import (
	"fmt"
	"time"
)

//ci := make(chan int)
//cs := make(chan string)
//cf := make(chan interface{})

var c chan int

/*
定义c 作为int 型的channel。就是说：这个channel 传输整数。注意这个变量是
全局的，这样goroutine 可以访问它；
..1 发送整数1 到channel c；
..2 初始化c；
..3 用关键字go 开始一个goroutine；
..4 等待，直到从channel 上接收一个值。注意，收到的值被丢弃了；
..5 两个goroutines，接收两个值。

虽然goroutine 是并发执行的，但是它们并不是并行运行的。如果不告诉Go 额外的
东西，同一时刻只会有一个goroutine 执行。利用runtime.GOMAXPROCS(n) 可以设置
goroutine 并行执行的数量。来自文档：
GOMAXPROCS 设置了同时运行的CPU 的最大数量，并返回之前的设置。如
果n < 1，不会改变当前设置。当调度得到改进后，这将被移除。

当在Go 中用ch := make(chan bool) 创建chennel 时，bool 型的无缓冲channel 会
被创建。这对于程序来说意味着什么呢？首先，如果读取（value := <ch）它将会
被阻塞，直到有数据接收。其次，任何发送（ch< 5）将会被阻塞，直到数据被读出。
无缓冲channel 是在多个goroutine 之间同步很棒的工具。
*/
func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready")
	c <- 1
}

func main() {
	//I am waiting
	//coffee is ready
	//tea is ready
	//ci <- 1
	//i := <-ci
	c = make(chan int)
	i := 0
	go ready("tea", 2)
	go ready("coffee", 1)
	fmt.Println("I am waiting but not too long")
	//time.Sleep(6 * time.Second)
	//<-c
	//<-c
	//select。通过select（和其他东西）可以监听channel 上输入的数据。
L:
	for {
		select {
		case <-c:
			i++
			if i > 1 {
				break L
			}
		}
	}
}
