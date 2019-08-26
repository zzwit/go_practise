package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

/***
  使用 用 buffer 读取文件
 */

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}
func main() {
	flag.Parse() //对所有的参数进行扫描
	if flag.NArg() == 0 {
		// 查看文件的内容
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}

}
