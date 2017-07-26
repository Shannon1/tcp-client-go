package main

import (
	"os"
	"strconv"
	"net"
	"fmt"
	"time"
	"github.com/golang/protobuf/proto"
	"testpbgo"
)

const (
	REMOTE_ADDR = "127.0.0.1:7890"
)

func main() {

	// connect to this socket
	conn, err := net.Dial("tcp", REMOTE_ADDR)
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connecting to " + REMOTE_ADDR)
	done := make(chan string)

	go handleWrite(conn, done)
	go handleRead(conn, done)

	fmt.Println(<-done)
	fmt.Println(<-done)

	time.Sleep(time.Second * 20)
}

func handleWrite(conn net.Conn, done chan string) {
	for i := 1; i > 0; i-- {
		fmt.Println("count ", strconv.Itoa(i))
		_, e := conn.Write([]byte("hello " + strconv.Itoa(i) + "\r\n"))
		if e != nil {
			fmt.Println("Error to send message because of ", e.Error())
			break
		}
		time.Sleep(time.Second)
	}
	done <- "Sent"
}


func handleRead(conn net.Conn, done chan string) {
	for {
		buf := make([]byte, 1024)
		reqLen, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error to read message because of ", err)
			return
		}
		fmt.Println("recv len: ", reqLen)
		decode(buf[:reqLen-1])
		done <- "Receive"
	}
}


func decode(data []byte) {
	// 进行解码
	newTest := &testpbgo.Test{}
	err := proto.Unmarshal(data, newTest)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
	}

	fmt.Println("label: ", newTest.GetLabel())
	fmt.Println("type: ", newTest.GetType())
	for _, one := range newTest.GetReps() {
		fmt.Println("reps: ", one)
	}

	fmt.Println("Optionalgroup: ", newTest.GetOptionalgroup().RequiredField)

}