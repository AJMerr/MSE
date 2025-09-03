# MSE - In Memory KV Store
A tiny, thread-safe, binary-safe key–value store written in Go.
Built as part of a multi-repo capstone (HTTP framework → KV store → reverse proxy → Todo app).

## Features
- Concurrent and safe: sync.RWMutex guards a Go map
- Binary-safe values: []byte values; no encoding assumptions
- Copy semantics:
- Set copies input (callers can’t mutate internal state)
- Get returns a copy (callers can mutate safely)
- Core operations: Set, Get, Del, Exists (+ Len)
- Race-detector clean: go test -race passes

## Installation
`go get github.com/AJMerr/MSE/pkg/store`

## Example Quick Start
```
package main

import (
	"fmt"
	"github.com/AJMerr/MSE/pkg/store"
)

func main() {
	s := store.NewStore()

	_ = s.Set("example", []byte("hello"))
	v, ok := s.Get("example")
	if ok {
		fmt.Println(string(v)) // "hello"
	}

	fmt.Println(s.Exists("example")) // true
	fmt.Println(s.Del("example"))    // true
	fmt.Println(s.Exists("example")) // false
}
```

## API
```
// constructor
func NewStore() *Store

// core ops
func (s *Store) Set(key string, value []byte) error   // returns ErrEmptyKey on ""
func (s *Store) Get(key string) ([]byte, bool)        // (nil,false) if missing or key == ""
func (s *Store) Del(key string) bool                  // false if missing or key == ""
func (s *Store) Exists(key string) bool               // false if missing or key == ""
func (s *Store) Len() int                             // snapshot count of keys

// sentinel error
var ErrEmptyKey error
```
