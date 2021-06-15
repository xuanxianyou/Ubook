package transports

import (
	"GoProject/WebProject/MicroService/UserBook/book/endpoints"
	"GoProject/WebProject/MicroService/UserBook/book/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

// UserTransport implement http protocol
var (
	ErrorRequest = errors.New("bad request")
)

func MakeHttpHandler(ctx context.Context,endpoint *endpoints.BookEndpoints)http.Handler{
	r:=mux.NewRouter()
	kitLog:=log.NewLogfmtLogger(os.Stderr)
	kitLog=log.With(kitLog,"ts",log.DefaultTimestampUTC)
	kitLog=log.With(kitLog, "caller", log.DefaultCaller)

	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(kitLog)),
	}

	r.Methods(http.MethodGet,http.MethodOptions).Path("/getbook").Handler(kithttp.NewServer(
		endpoint.GetBookEndpoint,
		decodeGetBookRequest,
		encodeJsonResponse,
		options...
	))

	r.Methods(http.MethodGet).Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		writer.Header().Set("Content-Type","application/json;charset=utf-8")
		_, _ = writer.Write([]byte(`{status:"ok"}`))
	})

	r.Use(mux.CORSMethodMiddleware(r))
	return r
}

func decodeGetBookRequest(ctx context.Context,r *http.Request)(interface{},error){
	params:=r.URL.Query()
	userId := params.Get("userId")
	fmt.Println("userId",userId)
	id,err:=strconv.Atoi(userId)
	if err!=nil{
		return nil, err
	}
	return &endpoints.GetBookRequest{UserId: int64(id)},nil
}

func encodeJsonResponse(ctx context.Context,w http.ResponseWriter, response interface{})error{
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Content-Type","application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(ctx context.Context,err error, w http.ResponseWriter){
	contentType,body:= "text/plain;charset=utf-8",[]byte(err.Error())
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type",contentType)
	if e,ok:=err.(*model.ServerError);ok{
		w.WriteHeader(e.Code)
		_, _ = w.Write(body)
	}else{
		w.WriteHeader(500)
		_, _ = w.Write(body)
	}
}
