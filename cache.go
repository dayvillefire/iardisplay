package main

import (
	"log"
	"sync"
	"time"

	"github.com/jbuchbinder/cadmonitor/monitor"
)

type CadCallStatusCache struct {
	Monitor        monitor.CadMonitor
	ExpiryDuration time.Duration

	cache       map[string]CadCallStatusCacheItem
	lock        sync.Mutex
	initialized bool
}

func (c *CadCallStatusCache) RetrieveWithCache(id string) (monitor.CallStatus, error) {
	if c.Cached(id) {
		log.Printf("CadCallStatusCache: Found %s in cache", id)
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.cache[id].Item, nil
	}

	log.Printf("CadCallStatusCache: Caching %s", id)
	item, err := c.Monitor.GetStatus(id)
	if err != nil {
		return item, err
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	if !c.initialized {
		c.cache = map[string]CadCallStatusCacheItem{}
		c.initialized = true
	}
	c.cache[id] = CadCallStatusCacheItem{
		Item:   item,
		Expiry: time.Now().Add(c.ExpiryDuration),
	}
	return item, nil
}

func (c *CadCallStatusCache) Cached(id string) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	if !c.initialized {
		c.cache = map[string]CadCallStatusCacheItem{}
		c.initialized = true
	}
	item, exists := c.cache[id]
	if !exists || item.Expired() {
		return false
	}
	return true
}

type CadCallStatusCacheItem struct {
	Expiry time.Time
	Item   monitor.CallStatus
}

func (c CadCallStatusCacheItem) Expired() bool {
	return c.Expiry.Before(time.Now())
}
