package cache

import "time"

type element struct {
	value    string
	deadline time.Time
}

func (receiver element) IsExpired() bool {
	if receiver.deadline.IsZero() {
		return false
	}

	return time.Now().Sub(receiver.deadline).Seconds() > 0
}

type Cache struct {
	key map[string]element
}

func (receiver *Cache) cleanCache() {
	for k, v := range receiver.key {
		if v.IsExpired() {
			delete(receiver.key, k)
		}
	}
}

func NewCache() Cache {
	elements := make(map[string]element)
	return Cache{key: elements}
}

func (receiver Cache) Get(key string) (string, bool) {
	e, ok := receiver.key[key]

	if !ok {
		return "", false
	}

	if e.IsExpired() {
		delete(receiver.key, key)
		return "", false
	}

	return e.value, ok
}

func (receiver *Cache) Put(key, value string) {
	element := element{value: value}
	receiver.key[key] = element

	receiver.cleanCache()
}

func (receiver Cache) Keys() []string {
	receiver.cleanCache()
	keys := []string{}

	for k := range receiver.key {
		keys = append(keys, k)
	}

	return keys
}

func (receiver *Cache) PutTill(key, value string, deadline time.Time) {
	element := element{value: value, deadline: deadline}
	receiver.key[key] = element

	receiver.cleanCache()
}
