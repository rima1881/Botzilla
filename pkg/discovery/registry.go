package discovery

import (
	"sync"
)

// Each component can be a member of multiple groups
type Component struct {
	ID      string
	Group   []string
	Address string
	// Might Add last seen :O
}

type Registery struct {
	//Keys are the IDs
	components map[string]Component
	mu         sync.Mutex
}

func NewRegistery() *Registery {
	return &Registery{components: make(map[string]Component)}
}

func (r *Registery) Register(comp Component) {
	r.components[comp.ID] = comp
}

func (r *Registery) GetComponents() map[string]Component {
	return r.components
}
