package main

import (
	"fmt"
	"time"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"

	ret int
	err error
)

func main() {
	//ret,err := fmt.Printf("Yournameis%s.Yourageis%d\n","Herman",26)
	//fmt.Println(ret,err)
	//fmt.Println("Please input your full name: ")
	t, _ := time.Parse("2006-01-02 15:04:05", "2016-04-20 16:23:00")
	fmt.Println(t.Unix())
	y, m, d := time.Unix(1466344320, 0).Date()
	fmt.Println(y, m, d)

	//format后面的字符串必须是2006-01-02 15:04:05，据说go是这个时间诞生的
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02"))
	fmt.Println(time.Now().Format("20060102"))

	a, b, c := "", 0, false
	fmt.Scan(&a, &b, &c)
	fmt.Println(a, b, c)

	fmt.Scanln(&a, &b, &c)
	fmt.Println(a, b, c)

	fmt.Scanf("%3s%d%t", &a, &b, &c)
	fmt.Println(a, b, c)

	//fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf(“%s %s”, &firstName, &lastName)
	//fmt.Printf("Hi %s %s!\n", firstName, lastName)
	//fmt.Sscanf(input, format, &f, &i, &s)
	//fmt.Println("From the string we read: ", f, i, s)
	//fmt.Scanln(&firstName, &lastName)

}
