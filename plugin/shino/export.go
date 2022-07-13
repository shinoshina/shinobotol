package shino

import (
	"shinobot/sbot/request"
	"shinobot/sbot/route"
)

func Export()(p *route.Plugin){

	p = route.NewPlugin("shinoaudio")
	p.OnMessage("小忍","all",intro)
	p.OnMessage("[CQ:at,qq=2037310389]","all",intro)
	p.OnMessage("kaka","all",kaka1)
	p.OnMessage("kaka！！","all",kaka2)
	p.OnMessage("kaka~~","all",kaka3)
	p.OnMessage("小忍透透","part",func(d route.DataMap) {
		request.SendMessage("[CQ:record,file=file:///home/shinoshina/gocode/src/gocqserver/assets/shino/curse.mp3]",d["group_id"].(float64))
	})
		
	
	return
}