package main
type Database struct{
	db map[string] string 
	//func Set(db) {} 
	
}

func (d Database) Set(key string, value string){
	d.db[key]=value
}
func (d Database) Get(key string) (string,bool){
	if val, ok := d.db[key]; ok {
		return val, true
	}else{
		return "", false
	}
}
func (d Database) Del (key string) (bool){
	if _,ok:=d.db[key]; ok {
		delete(d.db, key)
		return true
	}else {
		return false
	}
}