package command

type packet struct {
	follower string
	body     string
}

func (p packet) Serialize() ([]byte, error) {

}

func Deserialize([]byte) (packet, error) {

}
