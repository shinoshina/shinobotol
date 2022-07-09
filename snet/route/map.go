package route

import "fmt"

type (
	DataMap map[string]interface{}

	MessageSet struct {
		m    map[string]func(d DataMap)
		mode map[string]string
	}

	ActivitySet map[string]func(d DataMap)
)

func DefaultHandler(d DataMap) {
	fmt.Println("hai hai hai!")
}
func NewMessageSet() (ms* MessageSet) {
	ms = new(MessageSet)
	ms.m = make(map[string]func(d DataMap))
	ms.mode = make(map[string]string)
	ms.m["/"] = DefaultHandler
	return
}
func (ms MessageSet) OnMessage(rm Message, handler func(d DataMap)) {
	ms.m[rm.raw] = handler
}
func test() {

	m := NewMessageSet()
	m.OnMessage(NewMessage("type", "nimasile"), func(d DataMap) {

	})
}
func NewActivitySet() (a ActivitySet) {
	a = make(ActivitySet)
	return
}
func (a ActivitySet) OnEvent(tp string, handler func(d DataMap)) {
	a[tp] = handler
}
