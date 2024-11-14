package discovery

import (
	"strconv"
	"sync"
)

var ID_COUNTER int = 1

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

func (r *Registery) Register(address string) {
	newComponent := Component{ID: strconv.Itoa(ID_COUNTER), Group: []string{}, Address: address}
	r.components[newComponent.ID] = newComponent
	ID_COUNTER += 1
}

func (r *Registery) GetComponents() map[string]Component {
	return r.components
}
