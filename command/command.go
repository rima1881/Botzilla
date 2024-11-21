package command

import (
	"fmt"
	"net"
)

func Handler(conn net.Conn) {

	defer conn.Close()

	// Might need to change :O
	rawMessage := make([]byte, 1024)
	_, err := conn.Read(rawMessage)
	if err != nil {
		fmt.Println("Error reading command:\n", err)
		conn.Close()
		return
	}

	request, err := Deserialize(rawMessage)
	response, err := route(request)

	conn.Write([]byte(response))

}
