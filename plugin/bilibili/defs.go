package bilibili

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shinobot/sbot/logger"
	"shinobot/sbot/repo/datas"
	"shinobot/sbot/request"
	"shinobot/sbot/route"
	"strconv"
	"strings"
	"time"
)

// 由于leveldb不支持list格式，只能把多个grouip存在一个byte slice中
// 或者用链式的结构？不过那样很乱
// 至于为啥用leveldb不用关系型，因为sql配置好麻烦，leveldb开箱即用！！
var (
	url           = "http://api.bilibili.com/x/space/acc/info"
	subscribeList map[string]([]float64)
	state         map[string]int

	db *datas.Db
)

const (
	ONSTREAMING int = 1
	OFFLINE     int = 0
)

func init() {
	db = datas.CreateDb("assets/bilibili")
	subscribeList = make(map[string][]float64)
	state = make(map[string]int)
	db.IterateAll(func(key, value string) {
		subscribeList[key] = ListToGroupId(value)
		state[key] = OFFLINE
	})
	logger.Info(subscribeList, state)
}
func ListToGroupId(list string) []float64 {
	flist := make([]float64, 0)
	gslist := strings.Split(list, ":")
	logger.Info("glist", gslist)
	for _, v := range gslist {
		id, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println(err)
		} else {
			flist = append(flist, id)
		}

	}
	return flist
}
func checkStates(d route.DataMap) {
	checkState()
}
func checkState() {

	request.SendMessage("jijijji?", 1012330112)
	var arl string
	for k, v := range state {

		arl = url + "?mid=" + k
		resp, err := http.Get(arl)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		logger.Info(string(body))

		var ui userInfo
		json.Unmarshal(body, &ui)
		status := ui.Data.LiveRoom.LiveStatus
		logger.Info("status: ", status)
		logger.Info("c:", v)
		if v^status == 1 {
			state[k] = status
			var msg string
			if status == ONSTREAMING {
				msg = "本群订阅主播: " + ui.Data.Name + " 已开播\n" +
					ui.Data.LiveRoom.Title + "\n" +
					"[CQ:image,file="+ui.Data.LiveRoom.Cover+"]" + "\n" +
					ui.Data.LiveRoom.URL
			} else if status == OFFLINE {
				msg = "本群订阅主播: " + ui.Data.Name + " 已下播\n"
			}
			for _, v := range subscribeList[k] {
				request.SendMessage(msg, v)
			}
		}
		timer := time.NewTimer(1 * time.Second)
		<-timer.C
		timer.Stop()
	}

}
func subscribe(d route.DataMap) {

	vmap := d["group_value"].(map[string](string))
	mid := vmap["mid"]

	if _, ok := subscribeList[mid]; ok {
		subscribeList[mid] = append(subscribeList[mid], d.GroupID())
	} else {
		subscribeList[mid] = make([]float64, 1)
		subscribeList[mid][0] = d.GroupID()
	}

	if _, ok := state[mid]; !ok {
		state[mid] = 0
	}

	ok, list := db.Get(mid)
	logger.Info("LIST", list)
	sid := strconv.FormatFloat(d.GroupID(), 'f', -1, 64)
	if ok {
		if !strings.Contains(list, sid) {
			list += (sid + ":")
			logger.Info("LIST", list)
			db.Put(mid, list)
		} else {
			request.SendMessage("不准重复订阅[CQ:face,id=11][CQ:face,id=11][CQ:face,id=11]", d.GroupID())
			return
		}
	} else {
		db.Put(mid, sid+":")
	}

	request.SendMessage("订阅成功", d.GroupID())
}
func unsubscribe(d route.DataMap) {

	vmap := d["group_value"].(map[string](string))
	mid := vmap["mid"]
	find := false
	if _, ok := subscribeList[mid]; ok {

		for i, v := range subscribeList[mid] {
			if v == d.GroupID() {
				find = true
				_, s := db.Get(mid)
				fmt.Println("S ",s)
				sl := strings.Split(s, strconv.FormatFloat(d.GroupID(), 'f', -1, 64)+":")
                fmt.Println("SL ",sl)
				if i == len(subscribeList[mid])-1 && i != 0 {
					subscribeList[mid] = subscribeList[mid][:len(subscribeList[mid])-1]
					db.Put(mid, sl[0])
				} else if i != 0 {
					subscribeList[mid] = append(subscribeList[mid][:i], subscribeList[mid][i+1:]...)
					db.Put(mid, sl[0]+sl[1])
				} else if i == 0 && len(subscribeList[mid]) != 1 {
					subscribeList[mid] = subscribeList[mid][1:]
					db.Put(mid, sl[1])
				} else if i == 0 && len(subscribeList[mid]) == 1 {
					delete(subscribeList, mid)
					delete(state, mid)
					db.Delete(mid)
				}
				logger.Info("subscribe list ",subscribeList[mid])
			}

		}
		if !find {
			request.SendMessage("找不到你[CQ:face,id=11][CQ:face,id=11][CQ:face,id=11]", d.GroupID())
		}
	} else {
		request.SendMessage("你就没订阅[CQ:face,id=11][CQ:face,id=11][CQ:face,id=11]", d.GroupID())
	}

}
