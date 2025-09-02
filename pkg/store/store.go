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

// Function for Get
func (s *Store) Get(key string) ([]byte, bool) {
	if key == "" {
		return nil, false
	}
	s.mutex.RLock()
	v, ok := s.data[key]
	s.mutex.RUnlock()
	if !ok {
		return nil, false
	}
	return bytes.Clone(v), true
}

// Function for Set
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

// Function for Del
func (s *Store) Del(key string) bool {
	if key == "" {
		return false
	}
	s.mutex.Lock()
	_, existed := s.data[key]
	if existed {
		delete(s.data, key)
	}
	s.mutex.Unlock()
	return existed
}

// Function for Exists
func (s *Store) Exists(key string) bool {
	if key == "" {
		return false
	}
	s.mutex.RLock()
	_, ok := s.data[key]
	s.mutex.RUnlock()
	return ok
}

func (s *Store) Len() int {
	s.mutex.RLock()
	length := len(s.data)
	s.mutex.RUnlock()
	return length
}
