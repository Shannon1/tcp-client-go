package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"github.com/golang/protobuf/proto"
	"testpbgo"
)


func main() {
	host := flag.String("host", "", "host")
	port := flag.String("port", "7890", "port")

	flag.Parse()

	l, err := net.Listen("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Listening on " + *host + ":" + *port)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}

		//logs an incoming message
		fmt.Printf("%s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}


func handleRequest(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		reqLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}


		msg_recv := string(" Message received ")
		msg_recv += string(buf[:reqLen-1])
		fmt.Println(msg_recv)

		msg_reply, _ := createTestPb()

		_, err = conn.Write([]byte(msg_reply))
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}


	}
}


func createTestPb() ([]byte, error) {
	// 创建一个消息 Test
	test := &testpbgo.Test{
		// 使用辅助函数设置域的值
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Optionalgroup: &testpbgo.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}

	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("marshaling error: ", err)
	}

	return data, err
}