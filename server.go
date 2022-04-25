package main

import (
	
	"github.com/gin-gonic/gin"
	"gocqserver/handlers"
)



func main() {

	r := gin.Default()
	r.POST("/", handlers.MessageHandler)
	r.Run(":5701") // listen and serve on 0.0.0.0:8080
}
