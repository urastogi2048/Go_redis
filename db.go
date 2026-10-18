package main
import (
	"sync"
)
type Database struct{
	db map[string] string 
	mu sync.RWMutex
	//func Set(db) {} 
	
}
//os reference yooo
func (d *Database) Set(key string, value string){
	d.mu.Lock()//wait
	defer d.mu.Unlock()//execute
	d.db[key]=value
}
func (d *Database) Get(key string) (string,bool){
	d.mu.RLock()
	defer d.mu.RUnlock()
	if val, ok := d.db[key]; ok {
		return val, true
	}else{
		return "", false
	}
}
func (d *Database) Del (key string) (bool){
	d.mu.Lock()
	defer d.mu.Unlock()
	if _,ok:=d.db[key]; ok {
		delete(d.db, key)
		return true
	}else {
		return false
	}
}