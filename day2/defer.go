package main

import "fmt"

func main() {
	function1()
}


/**
关键字 defer 允许我们进行一些函数执行完成后的收尾工作，例如：
1.关闭文件流
	defer file.Close()
2. 解锁一个加锁的资源
	mu.Lock()
	defer mu.Unlock()
3.打印最终报告
	printHeader()
	defer printFooter()
4.关闭数据库链接
	defer disconnectFromDB()
*/
//defer 和追踪
func function1() {
	fmt.Printf("In function1 at the top\n")
	//关键字 defer 允许我们进行一些函数执行完成后的收尾工作，例如：
	defer function2() //执行完所有的函数在执行的defer的函数
	function3()
	a()
	f()
	doDBOperations()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function! \n")
}

func function3()  {
	var a , c  = 1, 3
	fmt.Printf("a+c= %d \n",a + c)
}

func a() {
	i := 0
	defer fmt.Println("a =",i)
	i++
	return
}


//当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）：
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("f = %d \n", i)
	}
}


func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() //function called here with defer
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return //terminate the program
	// deferred function executed here just before actually returning, even if
	// there is a return or abnormal termination before
}
