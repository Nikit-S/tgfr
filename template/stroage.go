package template

import "fmt"

type IStorage interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	Delete(key string) error
}

type Storage struct {
	db map[string][]byte
}

func NewStorage() *Storage {
	return &Storage{db: make(map[string][]byte)}
}

func (s *Storage) Get(key string) ([]byte, error) {
	if value, ok := s.db[key]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("key %s not found", key)
}

func (s *Storage) Set(key string, value []byte) error {
	s.db[key] = value
	return nil
}

func (s *Storage) Delete(key string) error {
	delete(s.db, key)
	return nil
}
