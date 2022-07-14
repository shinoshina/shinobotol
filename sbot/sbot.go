package sbot

import (
	"fmt"
	"shinobot/sbot/route"
	"time"

	"github.com/gin-gonic/gin"
)

type Sbot struct {
	r     *route.Router
	e     *gin.Engine
	bh func()
}

func NewBot() (sb *Sbot) {
	sb = new(Sbot)
	sb.r = route.NewRouter()
	sb.e = gin.Default()
	sb.bh = func() {
		fmt.Println("免费啦啦啦啦啦")
	}
	return
}

func (sb *Sbot) mainHandler(c *gin.Context) {

	d := route.NewDataMap()
	c.BindJSON(&d)
	sb.r.Handle(d)

}
func (sb *Sbot) Run() {
	sb.e.POST("/", sb.mainHandler)
	go setTimeval(5,sb.bh)
	sb.e.Run(":5701")
}

func (sb *Sbot) LoadPlugin(p *route.Plugin) {
	sb.r.LoadPlugin(p)
}
func setTimeval(sec int64, h func()) {
	timer := time.NewTimer(time.Duration(sec) * time.Second)
	<-timer.C
	h()
	timer.Stop()
}
