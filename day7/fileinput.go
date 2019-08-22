package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/***
 文件的读写
 	- 文件使用指向 os.File 类型的指针来表示的，也叫做文件句柄。我们在前面章节使用到过标准输入 os.Stdin 和标准输出 os.Stdout

 */

func main()  {
	inputFile,inputError := os.Open("products.txt")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}