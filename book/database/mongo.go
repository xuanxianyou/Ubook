package database

import (
	"GoProject/WebProject/MicroService/UserBook/book/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoDB struct {
	Client *mongo.Client
}

// ConnectMongo : 连接mongo数据库
func ConnectMongo() *MongoDB{
	conf:= config.Config.Mongo
	var MongoDBURI = fmt.Sprintf("%s://%s:%s",conf.Type, conf.Host, conf.Port)
	clientOptions := options.Client().ApplyURI(MongoDBURI)
	client,err:=mongo.Connect(context.TODO(),clientOptions)
	if err!=nil{
		log.Fatal(err)
	}
	err=client.Ping(context.TODO(),nil)
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("connect successfully")
	return &MongoDB{Client: client}
}

// DisconnectMongo : 关闭mongo数据库连接
func (m *MongoDB)DisconnectMongo(){
	err:=m.Client.Disconnect(context.TODO())
	if err!=nil{
		log.Fatal(err)
	}
	log.Println("Disconnect successfully")
}

