package requester

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

type MsgPostAll struct {
	Message_type string `json:"message_type"`
	User_id      int64  `json:"user_id"`
	Group_id     int64  `json:"group_id"`
	Message      string `json:"message"`
	Auto_escape  bool   `json:"auto_escapq"`
}

type MsgPostPrivate struct {
	User_id     int64  `json:"user_id"`
	Group_id    int64  `json:"group_id"`
	Message     string `json:"message"`
	Auto_escape bool   `json:"auto_escape"`
}

type MsgPostGroup struct {
	Group_id    int64  `json:"group_id"`
	Message     string `json:"message"`
	Auto_escape bool   `json:"auto_escape"`
}

type RequestPoster struct {
	Client *http.Client
}

func (rp RequestPoster) PostPoke(mes map[string]interface{}) {
	var group_id int64 = int64(mes["group_id"].(float64))
	id := strconv.FormatInt(int64(mes["sender_id"].(float64)), 10)

	if id != "2037310389" {
		poke := "[CQ:poke,qq=" + id + "]"
		msg := MsgPostGroup{
			Group_id:    group_id,
			Message:     poke,
			Auto_escape: false,
		}
		jsonStr, _ := json.Marshal(msg)
		resp, err := rp.Client.Post("http://127.0.0.1:5700/send_msg", "application/json", bytes.NewBuffer(jsonStr))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	}
}
