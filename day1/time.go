package main

import (
	"fmt"
	"time"
)

// 日期的学习

var week time.Duration

func main()  {
	t := time.Now()
	fmt.Println(t) //2019-08-12 17:59:27.832025 +0800 CST m=+0.000345058
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	//12.08.2019
	t = time.Now().UTC() //设置时区
	fmt.Println(t) //2019-08-12 10:01:13.863511 +0000 UTC
	fmt.Println(time.Now()) //2019-08-12 18:01:13.863531 +0800 CST m=+0.00082964
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now) //2019-08-19 10:02:50.279873 +0000 UTC
	fmt.Println(t.Format(time.RFC822)) // 12 Aug 19 10:03 UTC
	fmt.Println(t.Format(time.ANSIC)) //Mon Aug 12 10:03:43 2019
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 12 Aug 2019 10:03
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
}
