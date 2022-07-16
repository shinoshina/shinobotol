package test

import (
	"shinobot/sbot/request"
	"shinobot/sbot/route"
	"shinobot/sbot/tick"
	"time"
)

func Export() (p *route.Plugin) {
	p = route.NewPlugin("test","shut")
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

	p.OnTrigger("testload","testshut",func(d route.DataMap, pluginState string) {
		if pluginState == "loaded" {
			request.SendMessage("testonloaded",d.GroupID())
		}else if pluginState == "shut"{
			request.SendMessage("testonshut",d.GroupID())
		}
	})

	p.OnTick("定时任务",func(t *tick.Timer) {
		timer := time.NewTimer(5*time.Second)
		<-timer.C
		request.SendMessage("定时任务",1012330112)
	})
	p.OnBoot(func() {
		p.StartTask("定时任务")
	})
	return

}
