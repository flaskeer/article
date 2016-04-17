// ch3 project main.go
package main

import (
	"even"
	"event"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(event.Even(3))
	fmt.Println(even.Odd(3))
}
