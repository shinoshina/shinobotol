package dplugin

import (
	"shinobot/sbot/request"
	"shinobot/sbot/route"
	"strconv"
)

func SpeakHandler(d route.DataMap) {

	// content := d["raw_message"].(string)[len("read:"):len(d["raw_message"].(string))]

	vmap := d["group_value"].(map[string]string)
	rc := vmap["content"]
	voice := "[CQ:tts,text=" + rc + "]"

	request.SendMessage(voice, d["group_id"].(float64))

}
func PokeHandler(d route.DataMap) {

	group_id := d["group_id"].(float64)
	id := strconv.FormatInt(int64(d["sender_id"].(float64)), 10)

	if id != "2037310389" {
		poke := "[CQ:poke,qq=" + id + "]"
		request.SendMessage(poke, group_id)
	}

}
