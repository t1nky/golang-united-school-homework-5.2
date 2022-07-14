package cache

import (
	"time"
)

type Value struct {
	expireAt time.Time
	value    string
}

type Cache struct {
	mem map[string]Value
}

func NewCache() Cache {
	return Cache{
		mem: make(map[string]Value),
	}
}

func (c Cache) Get(key string) (string, bool) {
	res, ok := c.mem[key]
	if !res.expireAt.IsZero() && !time.Now().Before(res.expireAt) {
		delete(c.mem, key)
		return "", false
	}
	return res.value, ok
}

func (c Cache) Put(key, value string) {
	c.mem[key] = Value{expireAt: time.Time{}, value: value}
}

func (c Cache) Keys() []string {
	keys := make([]string, len(c.mem))
	i := 0
	for k := range c.mem {
		keys[i] = k
		i++
	}

	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.mem[key] = Value{expireAt: deadline, value: value}
}
