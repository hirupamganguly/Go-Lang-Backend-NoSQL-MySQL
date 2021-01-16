package main

import (
	"bufio"
	"fmt"
	"net"
)

func handelTCPRequest(conn net.Conn) {
	mesz, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message received from client: ", string(mesz))
	message := "Data send from server"
	conn.Write([]byte(message + "\n"))
	conn.Close()
}
func main() {
	listener, _ := net.Listen("tcp", "localhost:8080")
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go handelTCPRequest(conn)

	}
}
