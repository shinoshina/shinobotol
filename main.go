package main

import (
	"shinobot/plugin/dplugin"

	"shinobot/sbot"

	"shinobot/plugin/test"

	"shinobot/plugin/shino"
)
func main(){

	a := sbot.NewBot()

	a.LoadPlugin(dplugin.Export())
	a.LoadPlugin(test.Export())
	a.LoadPlugin(shino.Export())


	a.Run()
}