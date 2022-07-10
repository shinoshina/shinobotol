package route

import "fmt"

type Router struct {
	ms *MessageSet
	es EventSet
}

func (r *Router) LoadPlugin(p *Plugin) {

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
func NewRouter() (r *Router) {
	r = new(Router)
	r.ms = NewMessageSet()
	r.es = NewEventSet()
	return
}
func (r *Router) Handle(d DataMap) {

	pt, pmt, pst := d.SpiltType()
	if pt == "meta" {
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
