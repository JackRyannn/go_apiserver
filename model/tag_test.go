package model

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTagById(t *testing.T) {
	initConfigAndDatabase()
	defer DB.Close()
	p, err := GetTagById("1")
	if err != nil {
		fmt.Println("wrong")
		fmt.Println(err)
	}
	fmt.Println("chaoren=", p.Id, p.Name, p.Create_Time, p.Source)
}

func TestListTag(t *testing.T) {
	defer DB.Close()
	initConfigAndDatabase()
	a, count, err := ListTag("", 0, 0)
	fmt.Println(a[0].Name)
	if err != nil {
		t.Error(err)
	}
	if count <= 0 {
		t.Error("Tags Count <= 0")
	}
	fmt.Println(count)
}

func TestCreateTag(t *testing.T) {
	defer DB.Close()
	initConfigAndDatabase()
	tag := TagModel{
		Name:        "测试",
		Source:      "百家号",
		Category:    9,
		Property:    10,
		State:       0,
		Create_Time: time.Now(),
		Update_Time: time.Now(),

		Operator: "admin",
	}
	print(time.Now().String())
	print(tag.Close_Time.String())
	err := tag.Create()
	if err != nil {
		t.Error(err)
	} else {
		t.Log("创建成功")
	}

}
