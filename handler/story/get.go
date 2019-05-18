package story

import . "apiserver/handler"
import "apiserver/service"

import "github.com/gin-gonic/gin"

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	key := c.Query("key")
	// Get the user by the `username` from the database.
	ret := service.Receive(key)

	SendResponse(c, nil, ret)
}
