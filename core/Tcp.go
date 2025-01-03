package core

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func StartTCPServer(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("There was an error starting the server: \n", err)
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("Commanding server has started on port 8080")

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: \n", err)
			continue
		}

		go handler(conn)
	}
}

func handler(conn net.Conn) {

	defer conn.Close()

	// Create a buffered reader
	reader := bufio.NewReader(conn)

	// Read the entire message (this will read until it finds a newline or EOF)
	rawMessage, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Error reading from connection: \n", err)
	}

	message := string(rawMessage)
	remoteAddr := conn.RemoteAddr().String()

	response, err := router(message, remoteAddr)

	conn.Write(response)
	conn.Write([]byte("\n"))

}
