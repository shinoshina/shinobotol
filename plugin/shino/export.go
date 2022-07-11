package shino

import (
	"shinobot/sbot/request"
	"shinobot/sbot/route"
)

func Export()(p *route.Plugin){

	p = route.NewPlugin()
	p.OnMessage("小忍","all",audioTest)
	p.OnMessage("[CQ:at,qq=2037310389]","all",func(d route.DataMap) {
		request.SendMessage("[CQ:record,file=file:///home/shinoshina/gocode/src/gocqserver/assets/shino/intro.mp3]",d["group_id"].(float64))
	})
	p.OnMessage("kaka","all",func(d route.DataMap) {
		request.SendMessage("[CQ:record,file=file:///home/shinoshina/gocode/src/gocqserver/assets/shino/correct.mp3]",d["group_id"].(float64))
	})
	p.OnMessage("kaka!","all",func(d route.DataMap) {
		request.SendMessage("[CQ:record,file=file:///home/shinoshina/gocode/src/gocqserver/assets/shino/correct_answer.mp3]",d["group_id"].(float64))
	})
	return
}