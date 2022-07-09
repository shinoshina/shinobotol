package route

import "fmt"



type MessageSet map[string]func(string)()

func DefaultHandler(s string){
	fmt.Println("hai hai hai!")
}
func NewMessageSet()(m MessageSet){
	m = make(MessageSet)
	m["/"] = DefaultHandler
	return
}
func(m MessageSet) On_Message(raw string,handler func(string)()){
	m[raw] = handler
}

