package model

import (
	"GoProject/WebProject/MicroService/UserBook/user/database"
	"log"
)

type User struct {
	Id       int64 `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username string `json:"username" gorm:"unique not null"`
	Phone    string `json:"phone" gorm:"unique not null"`
	Password string `json:"-" gorm:"not null"`
	Salt     string `json:"-" gorm:"not null"`
}

func init(){
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// 没有则创建表单
	if !db.Connection.HasTable(User{}) {
		db.Connection.CreateTable(User{})
	}
}

func NewUser(username string, phone string, password string, salt string) *User {
	return &User{
		Username: username,
		Phone:    phone,
		Password: password,
		Salt:     salt,
	}
}

func (user *User) CreateUser() {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// 没有则创建表单
	if !db.Connection.HasTable(User{}) {
		db.Connection.CreateTable(User{})
	}
	db.Connection.Create(user)
}

func (user *User) FindUser(key string, value string) {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Connection.Where(key+"=?", value).Find(user)
}
