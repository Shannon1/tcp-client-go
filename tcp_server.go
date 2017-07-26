package main

import (
	"flag"
	"fmt"
	"net"
	"os"
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


		msg_reply := string(" Message received ")
		msg_reply += string(buf[:reqLen-1])
		fmt.Println("msg reply to client ", msg_reply)

		_, err = conn.Write([]byte(msg_reply))
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}


	}
}