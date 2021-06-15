package main

import (
	"GoProject/WebProject/MicroService/UserBook/user/consul"
	"GoProject/WebProject/MicroService/UserBook/user/endpoints"
	"GoProject/WebProject/MicroService/UserBook/user/middlewares"
	"GoProject/WebProject/MicroService/UserBook/user/services"
	"GoProject/WebProject/MicroService/UserBook/user/transports"
	"context"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/time/rate"
	"log"
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
	name := flag.String("service.name", "user", "service name")
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
	// user 服务
	userService := services.UserService{}
	// 中间件封装
	limiter:=rate.NewLimiter(1,5)
	userEndpoints := &endpoints.UserEndpoints{
		RegisterEndpoint: middlewares.RateLimit(limiter)(endpoints.MakeRegisterEndpoint(userService)),
		CaptchaEndpoint:  middlewares.RateLimit(limiter)(endpoints.MakeCaptchaEndpoint(userService)),
		LoginEndpoint:    middlewares.RateLimit(limiter)(endpoints.MakeLoginEndpoint(userService)),
	}

	r := transports.MakeHttpHandler(ctx, userEndpoints)

	// 启动并注册服务
	go func() {
		client.RegisterService(Id, *name, *address, *port, *tag)
		errChan <- http.ListenAndServe(":"+strconv.Itoa(*port), r)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	// 收集错误，注销服务
	err := <-errChan
	client.UnRegisterService(Id)
	log.Fatal(err)
}
