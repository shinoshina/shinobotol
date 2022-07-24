package bilibili

import (
	"shinobot/sbot/route"
	"shinobot/sbot/tick"
)

func Export() (p *route.Plugin) {

	p = route.NewPlugin("bilibili", "loaded")
	p.OnMessage(`^关注主播喵:(?P<mid>.*)`, "regex", subscribe)
	p.OnMessage(`^取关了喵:(?P<mid>.*)`, "regex", unsubscribe)
	p.OnMessage("debug","all",checkStates)
	p.OnTick("checkstatus", tick.Every(5*tick.Minute), checkState)
	p.OnBoot(func ()  {
		p.StartTask("checkstatus")
	})

	return
}
