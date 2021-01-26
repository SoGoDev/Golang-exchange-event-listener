package bitmex

type Data struct {
	Symbol    string  `json:"symbol"`
	LastPrice float32 `json:"lastPrice"`
}

type Message struct {
	Action string `json:"action"`
	Data   []Data `json:"data"`
}
