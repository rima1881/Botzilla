package message

type selectorType int

const (
	IDSelector    = iota
	classSelector = iota
)

func (s selectorType) Index() int {
	return int(s)
}

type messageHeader struct {
	selectorType string
	selector     string
}

type Message struct {
	header []messageHeader
	body   string
}
