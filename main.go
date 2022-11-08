package main

import (
	"io"
	"key-value-cache/cache"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const CacheExpiration = 30 * time.Minute
const CacheCleanupInterval = 1 * time.Hour

var c *cache.Cache

func main() {
	// Log to file
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	// Init cache
	c = cache.New(CacheExpiration, CacheCleanupInterval)
	// Serve and handle http requests
	http.HandleFunc("/", processRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	key := strings.Split(r.URL.Path, "/")[1]
	switch r.Method {
	case "GET":
		log.Printf("Received GET request for key: %s", key)
		value, found := c.Get(key)
		if !found {
			w.WriteHeader(http.StatusNotFound)
		}
		_, err := w.Write([]byte(value))
		if err != nil {
			log.Printf("Failed to write response: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/plain")
		}
	case "POST":
		log.Printf("Received POST request for key: %s", key)
		value, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Bad request %s", err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			log.Printf("Setting cache entry for key: %s, value: %s", key, value)
			c.Set(key, string(value[:]))
			w.WriteHeader(http.StatusOK)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
