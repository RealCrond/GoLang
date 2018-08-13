package main

import (
	"time"
	"fmt"
)

func calc() {
	for i := 0;i < 10; i++ {
		time.Sleep(1*time.Second)
		fmt.Println("run ", i, " times")
	}
	fmt.Println("calc finish")
}

func calcEx() {
	for i := 0;i < 10; i++ {
		time.Sleep(1*time.Second)
		fmt.Println("HeiHei ", i, " times")
	}
	fmt.Println("HeiHei finish")
}

func main() {
	go calc()
	go calcEx()
	fmt.Println("i exited")
	time.Sleep(21*time.Second)
}