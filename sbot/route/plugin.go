package route

type Plugin struct {
	ms *MessageSet
	es EventSet
}

func NewPlugin() (p *Plugin) {
	p = new(Plugin)
	p.ms = NewMessageSet()
	p.es = NewEventSet()
	return
}

func (p *Plugin) OnMessage(r string, mode string, handler func(d DataMap)) {
	p.ms.onMessage(r, mode, handler)
}
func (p *Plugin) OnEvent(ev string, handler func(d DataMap)) {
	p.es.onEvent(ev, handler)
}
