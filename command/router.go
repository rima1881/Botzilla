package command

import (
	"botzilla/core"
	"net"
)

func route(p packet) (string, error) {

	if p.follower == "0000" {
		return sendToSelf(p)
	}

	return sendToFollower(p)

}

func sendToFollower(p packet) (string, error) {

	registery := core.GetRegistery()

	follower := registery.GetComponent(p.follower)
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

	command := p.body[:4]

	handler := core.GetHandler(command)

	if handler == nil {
		return "wrong command id", nil
	}

	return handler(p.body)

}
