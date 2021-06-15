package database

import (
	"GoProject/WebProject/MicroService/UserBook/book/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)


type Database struct {
	Connection *gorm.DB
}

func NewDatabase()(*Database,error){
	var(
		dialect,user,password,host,port,name string
	)

	conf := config.Config.Mysql
	dialect = conf.Type
	name = conf.Name
	user = conf.User
	password = conf.Password
	host = conf.Host
	port = conf.Port
	var source = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",user,password,host,port,name)
	conn,err := gorm.Open(dialect,source)
	if err!=nil{
		log.Fatal("connecting mysql error: ", err)
		return nil,err
	}
	log.Println("Connect Mysql Success")
	var db Database
	db.Connection = conn
	return &db,nil
}


func (db *Database)Close(){
	if err:= db.Connection.Close();err!=nil{
		log.Fatal(err)
	}
}
