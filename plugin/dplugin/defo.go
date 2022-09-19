package dplugin

import (
	"shinobot/sbot/logger"
	"shinobot/sbot/repo/datas"
	"shinobot/sbot/request"
	"shinobot/sbot/route"
)
var(
	db *datas.Db
)
var (
	gloabalMessageForRepeat       = []string{"nihao", "test2", "test3", "baka"}
	gloabalMessageForLearning     = []string{"nihao", "test2", "test3", "baka"}
	messageCounterForRepeat   int = 0
	messageCounterForLearning int = 0
)
func init(){
	db = datas.CreateDb("assets/default/dialog")
}

func ResetGlobalMessage(){

	gloabalMessageForRepeat       = []string{"nihao", "test2", "test3", "baka"}
	gloabalMessageForLearning =  []string{"nihao", "test2", "test3", "baka"}
}
//yes yes this!!!
func NormalMessageHandler(d route.DataMap) {

	msg := d["raw_message"].(string)
	group_id := d["group_id"].(float64)
	answerOk, answer := db.Get(msg)
	repeatOk := Check(msg)
	learnOk := CheckLearning(msg)

	if answerOk {
		logger.Info("find answer! %v\n",answer)
		request.SendMessage(answer,group_id)
	} else if repeatOk {
		logger.Info("repeat!")
		request.SendMessage(msg,group_id)
	} else if learnOk {
		logger.Info("learn!")
		//data.Repos(gloabalMessageForLearning[0], gloabalMessageForLearning[1])
		db.Put(gloabalMessageForLearning[0],gloabalMessageForLearning[1])
		request.SendMessage(msg,group_id)
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
	for i :=0;i<3;i++{
		gloabalMessageForLearning[i] = gloabalMessageForLearning[i+1]
	}
	gloabalMessageForLearning[3] = currentMessage
	if gloabalMessageForLearning[0] == gloabalMessageForLearning[2] &&
			gloabalMessageForLearning[1] == gloabalMessageForLearning[3] &&
			gloabalMessageForLearning[0] != gloabalMessageForLearning[1] {
			return true
		} else {
			return false
		}

}


