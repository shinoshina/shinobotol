package route


type Message struct {
	mType string
	raw   string
}

func NewMessage(t string, r string) Message {
	m := Message{
		mType: t,
		raw:   r,
	}
	return m
}
