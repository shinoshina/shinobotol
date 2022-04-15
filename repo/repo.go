package repo

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var CurrentId int64
var CrucialId int32

type data struct {
	ID       int    `gorm:"primary_key;index:id"`
	Question string `gorm:"type:varchar(100);not null;index:question"`
	Answer   string `gorm:"type:varchar(100);not null;index:answer"`
}

func Repos(question string,answer string) {
	dsn := "root:zr444251196@tcp(127.0.0.1:3306)/shino_data?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&data{ID: int(CurrentId), Question: question, Answer: answer})
	CurrentId++
	fmt.Printf("record is recorded\n")
}

func init() {

	dsn := "root:zr444251196@tcp(127.0.0.1:3306)/shino_data?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&data{})

	var dt data
	db.Last(&dt)
	CurrentId = int64(dt.ID) + 1

	fmt.Printf("anaan %v",CurrentId)


	//   db.Create(&data{ID: 1, Question: "nihao",Answer: "zaijian"})
	//   db.Create(&data{ID: 2, Question: "nihao",Answer: "zaijian"})

}