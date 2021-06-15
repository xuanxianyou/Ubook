package consul

import (
	"fmt"
	consulAPI "github.com/hashicorp/consul/api"
	"log"
)


type Client struct {
	ConsulClient *consulAPI.Client
}

func NewClient(address string)*Client{
	config := consulAPI.DefaultConfig()
	config.Address = address
	client,err:=consulAPI.NewClient(config)
	if err!=nil{
		log.Fatal(err)
	}
	return &Client{
		ConsulClient: client,
	}
}
// 服务注册
func (c *Client)RegisterService(Id string, name string, address string, port int, tag string){
	// 服务实例信息
	register := consulAPI.AgentServiceRegistration{}
	register.ID = Id
	register.Name = name
	register.Address = address
	register.Port = port
	register.Tags = []string{tag}
	// 健康检查信息
	check := consulAPI.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP	= fmt.Sprintf("http://host.docker.internal:%d/health",port)  // 对与宿主机和容器来说，从容器内访问宿主机端口通过host.docker.internal而不是localhost
	register.Check = &check
	// 注册服务
	err:=c.ConsulClient.Agent().ServiceRegister(&register)
	if err!=nil{
		log.Fatal(err)
	}
}

// 反注册
func (c *Client)UnRegisterService(Id string){
	_ = c.ConsulClient.Agent().ServiceDeregister(Id)
}


