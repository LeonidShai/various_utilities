package singleton

import (
	"fmt"
	"sync"
)

var (
	instance *singleton
	once     sync.Once
)

type singleton struct {
	title string
	mu    *sync.RWMutex
}

func GetSingletonInstance(name string) *singleton {
	once.Do(func() {
		fmt.Println("Create Singleton")
		instance = new(singleton)
		instance.title = name
		instance.mu = new(sync.RWMutex)
	})

	return instance
}

func (s *singleton) Title() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.title
}

func (s *singleton) SetNewTitle(newName string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.title = newName
}
