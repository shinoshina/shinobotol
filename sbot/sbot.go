package sbot

import (
	"shinobot/sbot/route"

	"github.com/gin-gonic/gin"
)

type Sbot struct {
	b *route.Router
	r *gin.Engine
}

func NewBot() (sb *Sbot) {
	sb = new(Sbot)
	sb.b = route.NewRouter()
	sb.r = gin.Default()
	return
}

func (sb *Sbot)mainHandler(c *gin.Context) {

	d := route.NewDataMap()
	c.BindJSON(&d)
	sb.b.Handle(d)

}
func (sb *Sbot) Run() {
	sb.r.POST("/", sb.mainHandler)
	sb.r.Run(":5701")
}

func (sb *Sbot) LoadPlugin(p *route.Plugin){
	sb.b.LoadPlugin(p)
}
