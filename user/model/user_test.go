package model

import (
	"fmt"
	"testing"
)

func TestNewUser(t *testing.T) {
	user:=NewUser("kaneziki","17331987381","aassas","1234566")
	fmt.Println(user)
}
