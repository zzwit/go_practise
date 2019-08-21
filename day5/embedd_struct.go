package main

import "fmt"

/**
内嵌结构体
 */

type A struct {
	ax ,ay int
}

type B struct {
	A
	bx ,by int
}

func main()  {
	b := B{A{1,2},3.0,4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}