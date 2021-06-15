package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	var conf = Config
	fmt.Println(conf.Server)
	fmt.Println(conf.Database)
	fmt.Println(conf.Sms)
}
