package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	if cache.cache == nil {
		t.Errorf("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	testCases := []struct {
		inputKey   string
		inputvalue []byte
	}{
		{
			inputKey:   "animal",
			inputvalue: []byte("lion"),
		},
		{
			inputKey:   "name",
			inputvalue: []byte("john"),
		},
	}

	cache := NewCache(time.Millisecond)

	for _, cas := range testCases {
		// Add to the cache
		cache.Add(cas.inputKey, cas.inputvalue)

		// test if value in cache.
		actual, ok := cache.Get(cas.inputKey)

		if !ok {
			t.Errorf("%s not found!", cas.inputKey)
			continue
		}

		if string(actual) != string(cas.inputvalue) {
			t.Errorf("Wanted %s, got %s", cas.inputvalue, actual)
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10

	cache := NewCache(interval)

	key := "hey"
	cache.Add(key, []byte("hello"))

	time.Sleep(interval + time.Millisecond)
	_, ok := cache.Get(key)

	if ok {
		t.Errorf("%s should have been reaped!", key)
	}
}
