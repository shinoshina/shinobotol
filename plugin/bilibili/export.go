package bilibili

import "shinobot/sbot/route"


func Export()(p *route.Plugin){

	p = route.NewPlugin("bilibili")
	p.OnMessage("看看批","all",getUerInfo)
	return
}