package main

import "net"
import "fmt"
import (
	"bufio"
	"unsafe"
)
// only needed below for sample processing

var mapEventFunc map[string]func(str string)
var(
	sysMsg = make(chan string,10)
	asysMsg = make(chan string,100)
)

type tcpMsg struct {
	eventID 	int
	messages 	string
}

func main() {

	fmt.Println("Launching server...")

	Init()

	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()

	go func() {
		for{
			newmessage := <-sysMsg
			conn.Write([]byte(newmessage + "\n"))
		}
	}()

	// run loop forever (or until ctrl-c)
	for {
/*		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("\nMessage Received:", string(message))
		for recievedMsg := range mapEventFunc{
			if recievedMsg + "\n" == string(message){
				mapEventFunc[recievedMsg](recievedMsg)
			}
		}*/
		var data []byte
		reader := bufio.NewReader(conn)
		reader.Read(data)

		var ptestStruct *tcpMsg = *(**tcpMsg)(unsafe.Pointer(&data))
		fmt.Println(ptestStruct.messages)
/*		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))*/
	}
}

func Init(){
	mapEventFunc = make(map[string]func(str string))
	mapEventFunc["aaa"] = OnEvent1
	mapEventFunc["bbb"] = OnEvent2
	mapEventFunc["ccc"] = OnEvent3
}

func OnEvent1(strEvent string){
	sysMsg<- "OnEvent1"
	fmt.Println("[OnEvent1]",strEvent,"has been handled!!")
}

func OnEvent2(strEvent string){
	sysMsg<- "OnEvent2"
	fmt.Println("[OnEvent2]",strEvent,"has been handled!!")
}

func OnEvent3(strEvent string){
	sysMsg<- "OnEvent3"
	fmt.Println("[OnEvent3]",strEvent,"has been handled!!")
}
