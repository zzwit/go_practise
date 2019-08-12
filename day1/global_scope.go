package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	a = "O" // 修改的变量 也修改全局边的数值
	print(a)
}