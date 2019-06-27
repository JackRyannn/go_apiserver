package tag

import (
	"apiserver/model"
)

const (
	STATE_DOWN = 0
	STATE_UP   = 1
	STATE_DEL  = 2
)

type CreateRequest struct {
	//接口传来的值都是string类型，再自己转换
	Name     string `json:"name"`
	Source   string `json:"source"`
	Category string `json:"category"`
	Property string `json:"property"`
	State    string `json:"state"`
	ClosedAt string `json:"close_time"`
	Operator string `json:"operator"`
}

type CreateResponse struct {
	Name string `json:"name"`
}

//定义List方法的入参格式
type ListRequest struct {
	Name   string `json:"name"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

//定义作品的返回格式 json类型
type ListResponse struct {
	TotalCount uint64           `json:"totalCount"`
	TagList    []*model.TagInfo `json:"tagList"`
}
