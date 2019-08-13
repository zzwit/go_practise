package main

import "fmt"

// 控制结构

//var prompt = "Enter a digit, e.g. 3 "+ "or %s to quit."
//
//func init() {
//	if runtime.GOOS == "windows" {
//		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
//	} else { //Unix-like
//		prompt = fmt.Sprintf(prompt, "Ctrl+D")
//	}
//}

func main()  {
	bool1 := true
	if bool1 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}

}