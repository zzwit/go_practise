package main

import (
	"fmt"
)


//函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调。下面是一个将函数作为参数的简单
func main() {
	callback(1, Add)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	 // y = 1 f Add
	 // Add(1,2)
	f(y, 2) // this becomes Add(1, 2)
}

func IsAscii(c int) bool {
	if c > 255 {
		return false
	}
	return true
}
