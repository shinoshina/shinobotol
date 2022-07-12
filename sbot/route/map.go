package route

import (
	"fmt"
	"regexp"
	"shinobot/sbot/request"
	"strings"
)

type (
	DataMap map[string]interface{}

	MessageSet struct {
		ma map[string]func(d DataMap)
		mp map[string]func(d DataMap)
		mr map[string]func(d DataMap)
	}

	EventSet map[string]func(d DataMap)
)

func DefaultHandler(d DataMap) {
	gi := d["group_id"].(float64)
	request.SendMessage("gei ye pa!", gi)
}
func NewMessageSet() (ms *MessageSet) {
	ms = new(MessageSet)
	ms.ma = make(map[string]func(d DataMap))
	ms.mp = make(map[string]func(d DataMap))
	ms.mr = make(map[string]func(d DataMap))
	return
}

func NewDataMap() (d DataMap) {
	d = make(DataMap)
	return
}
func (ms *MessageSet) onMessage(r string, mode string, handler func(d DataMap)) {
	if mode == "all" {
		ms.ma[r] = handler
	} else if mode == "part" {
		ms.mp[r] = handler
	} else if mode == "regex" {
		ms.mr[r] = handler
	}
}
func (ms *MessageSet) handle(d DataMap) {

	msg := d["raw_message"].(string)
	_, ok := ms.ma[msg]
	if ok {
		ms.ma[msg](d)
		return
	} else {
		for key, _ := range ms.mp {
			if strings.Contains(msg, key) {
				ms.mp[key](d)
				return
			}
		}
		for key, _ := range ms.mr {
			rule := regexp.MustCompile(key)
			result := rule.FindStringSubmatch(msg)
			if result != nil {
				group := rule.SubexpNames()
				if len(group) > 1 {
					vmap := make(map[string]string)
					for i, name := range group {
						if i != 0 && name != "" {
							vmap[name] = result[i]
						}
					}
					d["group_value"] = vmap
				}
				ms.mr[key](d)
				return
			} else {
				fmt.Println("here round no match")
			}
		}
	}
	ms.ma["/"](d)

}
func NewEventSet() (es EventSet) {
	es = make(EventSet)
	return
}
func (es EventSet) onEvent(ev string, handler func(d DataMap)) {
	es[ev] = handler
}
func (es EventSet) handle(ev string, d DataMap) {
	es[ev](d)
}
func (d DataMap) SpiltType() (pt string, pmt string, pst string) {

	pt = d["post_type"].(string)
	if pt == "meta_event" {
		pmt = "meta"
		pst = "meta"
		return
	} else if pt == "message" {
		pmt = d["message_type"].(string)
		if pmt == "group" {
			pst = d["sub_type"].(string)
			return
		} else {
			pt = "disallow_private"
			pst = "hei"
			return
		}
	} else if pt == "notice" {
		pmt = d["notice_type"].(string)
		if tmp, ok := d["sub_type"]; ok {
			pst = tmp.(string)
			return
		}
	}
	return
}
