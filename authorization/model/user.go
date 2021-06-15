package model

type User struct {
	UserId int64 // user 标识

	Username string

	Password string

	Permissions []string // user 权限
}
