package test

import (
	"shinobot/sbot/request"
	"shinobot/sbot/route"
)

func Export() (p *route.Plugin) {
	p = route.NewPlugin()
	p.OnMessage("multi plugin test","part",func(d route.DataMap) {
		request.SendMessage("multi plugin test",d["group_id"].(float64))
	})
	
	return

}
