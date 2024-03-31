package pokecache

import "testing"

func TestCreateCache(t *testing.T) {
	// create new cache.
	cache := NewCache()

	if cache.cache == nil {
		t.Errorf("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache()

	// Add to the cache
	cache.Add("a", []byte("test"))

	// test if value in cache.
	actual, ok := cache.Get("a")
	if !ok {
		t.Errorf("a not found!")
	}

	if string(actual) != "test" {
		t.Errorf("Wanted %s, got %s", []byte("test"), actual)
	}

}
