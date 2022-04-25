package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)


func MessageHandler(c *gin.Context) {
	mes := make(map[string]interface{})
	c.BindJSON(&mes)
	mainType,subType,ok := SplitMessageType(mes)
	if ok {
		fmt.Printf("main: %v,sub: %v \n",mainType,subType)
		Do(mainType,subType,c,mes)
	}else{
		fmt.Println("gocq end check aliving")
	}

	



}