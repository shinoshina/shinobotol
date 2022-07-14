package bilibili

import "shinobot/sbot/route"


func Export()(p *route.Plugin){

	p = route.NewPlugin("bilibili","loaded")
	p.OnMessage("看看","all",getUerInfo)
	return
}