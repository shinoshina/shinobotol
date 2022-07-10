package dplugin

import "shinobot/sbot/route"

func Export() (p *route.Plugin) {
	p = route.NewPlugin()
	p.OnMessage("/", "all", NormalMessageHandler)
	p.OnMessage("read:","part",SpeakHandler)
	p.OnEvent("poke",PokeHandler)
	return

}
