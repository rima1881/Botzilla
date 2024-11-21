package main

import (
	"botzilla/command"
	"botzilla/core"
	"botzilla/message"
	"botzilla/stream"
	"fmt"
)

func main() {

	fmt.Println("Starting the Network...")

	go core.StartListener(7794, message.Handler)

	go core.StartListener(9112, stream.Handler)

	core.StartListener(6985, command.Handler)

}
