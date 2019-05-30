package story

import (
	. "apiserver/handler"
	"strings"
)
import "apiserver/service"

import "github.com/gin-gonic/gin"

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	ret := ""
	key := c.Query("key")
	// Get the user by the `username` from the database.
	if strings.Contains(key, "getShareText") {
		ret = service.GetShareText()
	} else if strings.Contains(key, "setShareText") {
		ret = service.SetShareText(key)
	} else {
		ret = service.Receive("10001", key)
	}

	SendResponse(c, nil, ret)
}
