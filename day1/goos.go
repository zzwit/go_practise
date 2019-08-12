package main

import (
	"fmt"
	"os"
	"runtime"
)

//如何通过runtime包在运行时获取所在的操作系统类型，以及如何通过 os 包中的函数 os.Getenv() 来获取环境变量中的值，并保存到 string 类型的局部变量 path 中
func main()  {
	var goos string = runtime.GOOS
	fmt.Printf("This operating is :%s \n",goos)
	path :=os.Getenv("PATH")
	fmt.Printf("Path is %s \n",path)
	//简短形式，使用 := 赋值操作符
	a :=50
	fmt.Printf("a:%s \n",a)
	//a :=20
	//fmt.Printf("a:%s \n",a)
}