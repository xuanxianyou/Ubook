package services

import (
	"GoProject/WebProject/MicroService/UserBook/user/database"
	"GoProject/WebProject/MicroService/UserBook/user/model"
	"GoProject/WebProject/MicroService/UserBook/user/utils"
	"context"
	"log"
	"net/http"
	"strconv"
)

const (
	SuccessMessage="注册成功！"
	UserExistMessage="用户已存在！"
)

type Response struct {
	Code int `json:"code"`
	Success bool `json:"success"`
	Message string `json:"message"`
}

type UserServiceInterface interface {
	Register(ctx context.Context,username string,phone string,password string)Response
	Captcha(ctx context.Context,phone string)Response
	Login(ctx context.Context,username string,password string)Response
}

type UserService struct {

}


func (userService *UserService)Register(ctx context.Context,username string,phone string,password string)Response{
	// 判断用户是否已经存在
	var user model.User
	user.FindUser("phone",phone)
	if user.Phone!=""{
		return Response{
			Code:    http.StatusOK,
			Success: false,
			Message: UserExistMessage,
		}
	}
	// 生成salt
	crypto:=utils.NewCrypto(password)
	password,err := crypto.Encrypt()
	if err!=nil{
		log.Fatal(err)
	}
	salt:=crypto.Salt()
	// 创建用户
	user=model.User{
		Username: username,
		Phone:    phone,
		Password: password,
		Salt:     salt,
	}
	// 保存用户
	user.CreateUser()
	return Response{
		Code:    http.StatusOK,
		Success: true,
		Message: SuccessMessage,
	}

}


func (userService *UserService)Captcha(ctx context.Context,phone string)Response{
	// 从数据库中获取验证码
	r:=database.NewRedis()
	client:=utils.NewSMSClient()
	c, err := r.Get(phone)
	if err!=nil{
		log.Fatal(err)
	}
	if c!=""{
		code,_:=strconv.Atoi(c)
		if client.SendMessage(phone, int32(code)){
			return Response{
				Code:    http.StatusOK,
				Success: true,
				Message: "验证码已发送！",
			}
		}
	}
	// 没有则向其发送
	v:=utils.NewVerifyCode(phone)
	code:=v.Code()
	if client.SendMessage(phone,code){
		c:=strconv.Itoa(int(code))
		err:=r.Set(phone,c,15*60)
		if err!=nil{
			log.Fatal(err)
		}
		return Response{
			Code:    http.StatusOK,
			Success: true,
			Message: "发送成功！",
		}
	}
	return Response{
		Code:    http.StatusOK,
		Success: false,
		Message: "发送失败！",
	}

}

func (userService *UserService)Login(ctx context.Context,mode int32,identity string,voucher string)Response{
	// 查找用户信息
	// 验证用户信息
	// 返回验证结果
	if mode==0{
		c:=NewLoginContext(LoginWithCode)
		if c.execute(identity,voucher){
			return Response{
				Code:    http.StatusOK,
				Success: true,
				Message: "验证成功！",
			}
		}else{
			return Response{
				Code:    http.StatusOK,
				Success: false,
				Message: "验证失败！",
			}
		}
	}
	if mode==1{
		c:=NewLoginContext(LoginWithPwd)
		if c.execute(identity,voucher){
			return Response{
				Code:    http.StatusOK,
				Success: true,
				Message: "验证成功！",
			}
		}else{
			return Response{
				Code:    http.StatusOK,
				Success: false,
				Message: "验证失败！",
			}
		}
	}
	return Response{
		Code: http.StatusOK,
		Success: false,
		Message: "验证成功！",
	}
}

type LoginContext struct {
	strategy func(string,string)bool
}

func NewLoginContext(strategy func(string,string)bool) *LoginContext{
	return &LoginContext{
		strategy: strategy,
	}
}

func (context *LoginContext)execute(identity string,voucher string)bool{
	return context.strategy(identity,voucher)
}

func LoginWithPwd(identity string,voucher string)bool{
	db,err:=database.NewDatabase()
	if err!=nil{
		return false
	}
	defer db.Close()
	// 判断username/phone
	// 查询用户信息
	user:=model.User{}
	user.FindUser("phone",identity)
	user.FindUser("username",identity)
	// 验证密码正确
	validate:=utils.NewValidate(user.Salt,user.Password)
	res:=validate.Verify(voucher)
	// 返回验证结果
	return res
}

func LoginWithCode(identity string,voucher string)bool{
	// 从redis数据库查询信息
	r:=database.NewRedis()
	c, err := r.Get(identity)
	if err!=nil{
		log.Fatal(err)
	}
	// 验证验证码的正确性
	return c==voucher
}