package main

import "fmt"

//如果参数被存储在一个 slice 类型的变量 slice 中，则可以通过 slice... 的形式来传递参数，调用变参函数。
func main() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7,9,3,5,1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)
}

func min(s ...int) int {
	if len(s)==0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		fmt.Printf("this is v = : %d m: = %d\n",v,min)
		if v < min {
			min = v
		}
	}
	return min
}