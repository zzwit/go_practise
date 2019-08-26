package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("Not found error")



func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New ("math - square root of negative number")
	}
	// implementation of Sqrt
}


func main() {
	fmt.Printf("error:%v", errNotFound)
	if f, err := Sqrt(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
