package store

import "sync"

type Car struct {
	ID int `json:"id"`
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
	Year int `json:"year"`
}

type Store struct {
	mu sync.RWMutex
	Cars map[string]Car
	NextID int
}

func NewStore() *Store {
	return &Store{
		Cars: make(map[string]Car),
		NextID: 1,
	}
}

func (s *Store) GetAll() []Car {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]Car, 0, len(s.Cars))
	for _, el := range s.Cars {
		list = append(list, el)
	}

	return list
}

func (s *Store) GetByName(name string) (Car, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	car, ok := s.Cars[name]
	return car, ok
}

func (s *Store) Add(car Car) Car {
	s.mu.Lock()
	defer s.mu.Unlock()

	car.ID = s.NextID
	s.NextID++

	s.Cars[car.Name] = car
	return car
}