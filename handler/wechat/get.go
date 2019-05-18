package wechat

import (
	. "apiserver/handler"
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
)

import "github.com/gin-gonic/gin"

const (
	token = "chaoren" //设置token
)

func makeSignature(timestamp, nonce string) string { //本地计算signature
	si := []string{token, timestamp, nonce}
	sort.Strings(si)            //字典序排序
	str := strings.Join(si, "") //组合字符串
	s := sha1.New()             //返回一个新的使用SHA1校验的hash.Hash接口
	io.WriteString(s, str)      //WriteString函数将字符串数组str中的内容写入到s中
	return fmt.Sprintf("%x", s.Sum(nil))
}

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	signature := c.Query("signature")
	echostr := c.Query("echostr")

	signatureGen := makeSignature(timestamp, nonce)
	if signature != signatureGen {
		SendResponse(c, nil, false)
	} else {
		SendResponse(c, nil, echostr)
	}

}
