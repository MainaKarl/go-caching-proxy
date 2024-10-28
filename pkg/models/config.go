package models

import (
	"sync"
)

var (
	Cache      = make(map[string][]byte) // in-memory cache
	CacheMutex sync.RWMutex              // mutex to synchronize cache access
	Port       int
	Origin     string
	ClearCache bool
)
