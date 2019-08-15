package main

//Break 与 continue

func main()  {
	for i:=0; i<3; i++ {
		for j:=0; j<10; j++ {
			if j>5 {
				break
			}
			print(j)
		}
		print("  ")
	}
	//关键字 continue 忽略剩余的循环体而直接进入下一次循环的过程，但不是无条件执行下一次循环，执行之前依旧需要满足循环的判断条件。

	for i :=0;i<10; i++ {
		if i==5 {
			continue
		}
		print(i)
		print(" ")
	}
}
