package data

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var CurrentId int64
var CGID int64
var CrucialId int32

var Dsn = "root:zr444251196@tcp(127.0.0.1:3306)/shino_data?charset=utf8mb4&parseTime=True&loc=Local"

type data struct {
	ID       int    `gorm:"primary_key;index:id"`
	Question string `gorm:"type:varchar(256);not null;index:question"`
	Answer   string `gorm:"type:varchar(256);not null;index:answer"`
}
type streamRoom struct {
	ID      int    `gorm:"primary_key;index:id"`
	GroupId string `gorm:"type:varchar(36);not null;index:groupid"`
	Mid     string `gorm:"type:varchar(24);not null;index:mid"`
}

func init() {

	db, err := gorm.Open(mysql.Open(Dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&data{})
	db.AutoMigrate(&streamRoom{})

	var dt data
	db.Last(&dt)
	CurrentId = int64(dt.ID) + 1

	var sr streamRoom
	db.Last(&sr)
	CGID = int64(sr.ID) + 1
	fmt.Printf("anaan %v", CurrentId)

	InitDbPool(16)
	//   db.Create(&data{ID: 1, Question: "nihao",Answer: "zaijian"})
	//   db.Create(&data{ID: 2, Question: "nihao",Answer: "zaijian"})
}

func Repos(question string, answer string) {
	db, index := GetDb()
	db.Create(&data{ID: int(CurrentId), Question: question, Answer: answer})
	CurrentId++
	fmt.Printf("record is recorded\n")

	defer FinishTask(index)
}
func R(groupid string,room string){
	db,index := GetDb()
	db.Create(&streamRoom{ID:int(CGID),GroupId: groupid,Mid: room})
	CGID++
	defer FinishTask(index)
}

func Find(question string) (ok bool, answer string) {

	var dt data

	db, index := GetDb()

	db.Where("question = ?", question).First(&dt)

	defer FinishTask(index)

	if dt.ID == 0 {
		return false, ""
	} else {
		return true, dt.Answer
	}

}

func RemoveRaw(str string) (ok bool) {

	var dt data

	db, index := GetDb()

	db.Where("answer = ?", str).Delete(&dt)

	fmt.Println(dt)

	defer FinishTask(index)

	return true
}

func RemoveImg(str string) (ok bool) {

	var dt data

	db, index := GetDb()

	db.Where("answer like ?", "%"+str+"%").Delete(&dt)

	fmt.Println(dt)

	defer FinishTask(index)

	return true

}
