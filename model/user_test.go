package model

import (
	"apiserver/config"
	"fmt"
	"testing"
)

var (
	cfg = "../conf/config.yaml"
)

func initConfigAndDatabase() {
	// init config 这里要设置配置文件路径，根据你的test包的位置要调整相对路径的
	if err := config.Init(cfg); err != nil {
		panic(err)
	}
	//初始化数据库
	DB.Init()
}
func TestGetUser(t *testing.T) {
	defer DB.Close()
	initConfigAndDatabase()
	userModel, err := GetUser("admin")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(userModel.Username, userModel.Password)
}

func TestListUser(t *testing.T) {
	defer DB.Close()
	initConfigAndDatabase()
	_, count, err := ListUser("admin", 0, 0)
	if err != nil {
		t.Error(err)
	}
	if count <= 0 {
		t.Error("ListUser Count <= 0")
	}
	fmt.Println(count)
}

func TestUserModel_Update(t *testing.T) {
	defer DB.Close()
	initConfigAndDatabase()
	userModel := UserModel{
		Username: "admin",
	}
	userModel.Id = 0
	err := userModel.Update(&userModel)
	fmt.Println(err)
}
