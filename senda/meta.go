package senda

import (
	"bytes"
	"encoding/json"
	"net/http"
)


type GroupMsg struct{
	Group_id int64 `json:"group_id"`
    Message string `json:"message"`
	Auto_escape bool `json:"auto_escape"`
}



func SendMessage(mes string,groupid float64){

	msg := GroupMsg{
		Group_id: int64(groupid),
		Message: mes,
		Auto_escape: false,
	}
	jsonmsg,_ := json.Marshal(msg)
	resp,err := http.Post("http://127.0.0.1:5700/send_group_msg","application/json",bytes.NewBuffer(jsonmsg))
    if err != nil{
		panic(err)
	}

	defer resp.Body.Close()


}