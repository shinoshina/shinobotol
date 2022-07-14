package route

import "fmt"

type Router struct {
	ms     *MessageSet
	es     EventSet
	ps     map[string](*Plugin)
	booted bool
}

func (r *Router) LoadPlugin(p *Plugin) {

	if p.state == "loaded" {
		r.load(p)
	} else if p.state == "shut" {
		fmt.Println("不加载")
	}

	keyload := p.trigger.kload
	keyshut := p.trigger.kshut
	r.ms.ma[keyload] = func(d DataMap) {

		if p.state == "shut" {
			r.load(p)
			p.state = "loaded"
			p.trigger.th(d, "loaded")
		} else {
			fmt.Println("乌鱼子111")
		}
	}
	r.ms.ma[keyshut] = func(d DataMap) {
		if p.state == "loaded" {
			r.unload(p)
			p.state = "shut"
			p.trigger.th(d, "shut")
		} else {
			fmt.Println("乌鱼子222")
		}
	}

	r.ps[p.name] = p
}
func (r *Router) load(p *Plugin) {
	for key, handler := range p.ms.ma {
		r.ms.ma[key] = handler
	}
	for key, handler := range p.ms.mp {
		r.ms.mp[key] = handler
	}
	for key, handler := range p.ms.mr {
		r.ms.mr[key] = handler
	}
	if h, ok := p.ms.ma["/"]; ok {
		r.ms.ma["/"] = h
	}
	for key, handler := range p.es {
		r.es[key] = handler
	}
}
func (r *Router) unload(p *Plugin) {
	for key, _ := range p.ms.ma {
		if _, ok := r.ms.ma[key]; ok {
			delete(r.ms.ma, key)
		}
	}
	for key, _ := range p.ms.mp {
		if _, ok := r.ms.mp[key]; ok {
			delete(r.ms.mp, key)
		}
	}
	for key, _ := range p.ms.mr {
		if _, ok := r.ms.mr[key]; ok {
			delete(r.ms.mr, key)
		}
	}
	for key, _ := range p.es {
		if _, ok := r.es[key]; ok {
			delete(r.es, key)
		}
	}
}
func NewRouter() (r *Router) {
	r = new(Router)
	r.ms = NewMessageSet()
	r.es = NewEventSet()
	r.ps = make(map[string]*Plugin)
	return
}
func (r *Router) Handle(d DataMap) {

	pt, pmt, pst := d.SpiltType()
	if pt == "meta" {
		if !r.booted {
			for _, p := range r.ps {
				 p.bh()
			}
			r.booted = true
			fmt.Println("plugin first booted")
		}
		fmt.Println("meta_event : check living")
	} else if pt == "message" {
		r.ms.handle(d)
	} else if pt == "notice" {

		not := d["notice_type"].(string)
		if not == "notify" {
			subt := d["sub_type"].(string)
			r.es.handle(subt, d)
		} else {
			r.es.handle(not, d)
		}
	} else {

	}
	fmt.Println(pt)
	fmt.Println(pmt)
	fmt.Println(pst)

}
