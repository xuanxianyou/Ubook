package database

import (
	"fmt"
	"log"
	"testing"
)

func TestNewRedis(t *testing.T) {
	_ = NewRedis()
}

func TestRedis_Set(t *testing.T) {
	r:=NewRedis()
	err:=r.Set("17331987381","0000",10)
	if err!=nil{
		log.Fatal(err)
	}
	code,err:=r.Get("17331987381")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(code)
}

func TestRedis_Get(t *testing.T) {
	r:=NewRedis()
	code,err:=r.Get("17331987381")
	if err!=nil{
		log.Fatal(err)
	}
	if code==""{
		fmt.Println("code不存在")
	}else{
		fmt.Println(code)
	}
}
