package handlers

import (
	"fmt"
	"gocqserver/data"

	"github.com/gin-gonic/gin"
)

//nonono this!!!! another copy!!!!!!!!!
var normal Handler = func(c *gin.Context, message map[string]interface{}) {

}

var (
	gloabalMessageForRepeat       = []string{"nihao", "test2", "test3", "baka"}
	gloabalMessageForLearning     = []string{"nihao", "test2", "test3", "baka"}
	messageCounterForRepeat   int = 0
	messageCounterForLearning int = 0
)

//yes yes this!!!
func NormalMessageHandler(c *gin.Context, message map[string]interface{}) {

	answerOk, answer := data.Find(message["raw_message"].(string))
	repeatOk := Check(message["raw_message"].(string))
	learnOk := CheckLearning(message["raw_message"].(string))

	if answerOk {
		fmt.Println("find answer!")
		c.JSON(200, gin.H{
			"reply":        answer,
			"auto_escape":  false,
			"at_sender":    false,
			"delete":       false,
			"kick":         false,
			"ban":          false,
			"ban_duration": 0,
		})
	} else if repeatOk {
		fmt.Println("repeat!")
		c.JSON(200, gin.H{
			"reply":        message["raw_message"],
			"auto_escape":  false,
			"at_sender":    false,
			"delete":       false,
			"kick":         false,
			"ban":          false,
			"ban_duration": 0,
		})
	} else if learnOk {
		fmt.Println("learn!")
		data.Repos(gloabalMessageForLearning[0], gloabalMessageForLearning[1])
		c.JSON(200, gin.H{
			"reply":        "recorded!",
			"auto_escape":  false,
			"at_sender":    false,
			"delete":       false,
			"kick":         false,
			"ban":          false,
			"ban_duration": 0,
		}) //[CQ:image,file=./sese/bukeyisese!.jpg,type=flash]
	}

}

func Check(currentMessage string) bool {

	var counter int = 1

	messageCounterForRepeat++
	messageCounterForRepeat %= 4
	gloabalMessageForRepeat[messageCounterForRepeat] = currentMessage

	for _, v := range gloabalMessageForRepeat {
		if v == currentMessage {
			counter++
		}
	}
	if counter >= 4 {
		messageCounterForRepeat = 0
		gloabalMessageForRepeat = []string{"nihao", "test2", "test3", "baka"} //QUESTION : HERERERERERE CAN DELETE
		return true
	} else {
		return false
	}
}
func CheckLearning(currentMessage string) bool {

	gloabalMessageForLearning[messageCounterForLearning] = currentMessage
	messageCounterForLearning++

	if messageCounterForLearning == 4 {
		messageCounterForLearning = 0
		if gloabalMessageForLearning[0] == gloabalMessageForLearning[2] &&
			gloabalMessageForLearning[1] == gloabalMessageForLearning[3] &&
			gloabalMessageForLearning[0] != gloabalMessageForLearning[1] {
			return true
		} else {
			return false
		}
	}
	return false

}
