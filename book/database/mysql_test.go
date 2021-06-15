package database

import (
	"log"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	db,err:=NewDatabase()
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()
}
