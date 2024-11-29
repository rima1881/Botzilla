package command

import (
	"botzilla/core"
	"fmt"
	"net"
)

func route(p packet) (string, error) {

	fmt.Println(p)

	if p.Follower == "0000" {
		return sendToSelf(p)
	}

	return sendToFollower(p)

}

func sendToFollower(p packet) (string, error) {

	registery := core.GetRegistery()

	follower := registery.GetComponent(p.Follower)
	address := follower.TCPAddress

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	message, err := p.Serialize()
	if err != nil {
		return "", err
	}

	_, err = conn.Write(message)
	if err != nil {
		return "", err
	}

	// Read the response from the server
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer), nil

}

func sendToSelf(p packet) (string, error) {

	command := p.Body[:4]

	handler := core.HandlerMap[command]

	if handler == nil {
		return "wrong command id", nil
	}

	return handler(p.Body)

}
