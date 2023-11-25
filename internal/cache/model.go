package cache

import "sync"

type Cache struct {
	mu        sync.Mutex
	size      int
	Requests  []string
	Responses map[string][]byte
}
