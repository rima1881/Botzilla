package command

import (
	"fmt"
)

type ConnectHandler struct{}

func (h *ConnectHandler) Handle(data []byte) error {
	// Process connect command
	fmt.Println("Handling connect command")
	return nil
}

type GetComponentsHandler struct{}

func (h *GetComponentsHandler) Handle(data []byte) error {
	fmt.Println("Returning all registered Componets")

	return nil
}

type CreateGroupHandler struct{}

func (h *CreateGroupHandler) Handle(data []byte) error {
	fmt.Println("Creating Group")

	return nil
}

type AssignGroupHander struct{}

func (h *AssignGroupHander) Handle(data []byte) error {
	fmt.Println("Assigning to the group")

	return nil
}

type RemoveGroupHandler struct{}

func (h *RemoveGroupHandler) Handle(data []byte) error {
	fmt.Println("Removing from the group")

	return nil
}

type InvalidCommandHandler struct{}

func (h *InvalidCommandHandler) Handle(data []byte) error {
	fmt.Println("Recieved an invalid command")

	return nil
}
