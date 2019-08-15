package main
import "fmt"

/**
在使用递归函数时经常会遇到的一个重要问题就是栈溢出：一般出现在大量的递归调用导致的程序栈内存分配耗尽。这个问题可以通过一个名为懒惰求值的技术解决，
在 Go 语言中，我们可以使用管道（channel）和 goroutine（详见第 14.8 节）来实现。练习 14.12 也会通过这个方案来优化斐波那契数列的生成问题。

 */
func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}