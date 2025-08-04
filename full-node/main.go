package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func firstStart(dir string) {
	dnsSeed := "127.0.0.1:7001"
	err := os.Mkdir(dir, 0755)
	if err != nil {
		log.Panic(err)
	}

	listener, err := net.Listen("tcp", "localhost:7000")
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	conn, err := net.Dial("tcp", dnsSeed)
	if err != nil {
		log.Panic(err)
	}
	conn.Close()
}

func main() {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Panic(err)
	}

	dir = dir + "/.ddns"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		firstStart(dir)
	} else {
		fmt.Println("The provided directory named", dir, "exists")
	}
}
