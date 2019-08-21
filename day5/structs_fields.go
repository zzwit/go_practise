package main

import "fmt"


/**
var t *T
t = new(T)
写这条语句的惯用方法是：t := new(T)，变量 t 是一个指向 T的指针，此时结构体字段的值是它们所属类型的零值。
 */
type struct1 struct {
	i1  int
	f1  float32
	str string
}


type myStruct struct { i int }

func main() {
	//ms := new(struct1)
	//ms.i1 = 10
	//ms.f1 = 15.5
	//ms.str= "Chris"
	//ms.sfv= "Chris"
	ms := &struct1{10, 15.5, "Chris"}
	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)
	//使用 fmt.Println 打印一个结构体的默认输出可以很好的显示它的内容，类似使用 %v 选项。
}
