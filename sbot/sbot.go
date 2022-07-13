package sbot

import (
	"shinobot/sbot/route"
	"time"

	"github.com/gin-gonic/gin"
)

type Sbot struct {
	b     *route.Router
	r     *gin.Engine
	boots []func()
}

func NewBot() (sb *Sbot) {
	sb = new(Sbot)
	sb.b = route.NewRouter()
	sb.r = gin.Default()
	sb.boots = make([]func(), 1)
	return
}

func (sb *Sbot) mainHandler(c *gin.Context) {

	d := route.NewDataMap()
	c.BindJSON(&d)
	sb.b.Handle(d)

}
func (sb *Sbot) Run() {
	sb.r.POST("/", sb.mainHandler)
	for _, h := range sb.boots {
		if h != nil {
			go setTimeval(5, h)
		}
	}
	sb.r.Run(":5701")
}

func (sb *Sbot) LoadPlugin(p *route.Plugin) {
	sb.b.LoadPlugin(p)
	sb.boots = append(sb.boots, p.Boot())
}
func setTimeval(sec int64, h func()) {
	timer := time.NewTimer(time.Duration(sec) * time.Second)
	<-timer.C
	h()
	timer.Stop()
}
