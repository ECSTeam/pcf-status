package helpers

import (
	"log"
	"time"
)

// Cache defines a collection of items with an expiration.
type Cache interface {

	// Clear the cache of stale items.
	Clear()

	// Add an item to the cache.
	Add(interface{}, string)

	// Get the item from the cache
	Get(string) (interface{}, bool)
}

// cacheItem defines the items in a cache.
type cacheItem struct {
	expiration time.Time
	item       interface{}
}

// cache is the implementation of a collection of items with an expiration.
type cache struct {
	items    map[string]cacheItem
	duration time.Duration
}

// Clear the cache of stale items
func (c *cache) Clear() {
	remove := []string{}
	now := time.Now()
	for name, item := range c.items {
		if item.expiration.Before(now) {
			remove = append(remove, name)
		}
	}

	for _, name := range remove {
		log.Printf("Removing stale manifest: %s", name)
		delete(c.items, name)
	}
}

// Add an item to the cache
func (c *cache) Add(item interface{}, label string) {
	if _, has := c.items[label]; !has {
		log.Printf("Add cache item: %s", label)
		c.items[label] = cacheItem{
			expiration: time.Now().Add(c.duration),
			item:       item,
		}
	}
}

// Get the item by name.
func (c *cache) Get(label string) (data interface{}, has bool) {
	var item cacheItem
	if item, has = c.items[label]; has {
		log.Printf("Get cache item: %s", label)
		data = item.item
	}
	return data, has
}

// NewCache create the cache
func NewCache(duration time.Duration) (c Cache, err error) {
	c = &cache{
		duration: duration,
		items:    map[string]cacheItem{},
	}

	return c, err
}
