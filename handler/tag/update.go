package tag

import (
	"strconv"

	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Update update a exist user account info.
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	// Get the user id from the url parameter.
	tagId, _ := strconv.Atoi(c.Param("id"))
	log.Infof("%s", tagId)

	// Binding the user data.
	var r TagRequest
	if err := c.BindJSON(&r); err != nil {
		log.Info(errno.ErrBind.Message)
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	state, _ := strconv.ParseUint(r.State, 12, 64)

	u := model.TagModel{
		BaseModel: model.BaseModel{
			Id: uint64(tagId),
		},
		State: state,
	}

	// We update the record based on the tag id.
	u.Id = uint64(tagId)
	log.Infof("%s", u.Id)
	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Save changed fields.
	if err := u.Update(&u); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
