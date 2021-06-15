package endpoints

import (
	"GoProject/WebProject/MicroService/UserBook/user/services"
	"context"
	"github.com/go-kit/kit/endpoint"
)


type UserEndpoints struct {
	RegisterEndpoint endpoint.Endpoint
	CaptchaEndpoint endpoint.Endpoint
	LoginEndpoint endpoint.Endpoint
}

// 注册部分
type RegisterRequest struct {
	Username string
	Phone string
	Password string
}

type RegisterResponse struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string `json:"message"`
}

func MakeRegisterEndpoint(userService services.UserService) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r:=request.(*RegisterRequest)
		response = userService.Register(ctx, r.Username, r.Phone, r.Password)
		return
	}
}

// 验证码部分
type CaptchaRequest struct {
	Phone string
}

type CaptchaResponse struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string `json:"message"`
}

func MakeCaptchaEndpoint(userService services.UserService) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r:=request.(*CaptchaRequest)
		response = userService.Captcha(ctx,r.Phone)
		return
	}
}

// 登录部分
type LoginRequest struct {
	Mode 	 int32	// 登录方式，0->账号，1->手机号
	Identity string	// 身份，账号或手机号
	Voucher  string // 凭证，密码或验证码
}

type LoginResponse struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string `json:"message"`
}

func MakeLoginEndpoint(userService services.UserService) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r:=request.(*LoginRequest)
		response = userService.Login(ctx,r.Mode,r.Identity,r.Voucher)
		return
	}
}


