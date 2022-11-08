package cache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {

	cache := New(30*time.Minute, 1*time.Hour)
	value, found := cache.Get("1")
	if found || value != "" {
		t.Error("Empty cache should not return value")
	}
	cache.Set("1", "data")
	value, found = cache.Get("1")
	if !found {
		t.Error("No value was found")
	} else if value != "data" {
		t.Error("Should return correct value", value)
	}
	cache.Set("1", "data2")
	value, found = cache.Get("1")
	if !found {
		t.Error("No value was found")
	} else if value != "data2" {
		t.Error("Should return correct overwritten value", value)
	}
}

func TestCacheExpiration(t *testing.T) {
	cache := New(1*time.Second, 1*time.Hour)
	cache.Set("1", "data")
	value, found := cache.Get("1")
	if !found {
		t.Error("No value was found")
	} else if value != "data" {
		t.Error("Should return correct value", value)
	}
	time.Sleep(1 * time.Second)
	value, found = cache.Get("1")
	if found || value != "" {
		t.Error("Expired cache item should not return value")
	}
}

func TestCacheCleanup(t *testing.T) {
	cache := New(1*time.Second, 2*time.Second)
	cache.Set("1", "data")
	value, found := cache.Get("1")
	if !found {
		t.Error("No value was found")
	} else if value != "data" {
		t.Error("Should return correct value", value)
	}
	time.Sleep(3 * time.Second)
	cacheSize := len(cache.items)
	if cacheSize != 0 {
		t.Errorf("Cleaned cache should be empty, size %d", cacheSize)
	}
}
