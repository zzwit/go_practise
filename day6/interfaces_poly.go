package main

import "fmt"

/***
  实现的两个接口
 */


type Shaper interface {
	Area() float32
}

type Square struct {
	side float32 // 5
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32  // 5 3
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}


func main() {

	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []Shaper{r, q}  //15
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println(n,'n')
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}