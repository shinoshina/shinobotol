package main

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var db *gorm.DB

var CurrentId int64
var CrucialId int32

type data struct {
	ID       int   `gorm:"primary_key;index:id"`
	Question string `gorm:"type:varchar(100);not null;index:question"`
	Answer   string `gorm:"type:varchar(100);not null;index:answer"`
}

func init() {

	dsn := "root:zr444251196@tcp(127.0.0.1:3306)/shino_data?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	  }

	  db.AutoMigrate(&data{})

	  db.Create(&data{ID: 1, Question: "nihao",Answer: "zaijian"})
	  db.Create(&data{ID: 2, Question: "nihao",Answer: "zaijian"})
	

}
