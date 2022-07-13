package leetcode

import "shinobot/sbot/route"

func Export()(p *route.Plugin){

	p = route.NewPlugin()
	p.OnMessage("每日一题","all",dailyQuestionInfo)
	return
}