package utils

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

type VerifyCode struct {
	Phone string
}

func NewVerifyCode(phone string)*VerifyCode{
	return &VerifyCode{
		Phone: phone,
	}
}

func (v *VerifyCode)Code()int32{
	rand.Seed(time.Now().UnixNano())
	phone, err := strconv.Atoi(v.Phone)
	if err!=nil{
		log.Fatal(err)
	}
	randCode := int64(phone) + rand.Int63()
	code:=randCode%10000
	return int32(code)
}
