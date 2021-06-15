package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

//Config 定义配置
var (
	Config *Conf
)

//Conf 配置结构体
type Conf struct {
	Server   Server   `yaml:"server"`
	Mysql 	 Mysql    `yaml:"mysql"`
	Sms      SMS	  `yaml:"sms"`
	Redis    Redis    `yaml:"redis"`
}

//Server HTTP服务配置结构体
type Server struct {
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read-timeout"`
	WriteTimeout time.Duration `yaml:"write-timeout"`
	SecretKey    string        `yaml:"secret-key"`
}

//Database 数据库配置结构体
type Mysql struct {
	Type        string `yaml:"type"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Name        string `yaml:"name"`
	TablePrefix string `yaml:"table-prefix"`
}

//
type SMS struct {
	AppId     string `yaml:"app-id"`
	SecretKey string `yaml:"secret-key"`
}

type Redis struct {
	Type string `yaml:"type"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

//init 初始化函数
func init() {
	Config = getConf()
	log.Println("[Setting] Config init success")
}

//getConf 读取配置文件
func getConf() *Conf {
	var c *Conf
	file, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Println("[Setting] config error: ", err)
	}
	err = yaml.UnmarshalStrict(file, &c)
	if err != nil {
		log.Println("[Setting] yaml unmarshal error: ", err)
	}
	return c
}
