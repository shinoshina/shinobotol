package route

import "fmt"

type (

	DataMap map[string]interface{}

	MessageSet map[string]func(d DataMap)

	Activity map[string]func(d DataMap)
)

func DefaultHandler(d DataMap) {
	fmt.Println("hai hai hai!")
}
func NewMessageSet() (m MessageSet) {
	m = make(MessageSet)
	m["/"] = DefaultHandler
	return
}
func (m MessageSet) OnMessage(raw string, handler func(d DataMap)) {
	m[raw] = handler
}
