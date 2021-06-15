package sms

type TemplateMessage struct {
	Mobile []string `json:"mobile"`
	Type int64 `json:"type"`
	Sign string `json:"sign"`
	TemplateId string `json:"template_id"`
	SendTime string `json:"send_time"`
	Params map[string]interface{} `json:"params"`
}

type PlainMessage struct {
	Mobile   []string `json:"mobile"`
	Body     string   `json:"body"`
	SendTime string   `json:"send_time"`
}

type CancelMessage struct {
	MessageId string `json:"message_id"`
}
