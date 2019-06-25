package tag

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"time"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	print("chaoren:" + r.Close_Time.String())
	u := model.TagModel{
		Name:        r.Name,
		Source:      r.Source,
		Category:    r.Category,
		Property:    r.Property,
		State:       r.State,
		Create_Time: time.Now(),
		Update_Time: time.Now(),
		Close_Time:  r.Close_Time,
		Operator:    r.Operator,
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	//// Encrypt the user password.
	//if err := u.Encrypt(); err != nil {
	//	SendResponse(c, errno.ErrEncrypt, nil)
	//	return
	//}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Name: r.Name,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
