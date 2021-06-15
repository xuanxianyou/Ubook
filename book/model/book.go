package model

import (
	"GoProject/WebProject/MicroService/UserBook/book/config"
	"GoProject/WebProject/MicroService/UserBook/book/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func init(){
	db,err:=database.NewDatabase()
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()
	//if !db.Connection.HasTable(Book{}){
	//	db.Connection.CreateTable(Book{})
	//}
}

type Book struct {
	// Id   int64  `json:"id" gorm:"primary_key;AUTO_INCREMENT"`// ID
	Name string `json:"name" gorm:"not null"`                // 书名
	Cover string `json:"cover" gorm:"not null"`              // 封面
	Author string `json:"author" gorm:"not null"`            // 作者
	Press string `json:"press" gorm:"not null"`              // 出版社
	Publication string `json:"publication"`                  // 出品方
	Page int `json:"page"`                                   // 页数
	Price float32 `json:"price"`                             // 价格
	Series string `json:"series"`                            // 丛书
	ISBN   string `json:"isbn"`                              // ISBN
	Score  float32 `json:"score"`                            // 评分
	Tags   []string `json:"tags"`                            // 标签
	ContentBrief  string `json:"content_brief"`              // 内容简介
	AuthorBrief string `json:"author_brief"`                 // 作者简介
	Catalogue	string `json:"catalogue"`                    //目录
}


func NewBook(name string,cover string,author string,press string,publication string,page int,price float32,series string,isbn string,score float32,tags []string,contentBrief string,authorBrief string,catalogue string)*Book{
	return &Book{
		Name:         name,
		Cover:        cover,
		Author:       author,
		Press:        press,
		Publication:  publication,
		Page:         page,
		Price:        price,
		Series:       series,
		ISBN:         isbn,
		Score:        score,
		Tags:         tags,
		ContentBrief: contentBrief,
		AuthorBrief:  authorBrief,
		Catalogue:    catalogue,
	}
}

//func (book *Book)CreateBook(){
//	db,err:=database.NewDatabase()
//	if err!=nil{
//		log.Fatal(err)
//	}
//	defer db.Close()
//	// 没有则创建表单
//	if db.Connection.HasTable(Book{}){
//		db.Connection.CreateTable(Book{})
//	}
//	db.Connection.Create(book)
//}
//
//func (book *Book)FindBook(key string,value string){
//	db,err:=database.NewDatabase()
//	if err!=nil{
//		log.Fatal(err)
//	}
//	defer db.Close()
//	db.Connection.Where(key + "= ?",value).Find(book)
//}
//
//func FindAllBooks()[]Book{
//	var books []Book
//	db,err:=database.NewDatabase()
//	if err!=nil{
//		log.Fatal(err)
//	}
//	defer db.Close()
//	db.Connection.Find(books)
//	return books
//}

func (book *Book)CreateBook()(insertID primitive.ObjectID){
	mongo:=database.ConnectMongo()
	defer mongo.DisconnectMongo()
	conf:= config.Config.Mongo
	collection:=mongo.Client.Database(conf.Database).Collection(conf.Collection)
	insertResult,err:=collection.InsertOne(context.TODO(),book)
	if err!=nil{
		log.Fatal(err)
	}
	if insertResult!=nil{
		insertID=insertResult.InsertedID.(primitive.ObjectID)
	}
	log.Println("Insert block successfully")
	return insertID
}

func QueryAllBook()[]*Book {
	mongo:=database.ConnectMongo()
	defer mongo.DisconnectMongo()
	conf:= config.Config.Mongo
	filter:=bson.D{}
	collection:=mongo.Client.Database(conf.Database).Collection(conf.Collection)
	cursor,err:=collection.Find(context.TODO(),filter)
	if cursor!=nil && cursor.Err()!=nil{
		log.Panicf("Query mongo Error:%v",cursor.Err().Error())
	}
	var books []*Book
	if cursor!=nil{
		err = cursor.All(context.TODO(),&books)
		if err!=nil{
			log.Fatal(err)
		}
	}
	return books
}

