package singleton

import (
	"errors"
	"fmt"
	"sync"
)

var (
	instance     *singleton
	once         sync.Once
	errBadObject = errors.New("Bad object")
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

func (s *singleton) Title() (string, error) {
	fmt.Printf("Title(): Addr s: %p and addr instance: %p\n", s, instance)
	if !s.IsCorrectObject() {
		fmt.Println("This is a bad object. Use another one!")

		return "Bad object", errBadObject
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.title, nil
}

func (s *singleton) SetNewTitle(newName string) {
	fmt.Printf("SetNewTitle(): Addr s: %p and addr instance: %p\n", s, instance)
	if !s.IsCorrectObject() {
		fmt.Println("This is a bad object. Use another one!")

		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	s.title = newName
}

func (s *singleton) IsCorrectObject() bool {
	return s == instance
}
