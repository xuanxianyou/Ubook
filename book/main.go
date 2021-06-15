package main

import (
	"GoProject/WebProject/MicroService/UserBook/book/consul"
	"GoProject/WebProject/MicroService/UserBook/book/endpoints"
	"GoProject/WebProject/MicroService/UserBook/book/services"
	"GoProject/WebProject/MicroService/UserBook/book/transports"
	"context"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main(){
	// 服务注册基本信息
	consulHost := flag.String("consul.host", "localhost", "consul address")
	consulPort := flag.Int("consul.port", 8500, "consul port")
	name := flag.String("service.name", "book", "service name")
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
	// book 服务
	bookService := services.BookService{}

	bookEndpoints :=&endpoints.BookEndpoints{
		GetBookEndpoint: endpoints.MakeGetBookEndpoint(bookService),
	}

	r := transports.MakeHttpHandler(ctx, bookEndpoints)

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
