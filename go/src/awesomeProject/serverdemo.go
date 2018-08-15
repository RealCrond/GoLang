package main

import (
	"net"
	"fmt"
	"os"
	"encoding/json"
	"bufio"
	"io"
	"hash/crc32"
)

//https://www.jianshu.com/p/dbc62a879081
//数据包类型
const (
	HEART_BEAT_PACKET = 0x00
	REPORT_PACKET = 0x01
)

//默认的服务器地址
var (
	server = "127.0.0.1:9876"
)

//数据包
type Packet struct {
	PacketType      byte
	PacketContent   []byte
}

//心跳包
type HeartPacket struct {
	Version     string`json:"version"`
	Timestamp   int64`json:"timestamp"`
}

//数据包
type ReportPacket struct {
	Content   string`json:"content"`
	Rand         int`json:"rand"`
	Timestamp   int64`json:"timestamp"`
}

type TcpServer struct {
	listener *net.TCPListener
	hawkServer *net.TCPAddr
}

func main(){
	hawkServer,err := net.ResolveTCPAddr("tcp",server)
	checkErr(err)

	listen,err := net.ListenTCP("tcp",hawkServer)
	checkErr(err)

	defer listen.Close()
	tcpServer := &TcpServer{
		listener:listen,
		hawkServer:hawkServer,
	}
	fmt.Println("Start server successful.....")

	for{
		conn,err := tcpServer.listener.Accept()
		fmt.Println("accept tcp client %s",conn.RemoteAddr().String())
		checkErr(err)

		go Handle(conn)
	}

}

func Handle(conn net.Conn){
	defer conn.Close()
	state := 0x00
	length := uint16(0)
	// crc校验和
	crc16 := uint16(0)
	var recvBuffer []byte

	cursor := uint16(0)
	bufferReader := bufio.NewReader(conn)

	for{
		recvByte,err := bufferReader.ReadByte()
		if err != nil{
			if err == io.EOF{
				fmt.Printf("client %s is close!\n",conn.RemoteAddr().String())
			}
			return 
		}

		switch state {
		case 0x00:
			if recvByte == 0xFF{
				state = 0x01
				recvBuffer = nil
				length = 0
				crc16 = 0
			}else{
				state = 0x00
			}
			break
		case 0x01:
			if recvByte == 0xFF{
				state = 0x02
			}else{
				state = 0x00
			}
			break
		case 0x04:
			recvBuffer[cursor] = recvByte
			cursor++
			if(cursor == length){
				state = 0x05
			}
			break
		case 0x05:
			crc16 += uint16(recvByte) * 256
			state = 0x06
			break
		case 0x06:
			crc16 += uint16(recvByte)
			state = 0x07
			break
		case 0x07:
			if recvByte == 0xFF {
				state = 0x08
			}else{
				state = 0x00
			}
		case 0x08:
			if recvByte == 0xFE{
				if (crc32.ChecksumIEEE(recvBuffer) >> 16) & 0xFFFF == uint32(crc16) {
					var packet Packet
					json.Unmarshal(recvBuffer,&packet)
					go processRecvData(&packet,conn)
				}else{
					fmt.Println("丢弃数据!")
				}
			}
			state = 0x00
		}
	}

}

func processRecvData(packet *Packet,conn net.Conn){
	switch packet.PacketType {
	case HEART_BEAT_PACKET:
		var beatPacket HeartPacket
		json.Unmarshal(packet.PacketContent,&beatPacket)
		fmt.Printf("Recieve heat beat from [%s],data is [%v]\n",conn.RemoteAddr().String(),beatPacket)
		conn.Write([]byte("heartBeat\n"))
		return
	case REPORT_PACKET:
		var reportPacket ReportPacket
		json.Unmarshal(packet.PacketContent,&reportPacket)
		fmt.Printf("Recieve heat beat from [%s],data is [%v]\n",conn.RemoteAddr().String(),reportPacket)
		conn.Write([]byte("Report data has recive\n"))
		return
	}
}

func checkErr(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(-1)
	}
}