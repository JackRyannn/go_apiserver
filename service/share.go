package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetShareText() string {

	absPath, _ := filepath.Abs("./service/shareText.txt")
	f, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	return string(f)
}

func SetShareText(text string) string {
	absPath, _ := filepath.Abs("./service/shareText.txt")

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
		return "设置成功"
	} else {
		return "设置失败"
	}
}
