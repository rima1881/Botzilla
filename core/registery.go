package core

import "sync"

type Registery struct {
	components map[string]Component
	mu         sync.RWMutex // Read-write mutex for thread-safety}
}

var (
	registery *Registery
	once      sync.Once
)

func GetRegistery() *Registery {

	once.Do(func() {
		registery = &Registery{components: make(map[string]Component)}
	})

	return registery
}

func (r *Registery) GetComponent(name string) Component {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.components[name]
}

func (r *Registery) AddComponent(c Component) {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.components[c.name] = c

}
