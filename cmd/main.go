package main

import (
	"flag"
	"fmt"
	"go-caching-proxy/pkg/models"
	"go-caching-proxy/pkg/services"
	"log"
	"net/http"
)

func main() {
	// Parse command-line arguments
	flag.IntVar(&models.Port, "port", 3000, "Port on which the caching proxy server will run")
	flag.StringVar(&models.Origin, "origin", "", "URL of the server to which requests will be forwarded")
	flag.BoolVar(&models.ClearCache, "clear-cache", false, "Clear the cache")
	flag.Parse()

	// Handle cache clearing and exit
	if models.ClearCache {
		services.ClearCacheData()
		fmt.Println("Cache cleared successfully.")
		return
	}

	if models.Origin == "" {
		log.Fatal("The --origin flag is required to specify the origin server.")
	}

	// Start the caching proxy server
	http.HandleFunc("/", services.HandleRequest)
	log.Printf("Starting caching proxy server on port %d, forwarding to %s", models.Port, models.Origin)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", models.Port), nil))
}
