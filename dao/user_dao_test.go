package dao

import (
	"log"
	"testing"
)

func TestGetUserBasicInfoByIdt(t *testing.T) {
	res, err := GetUserBasicInfoById(1)
	if err == nil {
		log.Println(res)
	}
}

func TestGetUserBasicInfoByName(t *testing.T) {
	res, err := GetUserBasicInfoByName("unit test")
	if err == nil {
		log.Println(res)
	}
}

func TestInsertUser(t *testing.T) {
	user := UserBasicInfo{Name: "unit test", Password: "unit test"}
	flag := InsertUser(&user)
	log.Println(flag)
}
