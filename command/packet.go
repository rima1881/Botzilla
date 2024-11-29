package command

import "encoding/json"

type packet struct {
	Follower string `json:"follower"`
	Body     string `json:"body"`
}

// Serialize serializes the Packet to JSON.
func (p packet) Serialize() ([]byte, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Deserialize deserializes JSON data into a Packet.
func Deserialize(data []byte) (packet, error) {
	var p packet
	err := json.Unmarshal(data, &p)
	if err != nil {
		return packet{}, err
	}
	return p, nil
}
