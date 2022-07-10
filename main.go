package main

import (
	"shinobot/plugin/dplugin"
	"shinobot/sbot"
)
func main(){

	a := sbot.NewBot()

	a.LoadPlugin(dplugin.Export())

	a.Run()
}