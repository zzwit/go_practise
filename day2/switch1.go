package main

// shitch 学习
import "fmt"

func main() {
	var num1 int = 99

	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}

	switch {
	case num1 < 0:
		fmt.Println("Number is negative")
	case num1 > 0 && num1 < 100:
		fmt.Println("Number is between 0 and 100")
	default:
		fmt.Println("Number is 10 or greater")
	}

	// fallthrough：Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}


}