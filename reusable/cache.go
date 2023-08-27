package reusable

import (
	"time"
)

type Cache interface {
	Set(string, string, time.Duration) error
	Get(string) (string, error)
	Forget(string) error
}

type CacheMapData struct {
	value string
	ttl   time.Time
}

type CacheMap struct {
	values map[string]*CacheMapData
}

func NewCacheMap(values map[string]*CacheMapData) *CacheMap {
	return &CacheMap{values: values}
}

func (m CacheMap) Set(key string, value string, expiry time.Duration) error {
	ttl := time.Now().Add(expiry)
	m.values[key] = &CacheMapData{value: value, ttl: ttl}
	return nil
}

func (m CacheMap) Get(key string) (string, error) {
	if data, exists := m.values[key]; exists && data.ttl.After(time.Now()) {
		return data.value, nil
	}

	return "", nil
}

func (m CacheMap) Forget(key string) error {
	delete(m.values, key)
	return nil
}
