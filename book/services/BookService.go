package services

import (
	"GoProject/WebProject/MicroService/UserBook/book/model"
	"context"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Book model.Book `json:"book"`
}

type BookServiceInterface interface {
	GetBook(ctx context.Context,userId int64)Response
}

type BookService struct {

}

func (bookService *BookService)GetBook(ctx context.Context,userId int64)Response{
	books:=model.QueryAllBook()
	rand.Seed(time.Now().UnixNano())
	book:=books[(rand.Int()+int(userId))%len(books)]
	return Response{
		Code:    http.StatusOK,
		Success: true,
		Book:    *book,
	}
}