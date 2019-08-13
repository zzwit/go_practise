package main

import "fmt"

// 指针简单学习

/**
一个指针变量可以指向任何一个值的内存地址 它指向那个值的内存地址，在 32 位机器上占用 4 个字节，在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关。
当然，可以声明指针指向任何类型的值来表明它的原始性或结构性；你可以在指针类型前面加上 * 号（前缀）来获取指针所指向的内容，
这里的 * 号是一个类型更改器。使用一个指针引用一个值被称为间接引用。
 */


func main()  {
	var i1 = 5
	// i1 解释 ，然后使用 intP = &i1 是合法的，此时 intP 指向 i1。
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1) //指针的格式化标识符为 %p
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p) // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s) // prints same string
	//通过对 *p 赋另一个值来更改“对象”，这样 s 也会随之更改。
	/***
	指针的一个高级应用是你可以传递一个变量的引用（如函数的参数），这样不会传递变量的拷贝。指针传递是很廉价的，
	只占用 4 个或 8 个字节。当程序在工作中需要占用大量的内存，或很多变量，或者两者都有，使用指针会减少内存占用和提高效率。被指向的变量也保存在内存中，
	直到没有任何指针指向它们，所以从它们被创建开始就具有相互独立的生命周期。
	另一方面（虽然不太可能），由于一个指针导致的间接引用（一个进程执行了另一个地址），指针的过度频繁使用也会导致性能下降。
	指针也可以指向另一个指针，并且可以进行任意深度的嵌套，导致你可以有多级的间接引用，但在大多数情况这会使你的代码结构不清晰。
	如我们所见，在大多数情况下 Go 语言可以使程序员轻松创建指针，并且隐藏间接引用，如：自动反向引用。
	对一个空指针的反向引用是不合法的，并且会使程序崩溃：
	 */
}
