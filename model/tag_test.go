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
	fmt.Println("chaoren=", p.Id, p.Name, p.CreatedAt, p.Source)
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

	temp_time := time.Now()

	tag := TagModel{
		Name:     "测试",
		Source:   "百家号",
		Category: 9,
		Property: 10,
		State:    0,
		ClosedAt: &temp_time,
		Operator: "admin",
	}
	print(time.Now().String())
	err := tag.Create()
	if err != nil {
		t.Error(err)
	} else {
		t.Log("创建成功")
	}

}

func TestTagModel_Update(t *testing.T) {
	defer DB.Close()
	initConfigAndDatabase()
	tagModel := TagModel{
		Name: "jay",
		BaseModel: BaseModel{
			Id: 0,
		},
		State: 3,
	}
	err := tagModel.Update(&tagModel)
	fmt.Println(err)
}
