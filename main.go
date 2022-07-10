package main

import (
	"shinobot/plugin/dplugin"

	"shinobot/sbot"

	"shinobot/plugin/test"
)
func main(){

	a := sbot.NewBot()

	a.LoadPlugin(dplugin.Export())
	a.LoadPlugin(test.Export())

	a.Run()
}