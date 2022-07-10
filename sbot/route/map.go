package route

import (
	"strings"
	"shinobot/sbot/request"
)

type (
	DataMap map[string]interface{}

	MessageSet struct {
		ma map[string]func(d DataMap)
		mp map[string]func(d DataMap)
		mr map[string]func(d DataMap)
	}

	NoticeSet map[string]func(d DataMap)
)

func DefaultHandler(d DataMap) {
	gi := d["group_id"].(float64)
	request.SendMessage("gei ye pa!",gi)
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
func (ms *MessageSet) OnMessage(r string, mode string, handler func(d DataMap)) {
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
			if strings.Contains(msg,key){
				ms.mp[key](d)
				return
			}
		}
		// regex handle
		// wo bu hui xie hahah cnm
	}
	ms.ma["/"](d)

}
func NewNoticeSet() (n NoticeSet) {
	n = make(NoticeSet)
	return
}
func (a NoticeSet) OnEvent(tp string, handler func(d DataMap)) {
	a[tp] = handler
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
