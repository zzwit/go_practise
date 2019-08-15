package main
import (
	"io"
	"log"
)

//下面的代码展示了另一种在调试时使用 defer 语句的手法
func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

func main() {
	func1("Go")
}