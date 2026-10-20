package main

import (
	"sync"
	"time"
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
		} else if time.Now().After(val.expiresIn){
			delete(d.db, key)
			return "", false
		}else {
			return "",false
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
