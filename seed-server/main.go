package main

import (
	"fmt"
	"log"
	"net"

	"github.com/doctor112-1/decentralized-domain-name-system-go/utils"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:7001")
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		deserializedMessage := utils.DeserializeMessage(buffer[:n])
		deserializedGetNodes := utils.DeserializeGetNodes([]byte(deserializedMessage.Payload))
		remoteAddress := conn.RemoteAddr().(*net.TCPAddr)
		remoteAddressIP := remoteAddress.IP.String()
		if deserializedGetNodes.GetNodes == 1 {
			fmt.Println(remoteAddressIP)
		}
	}
}
