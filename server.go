package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"gocqserver/repo"
	"gocqserver/data"
)

type messageChecker []string

var gloabalMessage = messageChecker{"nihao", "test2", "test3", "baka"}
var gloabalMessageForLearning = messageChecker{"nihao", "test2", "test3", "baka"}
var messageCounter int = 0
var messageCounterForLearning int = 0


func (messageList messageChecker) Check(currentMessage string) bool {

	var counter int = 1

	messageCounter++
	messageCounter %= 4
	gloabalMessage[messageCounter] = currentMessage

	for _, v := range messageList {
		if v == currentMessage {
			counter++
		}
	}
	if counter >= 4 {
		messageCounter = 0
		gloabalMessage = messageChecker{"nihao", "test2", "test3", "baka"}
		return true
	} else {
		return false
	}
}
func (messageList messageChecker) CheckLearning(currentMessage string) bool {

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

func MessageHandler(c *gin.Context) {
	mes := make(map[string]interface{})
	c.BindJSON(&mes)

	//message repeat
	if mes["post_type"] == "message" {
		if mes["message_type"] == "group" {

			if gloabalMessage.Check(mes["raw_message"].(string)) {
				fmt.Println(mes["raw_message"])

				c.JSON(200, gin.H{
					"reply":        mes["raw_message"],
					"auto_escape":  false,
					"at_sender":    false,
					"delete":       false,
					"kick":         false,
					"ban":          false,
					"ban_duration": 0,
				})
			} else if gloabalMessage.CheckLearning(mes["raw_message"].(string)) {

				repo.Repos(gloabalMessageForLearning[0], gloabalMessageForLearning[1])
				c.JSON(200, gin.H{
					"reply":        "recorded!",
					"auto_escape":  false,
					"at_sender":    false,
					"delete":       false,
					"kick":         false,
					"ban":          false,
					"ban_duration": 0,
				})

			}
		} else if mes["message_type"] == "private" {
			c.JSON(200, gin.H{
				"reply":       mes["raw_message"],
				"auto_escape": false,
			})
		}

	}

	//group chat poke
	if mes["post_type"] == "notice" {
		if mes["notice_type"] == "notify" {
			if mes["sub_type"] == "poke" {
				if _, ok := mes["group_id"]; ok {
					if mes["self_id"] == mes["target_id"] {
						resPoster := data.RequestPoster{
							Client: &http.Client{},
						}
						resPoster.PostPoke(mes)
					}
				}
			}
		}
	}
}

func main() {

	r := gin.Default()
	r.POST("/", MessageHandler)
	r.Run(":5701") // listen and serve on 0.0.0.0:8080
}
