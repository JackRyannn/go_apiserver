package wechat

import (
	. "apiserver/handler"
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/lexkong/log"
	"io"
	"sort"
	"strings"
	"time"
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

	log.Infof("WeChat ProcSignature Param : timestamp=%s,nonce=%s,signature=%s,echostr=%s,signatureGen=%s", timestamp, nonce, signature, echostr, signatureGen)

	if signature != signatureGen {
		SendResponse(c, nil, false)
	} else {

		SendResponseWithoutFormat(c, echostr)
	}

}

type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   time.Duration
	MsgType      CDATA
	Content      CDATA
}

type CDATA struct {
	Text string `xml:",cdata"`
}

func makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = CDATA{fromUserName}
	textResponseBody.ToUserName = CDATA{toUserName}
	textResponseBody.MsgType = CDATA{"text"}
	textResponseBody.Content = CDATA{content}
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(textResponseBody, " ", "  ")
}

func Post(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	log.Info(string(buf[0:n]))
	var event TextResponseBody
	msg := c.PostForm("msg")
	log.Info(msg)
	xml.Unmarshal([]byte(msg), &event)
	fmt.Println(event)
	ret, _ := makeTextResponseBody("JackRyannn", "renchao", "this is answer")
	SendResponseWithoutFormat(c, string(ret))

}
