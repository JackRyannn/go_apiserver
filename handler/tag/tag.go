package tag

import (
	"apiserver/model"
	"time"
)

type CreateRequest struct {
	Name       string    `json:"name"`
	Source     string    `json:"source"`
	Category   uint64    `json:"category"`
	Property   uint64    `json:"property"`
	State      uint64    `json:"state"`
	Close_Time time.Time `json:"close_time"`
	Operator   string    `json:"operator"`
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
