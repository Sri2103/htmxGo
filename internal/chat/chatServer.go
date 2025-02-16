package chat

type Message struct {
	Name  string `json:"name"`
	Text  string `json:"text"`
	Email string `json:"email"`
}

type chatHandler struct{}

func NewChatServer() *chatHandler {
	return &chatHandler{}
}
