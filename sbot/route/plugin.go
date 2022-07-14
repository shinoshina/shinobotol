package route

import (
	"fmt"
	"shinobot/sbot/request"
)

type Plugin struct {
	ms      *MessageSet
	es      EventSet
	bh      func()
	name    string
	trigger Trigger
	state   string
}

type Trigger struct {
	kload string
	kshut string
	th    func(d DataMap, state string)
}

func NewPlugin(n string, defaultState string) (p *Plugin) {
	p = new(Plugin)
	p.name = n
	p.ms = NewMessageSet()
	p.es = NewEventSet()
	p.state = defaultState
	p.bh = func() {
		fmt.Println("plugin: " + p.name + " onboot\n" + "initial state: " + p.state)
	}
	p.trigger = Trigger{
		kload: p.name+"load",
		kshut: p.name+"shut",
		th: func(d DataMap, pluginState string) {
			if pluginState == "shut" {
				fmt.Println(p.name + "  shut")
				request.SendMessage(p.name+"onshut",d.GroupID())
			} else if pluginState == "loaded" {
				request.SendMessage(p.name+"onloaded",d.GroupID())
				fmt.Println(p.name + "  loaded")
			}
		},
	}
	return
}
func (p *Plugin) OnTrigger(keyLoad string, keyShut string, hook func(d DataMap, pluginState string)) {
	p.trigger.kload = keyLoad
	p.trigger.kshut = keyShut
	p.trigger.th = hook
}
func (p *Plugin) OnMessage(r string, mode string, handler func(d DataMap)) {
	p.ms.onMessage(r, mode, handler)
}
func (p *Plugin) OnEvent(ev string, handler func(d DataMap)) {
	p.es.onEvent(ev, handler)
}

func (p *Plugin) OnBoot(handler func()) {
	p.bh = func() {
		fmt.Println("plugin: " + p.name + " onboot\n" + "initial state: " + p.state)
		handler()
	}
}
func (p *Plugin) Boot() func() {
	return p.bh
}
