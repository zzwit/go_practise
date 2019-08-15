package main

import "fmt"

const (
	FIZZ     = 3
	BUZZ     = 5
	FIZZBUZZ = 15
)

//写一个从 1 打印到 100 的程序，但是每当遇到 3 的倍数时，不打印相应的数字，但打印一次 "Fizz"。遇到 5 的倍数时，打印 Buzz 而不是相应的数字。对于同时为 3 和 5 的倍数的数，打印 FizzBuzz（提示：使用 switch 语句）。
func main() {
	for i := 0; i <= 100; i++ {
		switch {
		case i%FIZZBUZZ == 0:
			fmt.Println("FizzBuzz")
		case i%FIZZ == 0:
			fmt.Println("Fizz")
		case i%BUZZ == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
