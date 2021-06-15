package main

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/hashicorp/consul/api"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main(){
	var(
		consulHost = flag.String("host","localhost","consul server ip address")
		consulPort = flag.String("port","8500","consul server port")
	)
	flag.Parse()

	//创建日志组件
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// 创建consul client
	consulConfig := api.DefaultConfig()
	consulConfig.Address = "http://" + *consulHost + ":" + *consulPort
	consulClient, err := api.NewClient(consulConfig)
	if err!=nil{
		_ = logger.Log("error", err)
		os.Exit(1)
	}

	// 创建反向代理
	proxy:=NewReverseProxy(consulClient,logger)

	errChan := make(chan error,1)
	go func() {
		c:=make(chan os.Signal)
		signal.Notify(c,syscall.SIGINT,syscall.SIGTERM)
		errChan<- fmt.Errorf("%s",<-c)
	}()

	// 监听
	go func() {
		_ = logger.Log("transport", "HTTP", "addr", "9000")
		errChan<-http.ListenAndServe(":9000",proxy)
	}()

	_ = logger.Log("exit", <-errChan)
}

func NewReverseProxy(client *api.Client, logger log.Logger)*httputil.ReverseProxy{
	//创建Director
	director := func(req *http.Request) {

		//查询原始请求路径
		reqPath := req.URL.Path
		if reqPath == "" {
			return
		}
		//按照分隔符'/'对路径进行分解，获取服务名称serviceName
		pathArray := strings.Split(reqPath, "/")
		serviceName := pathArray[1]

		//调用consul api查询serviceName的服务实例列表
		result, _, err := client.Catalog().Service(serviceName, "", nil)
		if err != nil {
			_ = logger.Log("ReverseProxy failed", "query service instance error", err.Error())
			return
		}

		if len(result) == 0 {
			_ = logger.Log("ReverseProxy failed", "no such service instance", serviceName)
			return
		}

		//重新组织请求路径，去掉服务名称部分
		destPath := strings.Join(pathArray[2:], "/")

		//随机选择一个服务实例
		rand.Seed(time.Now().UnixNano())
		tgt := result[rand.Int()%len(result)]
		_ = logger.Log("service id", tgt.ServiceID)

		//设置代理服务地址信息
		req.URL.Scheme = "http"
		req.URL.Host = fmt.Sprintf("%s:%d", tgt.ServiceAddress, tgt.ServicePort)
		req.URL.Path = "/" + destPath
	}
	proxy:=httputil.ReverseProxy{Director: director}
	return &proxy
}
