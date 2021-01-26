package client

type Message struct {
	Action  string   `json:"action"`
	Symbols []string `json:"symbols"`
}

type QuotesMessage struct {
	Timestamp int64   `json:"timestamp"`
	Symbol    string  `json:"symbol"`
	Price     float32 `json:"price"`
}
