package main

import (
	"fmt"
	"./trans"
)

var twoPi = 2 * trans.Pi

// user_init.go 中导入了包 trans（需要init.go目录为./trans/init.go）并且使用到了变量 Pi：
func main() {
	fmt.Printf("2*Pi = %g\n", twoPi) // 2*Pi = 6.283185307179586
}