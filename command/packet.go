package command

import "encoding/json"

type packet struct {
	follower string
	body     string
}

// TODO
// Needs to be changed
func (p packet) Serialize() ([]byte, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func Deserialize(data []byte) (packet, error) {
	var p packet
	err := json.Unmarshal(data, &p)

	if err != nil {
		return packet{}, err
	}

	return p, nil
}
