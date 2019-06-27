package model

import (
	"sync"
	"time"
)

type BaseModel struct {
	Id        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreatedAt time.Time  `gorm:"column:create_time" json:"-"`
	UpdatedAt time.Time  `gorm:"column:update_time;null" json:"-"`
	DeletedAt *time.Time `gorm:"column:delete_time;null"  sql:"index" json:"-"`
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	SayHello  string `json:"sayHello"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ProductInfo struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Intro       string `json:"intro"`
	Label       string `json:"label"`
	Create_Time string `json:"create_time"`
	Update_Time string `json:"update_time"`
	State       int    `json:"state"`
	Pic         string `json:"pic"`
}

type TagInfo struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Source      string `json:"source"`
	Category    uint64 `json:"category"`
	Property    uint64 `json:"property"`
	Status      uint64 `json:"status"`
	Create_Time string `json:"create_time"`
	Update_Time string `json:"update_time"`
	Operator    string `json:"operator"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

type ProductList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*ProductInfo
}

type TagList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*TagInfo
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}
