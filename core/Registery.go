package core

import "sync"

type Registery struct {
	components map[string]*Component
	mu         sync.RWMutex
}

var (
	registery *Registery
	once      sync.Once
)

func GetRegistery() *Registery {
	once.Do(func() {
		registery = &Registery{components: make(map[string]*Component)}
	})
	return registery
}

func (r *Registery) GetComponent(name string) *Component {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return registery.components[name]
}

func (r *Registery) AddComponent(name string, address string) (string, error) {

	newComp := newComponent(name, address)
	r.mu.Lock()
	defer r.mu.Unlock()

	r.components[name] = newComp

	return newComp.GetToken(), nil
}
