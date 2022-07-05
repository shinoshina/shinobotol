package handlers

import (
	"fmt"
	"gocqserver/data"
	"gocqserver/requester/pv"
	"gocqserver/senda"
	"regexp"
	"strconv"
	"strings"

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

	senda.SendMessage(voice, message["group_id"].(float64))

}

func ShinoSpeakHandler(c *gin.Context, message map[string]interface{}) {

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

func RemoveHandler(c *gin.Context, message map[string]interface{}) {

	msgRule := regexp.MustCompile(`\[CQ:reply,id=(?P<message_id>\-?\d{1,})\]`)
	result := msgRule.FindStringSubmatch(message["raw_message"].(string))

	//[CQ:reply,id=\-(?P<message_id>\d{1,})]

	id_s_raw := result[1]
	var id_int int
	var id_s string
	if strings.Contains(id_s_raw, "-") {

		id_s = id_s_raw[len("-"):]
		id_int, _ = strconv.Atoi(id_s)
		id_int *= -1
	} else {
		id_int, _ = strconv.Atoi(id_s_raw)

	}

	str := senda.GetMessage(int32(id_int))


	if strings.Contains(str,"[CQ:image,file=") {
		imgRule := regexp.MustCompile(`\[CQ:image,file=.*\.image`)
		imgResult := imgRule.FindStringSubmatch(str)

		// subTypeRule := regexp.MustCompile(`,subType=\d]`)
		// subTypeResult := subTypeRule.FindStringSubmatch(str)

		str = imgResult[0] + ",subType=0]"
	    str1 := imgResult[0] + ",subType=1]"

		data.RemoveRaw(str1)
		data.RemoveRaw(str)

		ResetGlobalMessage()
		senda.SendMessage("emmm",message["group_id"].(float64))
		return

	}
	ResetGlobalMessage()

	fmt.Printf("shit enter %v\n",str)

	senda.SendMessage("emmm",message["group_id"].(float64))
	data.RemoveRaw(str)

}
