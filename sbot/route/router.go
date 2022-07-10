package route

import "fmt"

type Router struct{
	m *MessageSet
	n NoticeSet
}

func(r *Router) LoadPlugin(ms* MessageSet){

	for key,handler := range ms.ma{
		r.m.ma[key] = handler
	}
	for key,handler := range ms.mp{
		r.m.mp[key] = handler
	}
	for key,handler := range ms.mr{
		r.m.mr[key] = handler
	}
	r.m.ma["/"] = ms.ma["/"]
}
func(r *Router) RegisterActivity(on NoticeSet){

}
func NewRouter()(r *Router){
	r = new(Router)
	r.m = NewMessageSet()
	r.n = NewNoticeSet()
	return
}
func (r *Router) Handle(d DataMap){

	pt,pmt,pst := d.SpiltType()
	if(pt == "meta"){
		fmt.Println("meta_event : check living")
	}else if(pt == "message"){
		r.m.handle(d)
	}else if(pt == "notice"){

	}else{
		
	}
	fmt.Println(pt)
	fmt.Println(pmt)
	fmt.Println(pst)

}