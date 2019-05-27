package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func GetDiaryText(date string) string {
	pathStr := "./service/diary_" + date + ".txt"
	absPath, _ := filepath.Abs(pathStr)
	f, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	return string(f)
}

func SetDiaryText(text string) string {
	date := time.Now().Format("20060102")
	// 生成日记文本为diary_yyyyMMdd.txt
	pathStr := "./service/diary_" + date + ".txt"
	absPath, _ := filepath.Abs(pathStr)

	f, err := os.Create(absPath)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	defer f.Close()

	n2, err := f.WriteString(text)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}

	fmt.Println(n2)

	if n2 > 0 {
		return "写日记成功"
	} else {
		return "写日记失败"
	}
}
