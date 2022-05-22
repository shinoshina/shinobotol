package handlers

import (
	"gocqserver/senda"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PokeHandler(c *gin.Context, message map[string]interface{}) {

	group_id  := message["group_id"].(float64)
	id := strconv.FormatInt(int64(message["sender_id"].(float64)), 10)

	if id != "2037310389" {
		poke := "[CQ:poke,qq=" + id + "]"
		senda.SendMessage(poke,group_id)
	}

}
