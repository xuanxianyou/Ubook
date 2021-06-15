package main

import (
	"GoProject/WebProject/MicroService/UserBook/authorization/config"
	"GoProject/WebProject/MicroService/UserBook/authorization/consul"
	"GoProject/WebProject/MicroService/UserBook/authorization/endpoints"
	"GoProject/WebProject/MicroService/UserBook/authorization/model"
	"GoProject/WebProject/MicroService/UserBook/authorization/services"
	"GoProject/WebProject/MicroService/UserBook/authorization/transports"
	"context"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {

	// 服务注册基本信息
	consulHost := flag.String("consul.host", "localhost", "consul address")
	consulPort := flag.Int("consul.port", 8500, "consul port")
	name := flag.String("service.name", "authorization", "service name")
	address := flag.String("service.address", "localhost", "service address")
	port := flag.Int("service.port", 8000, "service port")
	tag := flag.String("service.tag", "primary", "service tag")
	flag.Parse()

	Id := *name + "-" +  uuid.New().String()
	// 生成consul客户端实例
	consulAddress := fmt.Sprintf("%s:%d", *consulHost, *consulPort)
	client := consul.NewClient(consulAddress)
	ctx := context.Background()
	errChan := make(chan error)

	var tokenService services.TokenService
	var tokenGranter services.TokenGranter
	var tokenEnhancer services.TokenEnhancer
	var tokenStore services.TokenStore
	var userDetailsService services.UserService
	var clientDetailsService services.ClientDetailsService
	var srv services.Service


	tokenEnhancer = services.NewJwtTokenEnhancer("secret")
	tokenStore = services.NewJwtTokenStore(tokenEnhancer.(*services.JwtTokenEnhancer))
	tokenService = services.NewTokenService(tokenStore, tokenEnhancer)

	userDetailsService = services.NewInMemoryUserService([] *model.User{
		{
		Username:    "aoho",
		Password:    "123456",
		UserId:      1,
		Permissions: []string{"Simple"},
		},
		{
			Username:    "admin",
			Password:    "123456",
			UserId:      1,
			Permissions: []string{"Admin"},
		},
	})
	clientDetailsService = services.NewInMemoryClientDetailService([] *model.Client{{
		"clientId",
		"clientSecret",
		1800,
		18000,
		"http://127.0.0.1",
		[] string{"password", "refresh_token"},
	}})

	tokenGranter = services.NewComposeTokenGranter(map[string]services.TokenGranter{
		"password": services.NewUsernamePasswordTokenGranter("password", userDetailsService,  tokenService),
		"refresh_token": services.NewRefreshGranter("refresh_token", userDetailsService,  tokenService),

	})

	// token endpoint
	tokenEndpoint := endpoint.MakeTokenEndpoint(tokenGranter, clientDetailsService)
	// log 封装
	tokenEndpoint = endpoint.MakeClientAuthorizationMiddleware(config.KitLogger)(tokenEndpoint)
	// check token endpoint
	checkTokenEndpoint := endpoint.MakeCheckTokenEndpoint(tokenService)
	// log 封装
	checkTokenEndpoint = endpoint.MakeClientAuthorizationMiddleware(config.KitLogger)(checkTokenEndpoint)

	srv = services.NewCommonService()

	//创建健康检查的Endpoint
	healthEndpoint := endpoint.MakeHealthCheckEndpoint(srv)

	endpts := endpoint.OAuth2Endpoints{
		TokenEndpoint:tokenEndpoint,
		CheckTokenEndpoint:checkTokenEndpoint,
		HealthCheckEndpoint: healthEndpoint,
	}

	//创建http.Handler
	r := transport.MakeHttpHandler(ctx, endpts, tokenService, clientDetailsService, config.KitLogger)

	go func() {
		client.RegisterService(Id, *name, *address, *port, *tag)
		config.Logger.Println("Http Server start at port:" + strconv.Itoa(*port))
		handler := r
		errChan <- http.ListenAndServe(":"  + strconv.Itoa(*port), handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	err := <-errChan
	client.UnRegisterService(Id)
	config.Logger.Println(err)
}
