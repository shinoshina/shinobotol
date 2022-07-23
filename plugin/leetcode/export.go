package leetcode

import (
	"shinobot/sbot/route"
	"shinobot/sbot/tick"
)

func Export() (p *route.Plugin) {

	p = route.NewPlugin("leetcode", "loaded")

	p.OnTick("leetcode", tick.Every(1*tick.Day).At("07:00"), dailyQuestionInfo)
	// p.OnBoot(func() {
	// 	p.StartTask("leetcode")
	// })
	return
}
