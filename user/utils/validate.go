package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

type Validate struct {
	salt string
	passwordEncrypted string
}

func NewValidate(salt string,passwordEncrypted string)*Validate{
	return &Validate{
		salt:              salt,
		passwordEncrypted: passwordEncrypted,
	}
}

func (v *Validate)Verify(password string)bool{
	hash := sha256.New()
	pwdHash := hash.Sum([]byte(password + v.salt))
	return hex.EncodeToString(pwdHash)==v.passwordEncrypted
}
