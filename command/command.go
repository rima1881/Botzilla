package command

import (
	"fmt"
	"net"
	"os"
)

func Start(port int) {

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

		go requestHandler(conn)
	}

}

func requestHandler(conn net.Conn) {

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
