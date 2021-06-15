package transports

import (
	"GoProject/WebProject/MicroService/UserBook/user/endpoints"
	"GoProject/WebProject/MicroService/UserBook/user/model"
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

func MakeHttpHandler(ctx context.Context,endpoint *endpoints.UserEndpoints)http.Handler{
	r:=mux.NewRouter()
	kitLog:=log.NewLogfmtLogger(os.Stderr)
	kitLog=log.With(kitLog,"ts",log.DefaultTimestampUTC)
	kitLog=log.With(kitLog, "caller", log.DefaultCaller)

	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(kitLog)),
	}

	r.Methods(http.MethodPost,http.MethodOptions).Path("/register").Handler(kithttp.NewServer(
		endpoint.RegisterEndpoint,
		decodeRegisterRequest,
		encodeJsonResponse,
		options...
	))
	r.Methods(http.MethodPost,http.MethodOptions).Path("/captcha").Handler(kithttp.NewServer(
		endpoint.CaptchaEndpoint,
		decodeCaptchaRequest,
		encodeJsonResponse,
	))
	r.Methods(http.MethodPost,http.MethodOptions).Path("/login").Handler(kithttp.NewServer(
		endpoint.LoginEndpoint,
		decodeLoginRequest,
		encodeJsonResponse,
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

func decodeRegisterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	username := r.FormValue("username")
	phone := r.FormValue("phone")
	password := r.FormValue("password")
	if username == "" || phone == "" || password == "" {
		return nil, ErrorRequest
	}
	return &endpoints.RegisterRequest{
		Username: username,
		Phone:    phone,
		Password: password,
	}, nil
}

func decodeCaptchaRequest(ctx context.Context,r *http.Request)(interface{},error){
	phone:=r.FormValue("phone")
	fmt.Println("phone",phone)
	if len(phone)!=11{
		return nil,ErrorRequest
	}
	return &endpoints.CaptchaRequest{Phone: phone},nil
}

func decodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	m := r.FormValue("mode")
	identity := r.FormValue("identity")
	voucher := r.FormValue("voucher")
	if m == "" || identity == "" || voucher == "" {
		return nil, ErrorRequest
	}
	mode,err:=strconv.Atoi(m)
	if err!=nil{
		return nil, ErrorRequest
	}
	return &endpoints.LoginRequest{
		Mode: int32(mode),
		Identity:    identity,
		Voucher: voucher,
	}, nil
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
