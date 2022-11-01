package entities

type Message struct {
	Sender   string `form:"sender"`
	Receiver string `form:"receiver"`
	Message  string `form:"message"`
}
