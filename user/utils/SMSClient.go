package utils

import (
	"GoProject/WebProject/MicroService/UserBook/user/config"
	"GoProject/WebProject/MicroService/UserBook/user/sms"
	"log"
)

type SMSClient struct {
	Assistant sms.Client
}

func NewSMSClient()*SMSClient{
	// 获取配置
	conf:=config.Config.Sms
	appId:=conf.AppId
	secretKey:=conf.SecretKey
	// 初始化客户端
	assistant:=sms.NewClient()
	assistant.SetAppId(appId)
	assistant.SetSecretKey(secretKey)


	return &SMSClient{
		Assistant: *assistant,
	}
}

func (client *SMSClient)SendMessage(phone string, code int32)bool{
	params := map[string]interface{}{"code":code}
	request := sms.NewRequest()
	request.SetMethod("sms.message.send")
	request.SetBizContent(sms.TemplateMessage{
		Mobile:     []string{phone},
		Type:       0,
		Sign:       "UBook",
		TemplateId: "ST_2020101100000007",
		SendTime:   "",
		Params:     params,
	})

	buf, err := client.Assistant.Execute(request)
	if err!=nil{
		log.Fatal(err)
		return false
	}
	log.Println(buf)
	return true
}
