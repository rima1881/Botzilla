package core

type Component struct {
	name             string
	TCPAddress       string
	websocketAddress string
	group            []string
}

func (c Component) sendCommnad(p)
