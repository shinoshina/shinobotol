package route

import (
	"shinobot/sbot/logger"
	"shinobot/sbot/request"
	"shinobot/sbot/tick"
)

type Plugin struct {
	ms      *MessageSet
	es      EventSet
	bl      *blacklist
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
	p.bl = NewBlackList()
	p.state = defaultState
	p.ctm = make(map[string]*tick.CronTask)
	p.bh = func() {
		logger.Info("plugin: " + p.name + " onboot\n" + "initial state: " + p.state)
		if p.state == "shut" {
			logger.Info("nothing happen")
		} else if p.state == "loaded" {
			logger.Info("default handler start")
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
			}
		},
	}
	return
}
func (p *Plugin) OnTick(name string, schedule tick.Schedule, task func()) {

	ct := tick.NewCronTask(name, schedule, task)
	p.ctm[name] = ct

}
func (p *Plugin) OnTrigger(keyLoad string, keyShut string, hook func(d DataMap, pluginState string)) {
	p.trigger.kload = keyLoad
	p.trigger.kshut = keyShut
	p.trigger.th = func(d DataMap, pluginState string) {
		if pluginState == "shut" {
			p.ShutDownAllCronTask()
		} else if pluginState == "loaded" {
		}
		hook(d, pluginState)
	}
}
func (p *Plugin) OnMessage(r string, mode string, handler func(d DataMap)) {
	if p.bl.blf {
		handler = func(d DataMap) {
			if v, ok := p.bl.blm[d.GroupID()]; ok {
				if v {
					return
				}
			}
			handler(d)
		}
	}
	p.ms.onMessage(r, mode, handler)
}
func (p *Plugin) OnEvent(ev string, handler func(d DataMap)) {
	p.es.onEvent(ev, handler)
}

func (p *Plugin) OnBoot(handler func()) {
	p.bh = func() {
		logger.Info("plugin: " + p.name + " onboot\n" + "initial state: " + p.state)
		if p.state == "shut" {
			logger.Info("nothing happen")
		} else if p.state == "loaded" {
			logger.Info("handler start")
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
