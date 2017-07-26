package main

//import "bufio"
import (
	"os"
	"strconv"
	"net"
	"fmt"
	"time"
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
	for i := 10; i > 0; i-- {
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
		fmt.Println(string(buf[:reqLen-1]))
		done <- "Receive"
	}
}