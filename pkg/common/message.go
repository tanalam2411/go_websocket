package common

type Message struct {
	Type int    `json:"type"`
	Body interface{} `json:"body"`
}

type LoginData struct {
	Username string `json:"username"`
	Token string `json:"token"`
}


type ResponseData struct {
	Topic string `json:"topic"`
	Message map[string]int `json:"message"`
}
