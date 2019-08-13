package main

import "fmt"

//类型别名 在 type TZ int 中，TZ 就是 int 类型的新名称

type TZ int

func main() {
	var a, b TZ = 3, 4
	c := a + b
	fmt.Printf(" c has thie value: %d", c)
}
