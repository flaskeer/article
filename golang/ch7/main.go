// ch7 project main.go
package main

//import (
//	"bufio"
//	"os"
//)

//func main() {
//	buf := make([]byte, 1024)
//	f, _ := os.Open("d://test.txt")
//	defer f.Close()
//	//缓冲代码
//	r := bufio.NewReader(f)
//	w := bufio.NewWriter(os.Stdout)
//	defer w.Flush()
//	for {
//		//n, _ := f.Read(buf)
//		n, _ := r.Read(buf)
//		if n == 0 {
//			break
//		}
//		//os.Stdout.Write(buf[:n])
//		w.Write(buf[0:n])
//	}
//}
import ( // 引入了代码包fmt和runtime
	"fmt"
	"runtime"
)

/**
输出的第一行是对变量m格式化后的结果。这就意味着，在函数init的第一条语句执行时，
变量m已经被初始化并赋值了。这验证了一条规则：
当前代码包中所有全局变量的初始化会在代码包初始化函数执行前完成。

输出的第二行是对变量info格式化后的结果。变量info被定义时并没有被显式赋值，因此它被赋予类型string的零值——""（空字符串）。
之后，变量info在代码包初始化函数init中被赋值，并在入口函数main中被输出。可见，所有的包初始化函数都会在main函数之前执行完成。
*/
func init() { // 包初始化函数
	fmt.Printf("Map: %v\n", m) // 先格式化再打印
	// 通过调用runtime包的代码获取当前机器所运行的操作系统以及计算架构
	// 而后通过fmt包的Sprintf方法进行字符串格式化并赋值给变量info
	info = fmt.Sprintf("OS: %s, Arch: %s", runtime.GOOS, runtime.GOARCH)
}

var m map[int]string = map[int]string{1: "A", 2: "B", 3: "C"}

// 非局部变量，map类型，已被显式赋值

var info string // 非局部变量，string类型，未被显式赋值

func main() { // 命令源码文件必须有的入口函数
	fmt.Println(info) // 打印变量info
}
