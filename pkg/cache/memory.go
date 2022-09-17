package cache

import (
	"github.com/samber/lo"
	"sync"
)

type MemoryDriver struct {
	sync.Map
}

func (m *MemoryDriver) Get(key string) (interface{}, bool) {
	return m.Map.Load(key)
}

func (m *MemoryDriver) Gets(keys []string, prefix string) (map[string]interface{}, []string) {
	result := make(map[string]interface{})
	okList := make([]string, 0, len(keys))
	for _, key := range keys {
		if val, ok := m.Get(prefix + key); ok {
			result[key] = val
			okList = append(okList, key)
		}
	}
	d, _ := lo.Difference[string](keys, okList)
	return result, d
}

func (m *MemoryDriver) Set(key string, val interface{}, ttl int64) error {
	m.Map.Store(key, val)
	return nil
}

func (m *MemoryDriver) Sets(items map[string]interface{}, prefix string) error {
	for key, val := range items {
		_ = m.Set(prefix+key, val, 0)
	}
	return nil
}

func (m *MemoryDriver) Delete(key string) error {
	m.Map.Delete(key)
	return nil
}

func (m *MemoryDriver) Deletes(keys []string, prefix string) error {
	for _, key := range keys {
		_ = m.Delete(prefix + key)
	}
	return nil
}

func NewMemoryDriver() *MemoryDriver {
	return &MemoryDriver{}
}
