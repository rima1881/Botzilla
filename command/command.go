package command

import (
	"bufio"
	"fmt"
	"net"
)

func Handler(conn net.Conn) {

	defer conn.Close()

	// Create a buffered reader
	reader := bufio.NewReader(conn)

	// Read the entire message (this will read until it finds a newline or EOF)
	rawMessage, err := reader.ReadBytes('\n')

	if err != nil {
		fmt.Printf("Failed to read message: %v\n", err)
		return
	}

	request, err := Deserialize(rawMessage)
	if err != nil {
		fmt.Println("go fuck yourself")
	}
	response, err := route(request)
	if err != nil {
		fmt.Println("read previos error handling")
	}

	conn.Write([]byte(response))

}
