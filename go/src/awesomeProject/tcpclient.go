package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
/*	currency := 20 //并发数,记住，一个连接数是打开一个端口号，window和linux的端口号都是有限制的
	count := 10 //每条连接发送多少次连接
	for i:=0;i<currency;i++ {
		go func(){
			for j:=0;j<count;j++ {*/
			var sendMsg string
			for {
				reader := bufio.NewReader(os.Stdin)
				data,_,_ := reader.ReadLine()
				sendMsg = string(data)
				fmt.Println("Send Msg:",sendMsg)
				sendMessage(sendMsg)
			}
/*			}
		}()
	}*/
	select{}
}

func sendMessage(str string) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if(err != nil) {
		panic("error")
	}
	//header := "hello"
	//header := "Hello Server!!"
	fmt.Fprintf(conn, str)
	//var str []byte = []byte{'H','e','l','l','o'}
	//conn.Write([]byte{'H','e','l','l','o'})
}