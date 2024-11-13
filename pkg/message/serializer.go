package message

import "encoding/json"

func (m Message) Serialize() ([]byte, error) {

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Deserialize(rawData []byte) (Message, error) {

	var message Message
	err := json.Unmarshal(rawData, &message)

	if err != nil {
		return Message{}, err
	}

	return message, nil

}
