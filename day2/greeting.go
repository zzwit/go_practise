package main

// 简单的函数的学习

func main()  {
	println("In main before calling greeting")
	greeting()
	println("In main after calling greeting")
}

func greeting()  {
	println("In greeting: Hi!!!!!")
}