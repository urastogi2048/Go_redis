package main

import (
	"sync"
	"time"
	//"strconv"
)

type Entry struct {
	value     string
	expiresIn time.Time
}
type Database struct {
	db map[string]Entry
	mu sync.RWMutex
	//func Set(db) {}

}

// os reference yooo
func (d *Database) Set(key string, value string) {

	d.mu.Lock()         //wait
	defer d.mu.Unlock() //execute
	entry := Entry{
		value: value,
	}
	d.db[key] = entry
}
func (d *Database) Get(key string) (string, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if val, ok := d.db[key]; ok {
		if val.expiresIn.IsZero() {

			return val.value, true
		} else if time.Now().Before(val.expiresIn) {

			return val.value, true
		} else if time.Now().After(val.expiresIn) {
			delete(d.db, key)
			return "", false
		} else {
			return "", false
		}

	}
	return "", false
}
func (d *Database) Del(key string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, ok := d.db[key]; ok {
		delete(d.db, key)
		return true
	} else {
		return false
	}
}

func (d *Database) StartActiveExpiration(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for range ticker.C {
			now := time.Now()
			d.mu.Lock()
			for key, entry := range d.db {
				if !entry.expiresIn.IsZero() &&
					!now.Before(entry.expiresIn) {
					delete(d.db, key)
				}
			}
			d.mu.Unlock()
		}
	}()
}
func (d *Database) SetWithTTL(key string, value string, interval time.Time) {
	d.mu.Lock()
	defer d.mu.Unlock()
	entry := Entry{
		value:     value,
		expiresIn: interval,
	}
	d.db[key] = entry
}
func (d *Database) TTLApply(key string, interval time.Time) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	entry, ok := d.db[key]
	if !ok {
		return false
	}

	entry.expiresIn = interval
	d.db[key] = entry
	return true
}
func (d *Database) TTLGet(key string) int{
	d.mu.Lock()
	defer d.mu.Unlock()
	entry, ok:=d.db[key]
	if !ok {
		return -1
	}else {
		remaining := time.Until(entry.expiresIn)
		seconds := int(remaining.Seconds())
		
		return seconds
	}
}
