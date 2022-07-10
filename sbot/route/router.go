package route

import "fmt"

type Router struct{
	m *MessageSet
	n *NoticeSet
}

func(r *Router) RegisterSet(om MessageSet){

}
func(r *Router) RegisterActivity(on NoticeSet){

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

}