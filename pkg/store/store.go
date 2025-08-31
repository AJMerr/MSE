package store

import (
	"bytes"
	"errors"
	"sync"
)

// Shape for the KV store
type Store struct {
	mutex sync.RWMutex
	data  map[string][]byte
}

// Constructor for Store
func NewStore() *Store {
	return &Store{data: make(map[string][]byte)}
}

var ErrEmptyKey = errors.New("empty key")

// TODO(Austin): implement
func (s *Store) Get(key string) ([]byte, bool) {
	return nil, false
}

// TODO(Austin): implement
func (s *Store) Set(key string, value []byte) error {
	if key == "" {
		return ErrEmptyKey
	}
	clone := bytes.Clone(value)
	s.mutex.Lock()
	s.data[key] = clone
	s.mutex.Unlock()
	return nil
}

// TODO(Austin): implement
func (s *Store) Del(key string) bool {
	return false
}

// TODO(Austin): implement
func (s *Store) Exists(key string) bool {
	return false
}
