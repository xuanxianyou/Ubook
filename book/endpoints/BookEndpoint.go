package endpoints

import(
	"GoProject/WebProject/MicroService/UserBook/book/model"
	"GoProject/WebProject/MicroService/UserBook/book/services"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type BookEndpoints struct {
	GetBookEndpoint endpoint.Endpoint
}

type GetBookRequest struct {
	UserId int64
}

type GetBookResponse struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Book model.Book `json:"book"`
}

func MakeGetBookEndpoint(service services.BookService)endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r:=request.(*GetBookRequest)
		response = service.GetBook(ctx, r.UserId)
		return
	}
}