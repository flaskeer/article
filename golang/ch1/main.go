// ch1 project main.go
package main

import (
	"fmt"
	"unicode/utf8"
)

/*
%v   基本格式的值。当输出结构体时，扩展标志(%+v)添加成员的名字。the value in a default format.
     when printing structs, the plus flag (%+v) adds field names

*/
func main() {
	// solve1
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	//solve2
	//i := 0
	//Loop:
	//fmt.Printf("%d\n", i)
	//if i < 10 {
	//	i++
	//	goto Loop
	//}

	//solve3
	//var arr [10]int
	//for i := 0; i < 10; i++ {
	//	arr[i] = i
	//}
	//fmt.Printf("%v", arr)
	/*
		BUZZ

		26

		FIZZ

		28

		29

		FIZZBUZZ
	*/
	/*
		const (
			FIZZ = 3
			BUZZ = 5
		)
		var p bool
		for i := 1; i < 100; i++ {
			p = false
			if i%FIZZ == 0 {
				fmt.Printf("FIZZ")
				p = true
			}
			if i%BUZZ == 0 {
				fmt.Printf("BUZZ")
				p = true
			}
			if !p {
				fmt.Printf("%v", i)
			}
			fmt.Println()
			fmt.Println()
		}
	*/

	//str := "A"
	//for i := 0; i < 100; i++ {
	//	fmt.Printf("%s\n", str)
	//	str = str + "A"
	//}
	str := "sdfgfdhfgjhggfkjhkf"
	fmt.Printf("string %s\nLength:%d,Runes:%d", str, len([]byte(str)), utf8.RuneCount([]byte(str)))
	s := "??? ??????????? "
	r := []rune(s)
	copy(r[4:4+3], []rune("abc"))
	fmt.Printf("before:%s", s)
	fmt.Printf("after:%s", string(r))
	//reverse
	s1 := "foobar"
	a := []rune(s1)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Printf("%s\n", string(a))

}
