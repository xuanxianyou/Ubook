package utils

import (
	"fmt"
	"testing"
)

func TestVerifyCode_Code(t *testing.T) {
	v:=NewVerifyCode("17331987381")
	code:=v.Code()
	fmt.Println(code)
}

func TestSMSClient_SendMessage(t *testing.T) {
	client:=NewSMSClient()
	v:=NewVerifyCode("17331987381")
	code:=v.Code()
	client.SendMessage("17331987381",code)
}
