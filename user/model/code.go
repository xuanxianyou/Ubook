package model

// Code 验证码，存储在Redis数据库中
type Code struct {
	Phone string `json:"phone"`
	Captcha int32 `json:"captcha"`
}


