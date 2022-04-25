package handlers

import (
	"github.com/gin-gonic/gin"
	"gocqserver/requester"
)

func ImgHandler(c *gin.Context, message map[string]interface{}) {
	requester.GetImg()
	c.JSON(200, gin.H{
		"reply":        "[CQ:image,file=file:///home/shinoshina/gocode/src/gocqserver/sese/bukeyisese!2.jpg]",
		"auto_escape":  false,
		"at_sender":    false,
		"delete":       false,
		"kick":         false,
		"ban":          false,
		"ban_duration": 0,
	})

}

func SpeakHandler(c *gin.Context, message map[string]interface{}) {

	content := message["raw_message"].(string)[len("please read:"):len(message["raw_message"].(string))]

	voice := "[CQ:tts,text=" + content + "]"

	c.JSON(200, gin.H{
		"reply":        voice,
		"auto_escape":  false,
		"at_sender":    false,
		"delete":       false,
		"kick":         false,
		"ban":          false,
		"ban_duration": 0,
	})

}
