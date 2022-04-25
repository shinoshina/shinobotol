package handlers

import (
	"gocqserver/requester"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PokeHandler(c *gin.Context, message map[string]interface{}) {

	resPoster := requester.RequestPoster{
		Client: &http.Client{},
	}
	resPoster.PostPoke(message)
}
