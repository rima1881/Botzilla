package communication

import (
	"botzilla/pkg/message"
	"fmt"
	"net"
	"os"
)

func StartListener() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("There was an error starting the server: \n", err)
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("Botzilla has started on port 8080")

	for {

		connMaster, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: \n", err)
			continue
		}

		go handleMaster(connMaster)

	}

}

func handleMaster(conn net.Conn) {

	// Might need to change :O
	rawMessage := make([]byte, 1024)
	_, err := conn.Read(rawMessage)
	if err != nil {
		fmt.Println("Error reading command:\n", err)
		conn.Close()
		return
	}

	message := message.Deserialize(rawMessage)

	print(message)

}

func handleRequest(conn net.Conn) {

}
