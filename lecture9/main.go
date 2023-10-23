package main

import (
	"errors"
	"sync"
)

type Entity struct {
	ID   int
	Name string
}

type Repository interface {
	GetAll() []Entity
	GetByID(id int) (*Entity, error)
	Create(entity Entity) (*Entity, error)
	Update(id int, entity Entity) (*Entity, error)
	Delete(id int) error
}

type InMemoryRepository struct {
	mu      sync.RWMutex
	storage map[int]Entity
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		storage: make(map[int]Entity),
	}
}

func (r *InMemoryRepository) GetAll() []Entity {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entities := make([]Entity, 0, len(r.storage))
	for _, entity := range r.storage {
		entities = append(entities, entity)
	}

	return entities
}

func (r *InMemoryRepository) GetByID(id int) (*Entity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	entity, exists := r.storage[id]
	if !exists {
		return nil, errors.New("Entity not found")
	}
	return &entity, nil
}

func (r *InMemoryRepository) Create(entity Entity) (*Entity, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	entity.ID = len(r.storage) + 1
	r.storage[entity.ID] = entity
	return &entity, nil
}

func (r *InMemoryRepository) Update(id int, entity Entity) (*Entity, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.storage[id]
	if !exists {
		return nil, errors.New("Entity not found")
	}

	entity.ID = id
	r.storage[id] = entity
	return &entity, nil
}

func (r *InMemoryRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.storage[id]
	if !exists {
		return errors.New("Entity not found")
	}

	delete(r.storage, id)
	return nil
}

func main() {
	repo := NewInMemoryRepository()

	entity1 := Entity{Name: "Entity 1"}
	createdEntity, err := repo.Create(entity1)
	if err != nil {
		panic(err)
	}

	retrievedEntity, err := repo.GetByID(createdEntity.ID)
	if err != nil {
		panic(err)
	}

	updatedEntity := *retrievedEntity
	updatedEntity.Name = "Updated Entity 1"
	_, err = repo.Update(updatedEntity.ID, updatedEntity)
	if err != nil {
		panic(err)
	}

	err = repo.Delete(retrievedEntity.ID)
	if err != nil {
		panic(err)
	}
}
