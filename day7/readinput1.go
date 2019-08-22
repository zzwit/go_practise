package main

import "fmt"

/**
  用户输入
	- Scan 和 Sscan 开头的函数
 */
var (
	firstName, lastName, s string
	i                      int
	f                      float64
	input                  = "56.12/ 5212/ Go"
	format                 = "%f/%d/%s"
)

func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf(firstName, "firstName")
	fmt.Printf(lastName, "lastName")
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}
