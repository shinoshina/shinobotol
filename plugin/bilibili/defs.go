package bilibili

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shinobot/sbot/route"
)

var(
	url = "http://api.bilibili.com/x/space/acc/info"
)
func getUerInfo(d route.DataMap){

	arl := url+"?mid="+"2"
	resp,err := http.Get(arl)
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body,_ := ioutil.ReadAll(resp.Body)

	var ui userInfo
	json.Unmarshal(body,&ui)
	fmt.Println(ui.Data.LiveRoom.LiveStatus)

}