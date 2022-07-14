package sbot

import (
	"shinobot/sbot/route"
	"time"

	"github.com/gin-gonic/gin"
)

type Sbot struct {
	r     *route.Router
	e     *gin.Engine
	boots []func()
}

func NewBot() (sb *Sbot) {
	sb = new(Sbot)
	sb.r = route.NewRouter()
	sb.e = gin.Default()
	sb.boots = make([]func(), 1)
	return
}

func (sb *Sbot) mainHandler(c *gin.Context) {

	d := route.NewDataMap()
	c.BindJSON(&d)
	sb.r.Handle(d)

}
func (sb *Sbot) Run() {
	sb.e.POST("/", sb.mainHandler)
	for _, h := range sb.boots {
		if h != nil {
			go setTimeval(5, h)
		}
	}
	sb.e.Run(":5701")
}

func (sb *Sbot) LoadPlugin(p *route.Plugin) {
	sb.r.LoadPlugin(p)
	sb.boots = append(sb.boots, p.Boot())
}
func setTimeval(sec int64, h func()) {
	timer := time.NewTimer(time.Duration(sec) * time.Second)
	<-timer.C
	h()
	timer.Stop()
}
