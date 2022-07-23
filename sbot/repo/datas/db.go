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

func (db *Db) Put(key string, value string) {
	db.db.Put([]byte(key), []byte(value), nil)
}
func (db *Db) Delete(key string) {
	db.db.Delete([]byte(key), nil)
}
func (db *Db) Close() {
	db.db.Close()
}
