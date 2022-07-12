package shino

import (
	"shinobot/sbot/request"
	"shinobot/sbot/route"
)

var (
	kakaOrder = 1
)

func intro(d route.DataMap) {
	request.SendMessage("[CQ:record,file=file:///home/shinoshina/gocode/src/gocqserver/assets/shino/intro.mp3]", d["group_id"].(float64))
}

func kaka1(d route.DataMap) {
	if kakaOrder == 1 {
		request.SendMessage("[CQ:record,file=file:///home/shinoshina/gocode/src/gocqserver/assets/shino/kaka1.mp3]", d["group_id"].(float64))
		kakaOrder = 2
	}

}
func kaka2(d route.DataMap) {
	if kakaOrder == 2 {
		request.SendMessage("[CQ:record,file=file:///home/shinoshina/gocode/src/gocqserver/assets/shino/kaka2.mp3]", d["group_id"].(float64))
		kakaOrder = 3
	}
}
func kaka3(d route.DataMap) {
	if kakaOrder == 3 {
		request.SendMessage("[CQ:record,file=file:///home/shinoshina/gocode/src/gocqserver/assets/shino/kaka3.mp3]", d["group_id"].(float64))
		kakaOrder = 1
	}
}
