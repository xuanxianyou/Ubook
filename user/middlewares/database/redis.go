package database

import (
	"GoProject/WebProject/MicroService/UserBook/user/config"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

type Redis struct {
	Connection redis.Conn
}

func NewRedis()*Redis{
	conf:=config.Config.Redis
	var dialect = fmt.Sprintf("%s://%s:%s",conf.Type,conf.Host,conf.Port)
	connection,err:=redis.DialURL(dialect)
	if err!=nil{
		log.Fatal(err)
	}
	return &Redis{Connection: connection}
}

func (r *Redis)Set(key string, value string, expired int)error{
	_,err :=r.Connection.Do("SET",key,value,"EX",expired)
	return err
}

func (r *Redis)Get(key string)(string,error){
	exist,err := redis.Bool(r.Connection.Do("EXISTS",key))
	if err != nil {
		return "",err
	}
	if !exist{
		return "",nil
	}
	value, err := redis.String(r.Connection.Do("GET", key))
	if err != nil {
		return "", err
	}
	return value,nil
}
