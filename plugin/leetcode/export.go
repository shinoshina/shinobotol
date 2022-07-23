package leetcode

import (
	"shinobot/sbot/route"
	"shinobot/sbot/tick"
)

func Export() (p *route.Plugin) {

	p = route.NewPlugin("leetcode", "loaded")

	p.OnTick("leetcode", tick.Every(1*tick.Day).At("18:42"), SendLeetcodeInfo)
	// p.OnBoot(func() {
	// 	p.StartTask("leetcode")
	// })
	p.OnMessage("订阅每日一题","all",subscribe)
	p.OnMessage("取消订阅力扣","all",unsubscribe)
	return
}
