package model

import (
	"fmt"
	"testing"
)

func TestGetProductById(t *testing.T) {
	initConfigAndDatabase()
	defer DB.Close()
	p, err := GetProductById("1")
	if err != nil {
		fmt.Println("wrong")
		fmt.Println(err)
	}
	fmt.Println("chaoren=", p.Name, p.Author)
}

func TestListProduct(t *testing.T) {
	defer DB.Close()
	initConfigAndDatabase()
	a, count, err := ListProduct("霸道", 0, 0)
	fmt.Println(a[0].Name)
	if err != nil {
		t.Error(err)
	}
	if count <= 0 {
		t.Error("Products Count <= 0")
	}
	fmt.Println(count)
}
