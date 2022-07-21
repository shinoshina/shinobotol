package route

import (
	"fmt"
	"shinobot/sbot/request"
	"shinobot/sbot/tick"
)

type Plugin struct {
	ms      *MessageSet
	es      EventSet
	bh      func()
	name    string
	trigger Trigger
	state   string
	ctm     map[string](*tick.CronTask)
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
	p.ctm = make(map[string]*tick.CronTask)
	p.bh = func() {
		fmt.Println("plugin: " + p.name + " onboot\n" + "initial state: " + p.state)
		if p.state == "shut" {
			fmt.Println("nothing happen")
		} else if p.state == "loaded" {
			fmt.Println("default handler start")
			for k, v := range p.ctm {
				fmt.Println(k + "start")
				v.Start()
			}
		}
	}

	p.trigger = Trigger{
		kload: p.name + "load",
		kshut: p.name + "shut",
		th: func(d DataMap, pluginState string) {
			if pluginState == "shut" {
				request.SendMessage(p.name+"onshut", d.GroupID())
				p.ShutDownAllCronTask()
			} else if pluginState == "loaded" {
				request.SendMessage(p.name+"onloaded", d.GroupID())
				p.bh()
			}
		},
	}
	return
}
func (p *Plugin) OnTick(name string,rule string, task func()) {

	ct := tick.NewCronTask(name,rule,task)
	p.ctm[name] = ct

}
func (p *Plugin) OnTrigger(keyLoad string, keyShut string, hook func(d DataMap, pluginState string)) {
	p.trigger.kload = keyLoad
	p.trigger.kshut = keyShut
	p.trigger.th = func(d DataMap, pluginState string) {
		if pluginState == "shut" {
			p.ShutDownAllCronTask()
		} else if pluginState == "loaded" {
			p.bh()
		}
		hook(d, pluginState)
	}
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
		if p.state == "shut" {
			fmt.Println("nothing happen")
		} else if p.state == "loaded" {
			fmt.Println("handler start")
			handler()
		}
	}
}
func (p *Plugin) Boot() func() {
	return p.bh
}
func (p *Plugin) ShutDownAllCronTask() {
	for _, v := range p.ctm {
		v.Stop()
	}
}
func (p *Plugin) StartTask(name string) {
	if _, ok := p.ctm[name]; ok {
		p.ctm[name].Start()
	}
}
