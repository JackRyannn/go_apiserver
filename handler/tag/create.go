package tag

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"strconv"
	"time"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r TagRequest
	if err := c.BindJSON(&r); err != nil {
		log.Info(errno.ErrBind.Message)
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	closeTime, _ := time.Parse("2006-01-02 15:04:05", "2019-06-26 00:00:00")
	log.Info("chaoren1:" + r.Name + r.Source + r.ClosedAt)
	//将string类型转换为uint64
	category, _ := strconv.ParseUint(r.Category, 12, 64)
	property, _ := strconv.ParseUint(r.Property, 12, 64)
	state, _ := strconv.ParseUint(r.State, 12, 64)
	u := model.TagModel{
		Name:     r.Name,
		Source:   r.Source,
		Category: category,
		Property: property,
		State:    state,
		ClosedAt: &closeTime,
		Operator: r.Operator,
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
		log.Info("chaoren: " + err.Error())
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Name: r.Name,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
