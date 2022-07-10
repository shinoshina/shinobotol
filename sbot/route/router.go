package route

import "fmt"

type Router struct{
	m MessageSet
	a ActivitySet
}

func(r *Router) RegisterSet(om MessageSet){

}
func(r *Router) RegisterActivity(oa ActivitySet){

}
func (r *Router) Handle(d DataMap){

	pt,pmt,pst := d.SpiltType()
	if(pt == "meta"){
		fmt.Println("meta_event : check living")
	}else if(pt == "message"){
	}

}