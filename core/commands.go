package core

import (
	"fmt"
)

func GetHandler(cmd string) func(data string) (string, error) {

	switch cmd {
	case "0000":
		return connect
	case "0001":
		return getComponents
	case "0002":
		return createGroup
	case "0003":
		return assignGroup
	case "0004":
		return removeGroup
	}

	return nil
}

func connect(data string) (string, error) {
	// Process connect command
	fmt.Println("Handling connect command")
	return "Device register", nil
}

func getComponents(data string) (string, error) {
	fmt.Println("Returning all registered Componets")

	return "device list:", nil
}

func createGroup(data string) (string, error) {
	fmt.Println("Creating Group")

	return "group created:", nil
}

func assignGroup(data string) (string, error) {
	fmt.Println("Assigning to the group")

	return "assigned to group", nil
}

func removeGroup(data string) (string, error) {
	fmt.Println("Removing from the group")

	return "group removed", nil
}
