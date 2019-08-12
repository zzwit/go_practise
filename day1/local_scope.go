package main
var a = "G"

func main() {
	n()
	m()
	n()
}

func n() { print(a) }

func m() {
	// 初始化函数，局部修改变量不会影响原来的 全局变量的数值
	a := "O"
	print(a)
}