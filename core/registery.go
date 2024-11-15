package core

import "sync"

type Registery struct {
	components map[string]Component
	mu         sync.Mutex
}

func NewRegistery() *Registery {
	return &Registery{components: make(map[string]Component)}
}
