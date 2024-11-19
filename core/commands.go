package core

import (
	"fmt"
	"log"
	"strings"
)

var HandlerMap = map[string]func(data string) (string, error){
	"0000": registerComponent,
	"0001": getComponents,
	"0002": createGroup,
	"0003": assignGroup,
	"0004": removeGroup,
}

func registerComponent(data string) (string, error) {
	// Log the action
	log.Println("Handling register command")

	registery := GetRegistery()

	// Check if message content was valid
	splitedData := strings.Split(data, ",")
	if len(splitedData) != 2 {
		return "", fmt.Errorf("invalid data format: expected 2 components, got %d", len(splitedData))
	}

	// Check if the component already exists
	isThere := registery.ComponentExists(splitedData[0])
	if isThere {
		return "", fmt.Errorf("component with name '%s' already exists", splitedData[0])
	}

	// Create a new component using the data
	comp := NewComponent(splitedData[0], splitedData[1])
	registery.AddComponent(comp)

	return "Device registered successfully", nil
}

func getComponents(_ string) (string, error) {
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
