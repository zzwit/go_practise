package main

func main() {
	// 1 - use 2 nested for loops
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			print("G")
		}
		print("\n")
		//println()
	}
	// 2 -  use only one for loop and string concatenation
	/**
	 println 和 Printf 区别
	 println 可以打印出字符串，和变量
	 Printf : 只可以打印出格式化的字符串,可以输出字符串类型的变量，不可以输出整形变量和整形
	 */
	str := "G"
	for i := 1; i <= 25; i++ {
		println(str)
		str += "G"
	}
}