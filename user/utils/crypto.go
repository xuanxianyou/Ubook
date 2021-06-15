package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"
)

type Crypto struct {
	salt string
	password string
}

func NewCrypto(password string)*Crypto{
	return &Crypto{
		salt: newSalt(),
		password: password,
	}
}

func newSalt() string {
	const (
		randomChar string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	)
	rand.Seed(time.Now().UnixNano())
	var salt bytes.Buffer
	for i := 0; i < 64; i++ {
		salt.WriteByte(randomChar[rand.Int63()%int64(len(randomChar))])
	}
	return salt.String()
}

func (c *Crypto)Salt()string{
	return c.salt
}

func (c *Crypto)Encrypt()(string,error){
	if c.password=="" || c.salt ==""{
		return "",errors.New("null crypto")
	}
	hash := sha256.New()
	pwdHash := hash.Sum([]byte(c.password + c.salt))
	return hex.EncodeToString(pwdHash),nil
}


