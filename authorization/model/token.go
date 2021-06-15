package model

import (
	"time"
)

type Token struct {
	RefreshToken *Token
	TokenType string
	TokenValue string
	ExpiredAt *time.Time
}

func (token *Token)IsExpired()bool{
	return token.ExpiredAt!=nil && token.ExpiredAt.Before(time.Now())
}

// 为授权服务器提供一致的方式获取客户端或者用户信息
type OAuth2Details struct {
	Client Client
	User   User
}