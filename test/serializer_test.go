package test

import (
	"botzilla/pkg/message"
	"io"
	"os"
	"testing"
)

func TestDeserialize(t *testing.T) {

	jsonPath := "sample_message.json"
	expectedBody := "HEHE"
	expectedHeader := []string{"123", "546"}

	file, err := os.Open(jsonPath)

	if err != nil {
		t.Errorf("Failed opening %s:\n %v", jsonPath, err)
		return
	}

	defer file.Close()

	fileData, err := io.ReadAll(file)

	if err != nil {
		t.Errorf("Failed reading %s:\n %v", jsonPath, err)
		return
	}

	msg, err := message.Deserialize(fileData)

	if err != nil {
		t.Errorf("Failed deserialing:\n %v", err)
		return
	}

	if msg.Body != expectedBody {
		t.Errorf("Expected %s for body but received %s", expectedBody, msg.Body)
	}

	if !(msg.Header[0] == expectedHeader[0] && msg.Header[1] == expectedHeader[1] && len(msg.Header) == len(expectedHeader)) {
		t.Errorf("Expected %#v for head but received %#v", expectedHeader, msg.Header)
	}

}

func TestSerialize(t *testing.T) {

	jsonPath := "sample_message.json"

	file, err := os.Open(jsonPath)

	if err != nil {
		t.Errorf("Failed opening %s:\n %v", jsonPath, err)
		return
	}

	defer file.Close()

	fileData, err := io.ReadAll(file)

	if err != nil {
		t.Errorf("Failed reading %s:\n %v", jsonPath, err)
		return
	}

	msg := message.Message{Header: []string{"123", "546"}, Body: "HEHE"}

	decodedData, err := msg.Serialize()
	if err != nil {
		t.Errorf("Failed Serializing the message:\n %v", err)
		return
	}

	equal := len(decodedData) == len(fileData)

	i := 0
	for i < len(decodedData) && equal {
		equal = decodedData[i] == fileData[i]
		i += 1
	}

	if !equal {
		t.Errorf("Serialized data doesn't match Expected data")
	}

}
