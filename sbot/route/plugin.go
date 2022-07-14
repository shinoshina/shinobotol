package route

import (
	"context"
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
	ptm     map[string](*PeriodicalTask)
}
type PeriodicalTask struct {
	name   string
	task   func()
	state  string
	ctx    context.Context
	cancel context.CancelFunc
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
	p.ptm = make(map[string]*PeriodicalTask)
	p.bh = func() {
		fmt.Println("plugin: " + p.name + " onboot\n" + "initial state: " + p.state)
	}
	p.trigger = Trigger{
		kload: p.name + "load",
		kshut: p.name + "shut",
		th: func(d DataMap, pluginState string) {
			if pluginState == "shut" {
				request.SendMessage(p.name+"onshut", d.GroupID())
				p.ShutDownAllPTask()
			} else if pluginState == "loaded" {
				request.SendMessage(p.name+"onloaded", d.GroupID())
				p.bh()
			}
		},
	}
	return
}
func (p *Plugin) OnTick(name string, task func()) {

	p.newPeriodicalTask(name, task)

}
func (p *Plugin) OnTrigger(keyLoad string, keyShut string, hook func(d DataMap, pluginState string)) {
	p.trigger.kload = keyLoad
	p.trigger.kshut = keyShut
	p.trigger.th = func(d DataMap, pluginState string) {
		if pluginState == "shut" {
			p.ShutDownAllPTask()
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
			handler()
		}

	}
}
func (p *Plugin) Boot() func() {
	return p.bh
}
func (p *Plugin) ShutDownAllPTask() {
	for _, v := range p.ptm {
		v.stop()
	}
}
func (p *Plugin) newPeriodicalTask(name string, task func()) (pt *PeriodicalTask) {
	pt = new(PeriodicalTask)
	pt.name = name
	pt.task = task
	pt.state = "off"
	pt.ctx, pt.cancel = context.WithCancel(context.Background())
	p.ptm[name] = pt
	return
}
func (p *Plugin) StartTask(name string) {
	if _, ok := p.ptm[name]; ok {
		p.ptm[name].start()
	}
}
func (pt *PeriodicalTask) start() {
	go func() {
		if pt.state == "off" {
			pt.state = "on"
			for {
				select {
				case <-pt.ctx.Done():
					return
				default:
					pt.task()
				}
			}
		} else if pt.state == "on" {
			return
		}
	}()
}
func (pt *PeriodicalTask) stop() {
	if pt.state == "on" {
		pt.cancel()
	} else if pt.state == "off" {
		return
	}
}
