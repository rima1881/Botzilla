package main

import (
	"botzilla/pkg/communication"
	"fmt"
)

func main() {

	fmt.Println("Starting the Network...")

	go communication.StartTCPServer()

	communication.StartWebsocketServer()

}
