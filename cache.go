package main

import (
	"errors"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/jbuchbinder/iarapi"

	"github.com/jbuchbinder/cadmonitor/monitor"
)

const (
	IarDispatchMessage  = "DispatchMessage"
	IarIncidentInfoData = "IncidentInfoData"
	IarLatestIncidents  = "LatestIncidents"
	IarNowResponding    = "NowResponding"
	IarOnSchedule       = "OnSchedule"
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

type IarCache struct {
	Iar            iarapi.IamRespondingAPI
	ExpiryDuration time.Duration

	cache       map[string]IarCacheItem
	lock        sync.Mutex
	initialized bool
}

func (c *IarCache) RetrieveWithCache(objType, id string) (interface{}, error) {
	if c.Cached(objType, id) {
		log.Printf("IarCache: Found type '%s' entry '%s' in cache", objType, id)
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.cache[objType+"##"+id].Item, nil
	}

	log.Printf("IarCache: Caching type '%s' entry '%s'", objType, id)
	var item interface{}
	var err error
	switch objType {
	case IarDispatchMessage:
		item, err = c.Iar.ListWithParser()
		break
	case IarLatestIncidents:
		item, err = c.Iar.GetLatestIncidents()
		break
	case IarIncidentInfoData:
		intid, _ := strconv.ParseInt(id, 10, 64)
		item, err = c.Iar.GetIncidentInfo(intid)
		break
	case IarNowResponding:
		item, err = c.Iar.GetNowRespondingWithSort()
		break
	case IarOnSchedule:
		item, err = c.Iar.GetOnScheduleWithSort()
		break
	default:
		return nil, errors.New("invalid cache type")
	}
	if err != nil {
		return item, err
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	if !c.initialized {
		c.cache = map[string]IarCacheItem{}
		c.initialized = true
	}
	c.cache[objType+"##"+id] = IarCacheItem{
		Item:   item,
		Expiry: time.Now().Add(c.ExpiryDuration),
	}
	return item, nil
}

func (c *IarCache) Cached(objType, id string) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	if !c.initialized {
		c.cache = map[string]IarCacheItem{}
		c.initialized = true
	}
	item, exists := c.cache[objType+"##"+id]
	if !exists || item.Expired() {
		return false
	}
	return true
}

type IarCacheItem struct {
	Expiry time.Time
	Item   interface{}
}

func (c IarCacheItem) Expired() bool {
	return c.Expiry.Before(time.Now())
}
