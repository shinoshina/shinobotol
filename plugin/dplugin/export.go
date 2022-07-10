package dplugin

import (
	"shinobot/sbot/request"
	"shinobot/sbot/route"
)

func Export() (p *route.Plugin) {
	p = route.NewPlugin()
	p.OnMessage("/", "all", NormalMessageHandler)
	p.OnMessage("read:","part",SpeakHandler)
	p.OnMessage("ceshi","part",func(d route.DataMap) {
		request.SendMessage("ceshi",d["group_id"].(float64))
	})
	p.OnEvent("poke",PokeHandler)
	return

}
