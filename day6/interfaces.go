package main

import "fmt"

/**
 实现的接口的方式
 */


 // 定义一个接口
type Shaper interface {
	Area() float32
}

//定义了一个结构体
type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}


func main() {
	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper
	areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}
