package product

import (
	"apiserver/model"
)

type CreateRequest struct {
	Name string `json:"name"`
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
	TotalCount  uint64               `json:"totalCount"`
	ProductList []*model.ProductInfo `json:"productList"`
}
