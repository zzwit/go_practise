package main

import "fmt"

// for GOTO 用法
//for、switch 或 select 语句都可以配合标签（label）形式的标识符使用，即某一行第一个以冒号（:）结尾的单词（gofmt 会将后续代码自动移至下一行）。

func main()  {
LABEL1:
	for i :=0;i<5 ; i++ {
		for j :=0;j<=5 ; j++  {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	//本例中，continue 语句指向 LABEL1，当执行到该语句的时候，就会跳转到 LABEL1 标签的位置。
	/***
	您可以看到当 j==4 和 j==5 的时候，没有任何输出：标签的作用对象为外部循环，因此 i 会直接变成下一个循环的值，而此时 j 的值就被重设为 0，即它的初始值。
	如果将 continue 改为 break，则不会只退出内层循环，而是直接退出外层循环了。另外，还可以使用 goto 语句和标签配合使用来模拟循环。
	 */

 //i :=0
 //HERE:
	//print(i)
	//i++
	//if i ==5 {
	//	return
	//}
	//goto HERE
 // 特别注意 使用标签和 goto 语句是不被鼓励的：它们会很快导致非常糟糕的程序设计，而且总有更加可读的替代方案来实现相同的需求。

	for i := 0; i<7 ; i++ {
		if i%2 == 0 { continue }
		fmt.Println("Odd:", i)
	}
}