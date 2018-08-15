package main

import "net"
import "fmt"
import "bufio"
import (
	"os"
	"unsafe"
)

const(
	EM_DLL_DYNAMICPWD_CHECK_SUCCESS = iota + 2000		//动态口令校验是否正确
	EM_DLL_DCSCONFINFO
	EM_DLL_HASMTCSKYSHARE
	EM_DLL_VOD
	EM_DLL_CALLING_PWD
)

type tcpMsg struct {
	eventID 	int
	messages 	string
}

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", ":8081")

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		var msgSt tcpMsg = tcpMsg{EM_DLL_DYNAMICPWD_CHECK_SUCCESS,"Hello World!"}
		data := *(*[]byte)(unsafe.Pointer(&msgSt))
		// send to socket
		//fmt.Fprintf(conn, text + "\n")
		if conn != nil{
			conn.Write(data)
			// listen for reply
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Message from server: "+message)
		}
	}
}