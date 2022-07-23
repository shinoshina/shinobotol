package test

import (
	"shinobot/sbot/request"
	"shinobot/sbot/route"
	"shinobot/sbot/tick"
)
func Export() (p *route.Plugin) {
	p = route.NewPlugin("test", "loaded")
	p.OnMessage("multi plugin test", "part", func(d route.DataMap) {
		request.SendMessage("multi plugin test", d["group_id"].(float64))
	})
	p.OnMessage(`^regex test`, "regex", func(d route.DataMap) {
		request.SendMessage("regex test", d["group_id"].(float64))
	})
	p.OnMessage(`^name:(?P<name>.*?) age:(?P<age>.*)`, "regex", func(d route.DataMap) {
		vmap := d["group_value"].(map[string](string))
		for k, v := range vmap {
			request.SendMessage("key: "+k+"  "+"value: "+"  "+v, d["group_id"].(float64))
		}
	})
	p.OnMessage("TEST", "all", func(d route.DataMap) {
		request.SendMessage(d.Message(), d.GroupID())
	})

	p.OnMessage("div", "all", func(d route.DataMap) {
		ct := tick.NewCronTask("sb", tick.Every(1*tick.Day).At("11:58"), func() {
			request.SendMessage(d.Message(), d.GroupID())
		})
		ct.Start()
	})

	p.OnTick("hello", tick.Every(1*tick.Day).At("11:58"), func() {
		request.SendMessage("hello ha", 1012330112)
	})

	p.OnTrigger("testload", "testshut", func(d route.DataMap, pluginState string) {
		if pluginState == "loaded" {
			request.SendMessage("testonloaded", d.GroupID())
		} else if pluginState == "shut" {
			request.SendMessage("testonshut", d.GroupID())
		}
	})
	return

}
