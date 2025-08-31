package store

import "sync"

// Shape for the KV store
type Store struct {
	mu   sync.RWMutex
	data map[string][]byte
}

// Constructor for Store
func NewStore() *Store {
	return &Store{data: make(map[string][]byte)}
}
