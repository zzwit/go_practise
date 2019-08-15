package main

import "fmt"

//使用 for 结构创建一个简单的循环。要求循环 15 次然后使用 fmt 包来打印计数器的值。
//使用 goto 语句重写循环，要求不能使用 for 关键字。
func main() {
	// 1:
	for i := 0; i < 15; i++ {
		fmt.Printf("The counter is at %d\n", i)
	}
	// 2:
	i := 0
	// 使用的 goto
START:
	fmt.Printf("The counter is at %d\n", i)
	i++
	if i < 15 {
		goto START
	}
}