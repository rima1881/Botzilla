package message

type Message struct {
	Header []string
	Body   string
}

func (m Message) String() string {
	return m.Body
}

/*
	For now message headers only include
	the device ids but in future device
	grouping will be implemented
*/
