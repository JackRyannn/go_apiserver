package model

import (
	"apiserver/pkg/constvar"
	"fmt"
	"time"
)

// 定义model及包含的字段
//注意这里首字母必须大写，否则匹配不上
type ProductModel struct {
	Id          uint64
	Name        string
	Author      string
	Intro       string
	Label       string
	Create_Time time.Time
	Update_Time time.Time
	State       int
	Pic         string
}

//设置该model从属的表名
func (*ProductModel) TableName() string {
	return "tb_product"
}

func GetProductById(id string) (*ProductModel, error) {
	p := new(ProductModel)
	d := DB.Self.Where("id=?", id).First(&p)
	return p, d.Error
}

func ListProduct(name string, offset, limit int) ([]*ProductModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	products := make([]*ProductModel, 0)
	var count uint64
	//模糊查询，前后两个百分号
	where := fmt.Sprintf("name like '%%%s%%'", name)
	if err := DB.Self.Model(&ProductModel{}).Where(where).Count(&count).Error; err != nil {
		return products, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&products).Error; err != nil {
		return products, count, err
	}

	return products, count, nil
}
