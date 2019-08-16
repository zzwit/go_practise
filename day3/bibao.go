package main

import "fmt"

// 闭包学习
/**
当我们不希望给函数起名字的时候，可以使用匿名函数，例如：func(x, y int) int { return x + y }。
这样的一个函数不能够独立存在（编译器会返回错误：non-declaration statement outside function body），但可以被赋值于某个变量，
即保存函数的地址到变量中：fplus := func(x, y int) int { return x + y }，然后通过变量名对函数进行调用：fplus(3,4)。
当然，您也可以直接对匿名函数进行调用：func(x, y int) int { return x + y } (3, 4)。
下面是一个计算从 1 到 1 百万整数的总和的匿名函数：
 */

func main()  {
	f()
}

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}
