package bilibili

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shinobot/sbot/request"
	"shinobot/sbot/route"
	"shinobot/sbot/tick"
	"sync"
	"time"
)

var (
	url                = "http://api.bilibili.com/x/space/acc/info"
	a   map[string]int = make(map[string]int)

	mutex sync.Mutex
)

func getUerInfo(mid string) (status int) {

	arl := url + "?mid=" + mid
	resp, err := http.Get(arl)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var ui userInfo
	json.Unmarshal(body, &ui)
	fmt.Println(ui.Data.LiveRoom.LiveStatus)
	return ui.Data.LiveRoom.LiveStatus

}

func Subscribe(d route.DataMap) {

	vmap := d["group_value"].(map[string](string))
	mid := vmap["mid"]
	a[mid] = getUerInfo(mid)
	fmt.Println(a[mid])
	request.SendMessage(mid, d.GroupID())
}
func Tick(t *tick.Timer) {
	timer := time.NewTimer(2 * time.Second)
	<-timer.C
	mutex.Lock()
	for k, _ := range a {
		fmt.Println(k + "locking")
		request.SendMessage("room: "+k, 1012330112)
	}
	mutex.Unlock()
}
