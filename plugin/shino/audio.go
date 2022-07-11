package shino

import (
	"shinobot/sbot/request"
	"shinobot/sbot/route"
)

func audioTest(d route.DataMap)  {
	request.SendMessage("[CQ:record,file=file:///home/shinoshina/gocode/src/gocqserver/assets/shino/intro.mp3]",d["group_id"].(float64))
}