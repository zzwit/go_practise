package main

var a string

func main() {
	a = "G" // 修改全局变量的数值
	print(a)
	f1()
}

func f1() {
	a := "O" // 指定当前的数值，不会影响全局的数值
	print(a)
	f2()
}

func f2() {
	print(a)
}