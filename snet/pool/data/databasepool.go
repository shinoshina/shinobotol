package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConnection struct {
	db    *gorm.DB
	state bool
}


var Dbpool []DbConnection
var currentIndex int16
var idleDbNum int16
var defaultDbNum int16

func InitDbPool(num int16) {

	defaultDbNum = num
	currentIndex = 0
	idleDbNum = num

	Dbpool = make([]DbConnection, num)
	for i := range Dbpool {
		Dbpool[i].db, _ = gorm.Open(mysql.Open(Dsn), &gorm.Config{})
		Dbpool[i].state = false
	}

}

func GetDb() (*gorm.DB, int16) {

	if idleDbNum > 0 {
		defer updateState()

		for ; currentIndex < defaultDbNum && Dbpool[currentIndex].state; currentIndex++ {
		}
		return Dbpool[currentIndex].db, currentIndex
	} else {
		return nil, -1
	}
}

func updateState() {
	Dbpool[currentIndex].state = true
	idleDbNum--
	currentIndex++
	if currentIndex == defaultDbNum {
		currentIndex = 0
	}
}

func FinishTask(index int16) {
	Dbpool[index].state = false
	idleDbNum++
}
