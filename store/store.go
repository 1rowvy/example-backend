package store

import (
	"sort"
	"sync"
)

type Car struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Image string `json:"image"`
	Year int `json:"year"`
	City string `json:"city"`
	Price int `json:"price"`
}

type Store struct {
	mu sync.RWMutex
	Cars map[int]Car
	NextID int
}

func NewStore() *Store {
	return &Store{
		Cars: map[int]Car{
			1:{
				ID: 1,
				Name: "Geely Coolray",
				Image: "https://s12.auto.drom.ru/photo/v2/TdKIktC2ZAUHrFi16urV9UGGjg_xn4nVrhdB9lVQFJxfdBOQel5IOIrx8KgvBo0ufrJlELRatUNZHYng/gen1200.jpg", 
				Year: 2021, 
				City: "Иркутск", 
				Price: 1520000,
			},
			2:{
				ID: 2,
				Name: "Toyota C-HR",
				Image: "https://s12.auto.drom.ru/photo/v2/0Yve6nQobD5oIZZ-9BVpKzVfoExMOciS7rfEBuj0QQ8tu7vZHzww-6NtLuoO_v64uKqB5shko6424IbT/gen1200.jpg", 
				Year: 2019, 
				City: "Иркутск", 
				Price: 1850000,
			},
			3:{
				ID: 3,
				Name: "Hyundai Palisade",
				Image: "https://s12.auto.drom.ru/photo/v2/XyjBnyHcv0ZVgbeTHzS8WpyZeC24oTG2_uKeDzwP73z65cEf8-zW4lkZOySsVj96E8nLBKhgrIpr-4_L/gen1200.jpg", 
				Year: 2018, 
				City: "Иркутск", 
				Price: 3290000,
			},
		},
		NextID: 4,
	}
}

func (s *Store) GetAll() []Car {
	s.mu.RLock()
	defer s.mu.RUnlock()

	keys := make([]int, 0, len(s.Cars))
	for id := range s.Cars {
		keys = append(keys, id)
	}

	sort.Ints(keys)

	list := make([]Car, 0, len(s.Cars))
	for _, id := range keys {
		list = append(list, s.Cars[id])
	}

	return list
}

func (s *Store) GetById(id int) (Car, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	car, ok := s.Cars[id]
	return car, ok
}

func (s *Store) Add(car Car) Car {
	s.mu.Lock()
	defer s.mu.Unlock()

	car.ID = s.NextID
	s.NextID++

	s.Cars[car.ID] = car
	return car
}