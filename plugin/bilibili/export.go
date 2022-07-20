package bilibili

import "shinobot/sbot/route"


func Export()(p *route.Plugin){

	p = route.NewPlugin("bilibili","loaded")
	p.OnMessage(`^关注主播喵:(?P<mid>.*)`,"regex",Subscribe)
	// p.OnTick("hha",Tick)
	return
}