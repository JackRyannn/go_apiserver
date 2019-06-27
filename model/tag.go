package model

import (
	"apiserver/pkg/constvar"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

// 定义model及包含的字段
//注意这里首字母必须大写，否则匹配不上
type TagModel struct {
	BaseModel
	Name     string `gorm:"column:name"`
	Source   string `gorm:"column:source;default:'企鹅号'"`
	Category uint64 `gorm:"column:category"`
	Property uint64 `gorm:"column:property"`
	State    uint64 `gorm:"column:state"`
	//这里用指针才能修改CloseAt的值
	ClosedAt *time.Time `gorm:"column:close_time;null" json:"-"`
	Operator string     `gorm:"column:operator;default:'admin'"`
}

//设置该model从属的表名
func (*TagModel) TableName() string {
	return "tb_tag"
}

// Create creates a new user account.
func (u *TagModel) Create() error {
	return DB.Self.Create(&u).Error
}

func GetTagById(id string) (*TagModel, error) {
	p := new(TagModel)
	d := DB.Self.Where("id=?", id).First(&p)
	return p, d.Error
}

func ListTag(name string, offset, limit int) ([]*TagModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	tags := make([]*TagModel, 0)
	var count uint64
	//模糊查询，前后两个百分号
	where := fmt.Sprintf("name like '%%%s%%'", name)
	if err := DB.Self.Model(&TagModel{}).Where(where).Count(&count).Error; err != nil {
		return tags, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&tags).Error; err != nil {
		return tags, count, err
	}

	return tags, count, nil
}

// Validate the fields.
func (u *TagModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
