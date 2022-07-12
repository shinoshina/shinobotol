package test

import (
	"shinobot/sbot/request"
	"shinobot/sbot/route"
)

func Export() (p *route.Plugin) {
	p = route.NewPlugin()
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
	return

}
