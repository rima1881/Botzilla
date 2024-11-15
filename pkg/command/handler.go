package command

import (
	"botzilla/pkg/message"
	"fmt"
)

type CommandHandler interface {
	Handle(data []byte) error
}

type CommandRouter struct {
	handlers map[string]CommandHandler
}

func NewCommandRouter() *CommandRouter {
	return &CommandRouter{
		handlers: map[string]CommandHandler{
			"0000": &ConnectHandler{},
			"0001": &GetComponentsHandler{},
			"0002": &CreateGroupHandler{},
			"0003": &RemoveGroupHandler{},
		},
	}
}

func RequestHander(data []byte) error {
	fmt.Println("Sending data to components...")

	return nil
}

func (r *CommandRouter) Route(msg *message.Message) {

	for _, clinet := range msg.Header {

		// If the message is not targeted at botzilla itself
		if clinet != "0000" {

			transmisionMessage := message.Message{Header: []string{clinet}, Body: msg.Body}
			encodedMSG, err := transmisionMessage.Serialize()
			if err != nil {
				fmt.Println("Error serializing message")
				continue
			}

			RequestHander(encodedMSG)

		}

		command := msg.Body[:4]
		hander := r.handlers[command]
		hander.Handle([]byte(msg.Body))

	}

}
