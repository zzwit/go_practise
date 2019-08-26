package main

import (
	"fmt"
	"os"
	"strings"
)

/***
	- 从命令行获取参数
	os包中有一个字符串类型的切片变量os.Args，用来处理一些基本的命令行参数，它在程序启动后读取命令行输入的参数。
 */
func main() {
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], "")
	}
	fmt.Println("Good Morning", who)
}
