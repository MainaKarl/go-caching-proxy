package services

import "go-caching-proxy/pkg/models"

// clearCacheData clears the in-memory cache
func ClearCacheData() {
	models.CacheMutex.Lock()
	defer models.CacheMutex.Unlock()
	models.Cache = make(map[string][]byte)
}
