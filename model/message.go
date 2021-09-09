package model

type Message struct {
	Channel []byte
	Text    []byte
}

type Dispatch struct {
	Token    []byte    `json:"bot_token"`
	Messages []Message `json:"channels"`
}
