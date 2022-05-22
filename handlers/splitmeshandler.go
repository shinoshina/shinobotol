package handlers

import (
	"gocqserver/requester/pv"
	"gocqserver/senda"

	"github.com/gin-gonic/gin"
)

func ImgHandler(c *gin.Context, message map[string]interface{}) {
	pv.GetImg()
	c.JSON(200, gin.H{
		"reply":        "[CQ:image,file=file:///home/shinoshina/gocode/src/gocqserver/sese/bukeyisese!2.jpg]",
		"auto_escape":  false,
		"at_sender":    false,
		"delete":       false,
		"kick":         false,
		"ban":          false,
		"ban_duration": 0,
	})
	pv.GetHTML()

}

func SpeakHandler(c *gin.Context, message map[string]interface{}) {

	content := message["raw_message"].(string)[len("read:"):len(message["raw_message"].(string))]

	voice := "[CQ:tts,text=" + content + "]"

	senda.SendMessage(voice,message["group_id"].(float64))

}

func ShinoSpeakHandler(c *gin.Context, message map[string]interface{} ){

	c.JSON(200, gin.H{
		"reply":        "[CQ:record,file=file:///home/shinoshina/gocode/src/gocqserver/audio/shinobu894.mp3]",
		"auto_escape":  false,
		"at_sender":    false,
		"delete":       false,
		"kick":         false,
		"ban":          false,
		"ban_duration": 0,
	})



}
