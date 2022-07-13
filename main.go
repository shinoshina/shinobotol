package main

import (
	"shinobot/plugin/bilibili"
	"shinobot/plugin/dplugin"
	"shinobot/plugin/leetcode"

	"shinobot/sbot"

	"shinobot/plugin/test"

	"shinobot/plugin/shino"
)
func main(){

	a := sbot.NewBot()

	a.LoadPlugin(dplugin.Export())
	a.LoadPlugin(test.Export())
	a.LoadPlugin(shino.Export())
	a.LoadPlugin(leetcode.Export())
	a.LoadPlugin(bilibili.Export())


	a.Run()
}