package command

import "botzilla/core"

func route(p packet) (string, error) {

	if p.follower == "0000" {
		return sendToSelf(p)
	}

	return sendToFollower(p)

}

func sendToFollower(p packet) (string, error) {

}

func sendToSelf(p packet) (string, error) {

	command := p.body[:4]

	handler := core.GetHandler(command)

	if handler == nil {
		return "wrong command id", nil
	}

	return handler(p.body)

}
