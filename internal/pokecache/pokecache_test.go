package pokecache

import "testing"

func TestCreateCache(t *testing.T) {
	cache := NewCache()

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

	cache := NewCache()

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
