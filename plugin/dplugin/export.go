package dplugin

import "shinobot/sbot/route"

func Export() (ms *route.MessageSet) {
	ms = route.NewMessageSet()
	ms.OnMessage("/", "all", NormalMessageHandler)
	return

}
