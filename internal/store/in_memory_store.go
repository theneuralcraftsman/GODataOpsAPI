package store

import (
	"sync"
	"GoDataOpsAPI/internal/model" // Update this line
	
)

type InMemoryStore struct {
	items  map[int]model.Item
	nextID int
	mu     sync.Mutex
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		items:  make(map[int]model.Item),
		nextID: 1,
	}
}

func (s *InMemoryStore) Create(item model.Item) model.Item {
	s.mu.Lock()
	defer s.mu.Unlock()

	item.ID = s.nextID
	s.nextID++
	s.items[item.ID] = item

	return item
}

func (s *InMemoryStore) GetAll() []model.Item {
	s.mu.Lock()
	defer s.mu.Unlock()

	itemList := make([]model.Item, 0, len(s.items))
	for _, item := range s.items {
		itemList = append(itemList, item)
	}
	return itemList
}

func (s *InMemoryStore) GetByID(id int) (model.Item, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.items[id]
	return item, exists
}

func (s *InMemoryStore) Update(id int, item model.Item) (model.Item, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.items[id]; !exists {
		return model.Item{}, false
	}

	item.ID = id
	s.items[id] = item
	return item, true
}

func (s *InMemoryStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.items[id]; !exists {
		return false
	}

	delete(s.items, id)
	return true
}
