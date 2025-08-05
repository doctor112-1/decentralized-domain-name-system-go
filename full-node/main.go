package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"unsafe"

	"github.com/doctor112-1/decentralized-domain-name-system-go/utils"
)

func askForNodes() {
	dnsSeed := "127.0.0.1:7001"

	listener, err := net.Listen("tcp", "localhost:7000")
	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	conn, err := net.Dial("tcp", dnsSeed)
	if err != nil {
		log.Panic(err)
	}

	var askForNodesMessage utils.Message
	var getNodes utils.GetNodes

	getNodes.GetNodes = 1

	serializedGetNodes := getNodes.Serialize()

	askForNodesMessage.Version = 1
	askForNodesMessage.Command = "getNodes"
	askForNodesMessage.Payload = string(serializedGetNodes)
	askForNodesMessage.Length = uint32(unsafe.Sizeof(askForNodesMessage.Version)) + uint32(unsafe.Sizeof(askForNodesMessage.Command)) + uint32(unsafe.Sizeof(askForNodesMessage.Payload))

	SerializedAskForNodesMessage := askForNodesMessage.Serialize()
	_, err = conn.Write(SerializedAskForNodesMessage)
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
		err := os.Mkdir(dir, 0755)
		if err != nil {
			log.Panic(err)
		}
		askForNodes()
	} else {
		fmt.Println("The provided directory named", dir, "exists")
	}
}
