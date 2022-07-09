package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type GroupMsg struct {
	Group_id    int64  `json:"group_id"`
	Message     string `json:"message"`
	Auto_escape bool   `json:"auto_escape"`
}

type MessageId struct {
	Message_id int32 `json:"message_id"`
}

func SendMessage(mes string, groupid float64) {

	msg := GroupMsg{
		Group_id:    int64(groupid),
		Message:     mes,
		Auto_escape: false,
	}
	jsonmsg, _ := json.Marshal(msg)
	resp, err := http.Post("http://127.0.0.1:5700/send_group_msg", "application/json", bytes.NewBuffer(jsonmsg))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}

func GetMessage(id int32) string {

	mid := MessageId{
		Message_id: id,
	}
	jsonmsg, _ := json.Marshal(mid)
	resp, err := http.Post("http://127.0.0.1:5700/get_msg", "application/json", bytes.NewBuffer(jsonmsg))

	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body))
	msgRule := regexp.MustCompile(`"message":"(?P<raw_message>.*?)",`)
	result := msgRule.FindStringSubmatch(string(body))

	fmt.Println(result[1])
	return result[1]
}
