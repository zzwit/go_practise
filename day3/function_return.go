package main

import "fmt"

func main() {
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 2:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
	var f = Adder82()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))



}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func Adder82() func(int) int {
	var x int
	// x的数值一直的变化的，每次都是上一次 x  + = delta的数值
	//我们可以看到，在多次调用中，变量 x 的值是被保留的，即 0 + 1 = 1，然后 1 + 20 = 21，最后 21 + 300 = 321：闭包函数保存并积累其中的变量的值，不管外部函数退出与否，
	// 它都能够继续操作外部函数中的局部变量。
	return func(delta int) int {
		fmt.Printf("this is delta = %d \n",delta)
		fmt.Printf("this is x = %d \n",x)
		x += delta
		return x
	}
}