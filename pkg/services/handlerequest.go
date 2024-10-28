package services

import (
	"go-caching-proxy/pkg/models"
	"io"
	"net/http"
)

// handleRequest forwards the request to the origin server and caches the response
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	url := models.Origin + r.URL.Path

	// Check cache for existing response
	models.CacheMutex.RLock()
	cachedResponse, found := models.Cache[url]
	models.CacheMutex.RUnlock()

	if found {
		// Return cached response with HIT header
		w.Header().Set("X-Cache", "HIT")
		w.Write(cachedResponse)
		return
	}

	// Forward request to the origin server if not in cache
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error forwarding request to origin", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading origin server response", http.StatusInternalServerError)
		return
	}

	// Cache the response
	models.CacheMutex.Lock()
	models.Cache[url] = body
	models.CacheMutex.Unlock()

	// Add MISS header and write response
	w.Header().Set("X-Cache", "MISS")
	w.Write(body)
}
