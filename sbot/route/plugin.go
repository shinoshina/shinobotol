package route

import "fmt"

type Plugin struct {
	ms   *MessageSet
	es   EventSet
	bh   func()
	name string
}

func NewPlugin(n string) (p *Plugin) {
	p = new(Plugin)
	p.name = n
	p.ms = NewMessageSet()
	p.es = NewEventSet()
	p.bh = func() {
		fmt.Println("plugin " + p.name + " booting")
	}
	return
}

func (p *Plugin) OnMessage(r string, mode string, handler func(d DataMap)) {
	p.ms.onMessage(r, mode, handler)
}
func (p *Plugin) OnEvent(ev string, handler func(d DataMap)) {
	p.es.onEvent(ev, handler)
}

func (p *Plugin) OnBoot(handler func()) {
	p.bh = handler
}
func (p *Plugin) Boot() func() {
	return p.bh
}
