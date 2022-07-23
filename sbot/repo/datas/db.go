package datas

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	db *leveldb.DB
)

type Db struct {
	db *leveldb.DB
}

func CreateDb(path string) *Db {
	db := new(Db)
	var err error
	db.db, err = leveldb.OpenFile(path, nil)
	if err != nil {
		panic(err)
	}
	return db
}
func (db *Db) Get(key string) (bool, string) {

	value, err := db.db.Get([]byte(key), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			fmt.Printf("key: %v not found\n", key)
			return false, ""
		} else {
			panic(err)
		}
	} else {
		return true, string(value)
	}
}
func (db *Db) IterateAll(fn func(key string, value string)) {
	iter := db.db.NewIterator(nil, nil)
	for iter.Next() {
		fn(string(iter.Key()), string(iter.Value()))
	}
	iter.Release()
}
func (db *Db) Put(key string, value string) {
	fmt.Println("put:",key,value)
	err := db.db.Put([]byte(key), []byte(value), nil)
    if err != nil{
		fmt.Println("put error",err)
	}
}
func (db *Db) Delete(key string) {
	db.db.Delete([]byte(key), nil)
}
func (db *Db) Has(key string) bool {
	ok, err := db.db.Has([]byte(key), nil)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return ok
	}
}
func (db *Db) Close() {
	db.db.Close()
}
