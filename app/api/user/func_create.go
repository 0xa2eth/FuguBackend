package user

import (
	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/validation"
	"FuguBackend/app/services/user"
	"FuguBackend/pkg/snowflake"
	"github.com/spf13/cast"
	"net/http"
)

type createRequest struct {
	Address       string `json:"address,omitempty" gorm:"column:address;type:varchar(255)"`
	TwitterID     string `json:"twitterID,omitempty" gorm:"column:twitter_id;type:varchar(255)"`
	TwitterAvatar string `json:"twitterAvatar,omitempty" gorm:"column:twitter_avatar;type:varchar(255)"`
	TwitterName   string `json:"twitterName,omitempty" gorm:"column:twitter_name;type:varchar(255)"`
	InviteCode    string `json:"inviteCode,omitempty"`
}

type createResponse struct {
	UserID string `json:"UserID"`
}

// Create 新增管理员
// @Summary 新增管理员
// @Description 新增管理员
// @Tags API.admin
// @Accept application/json
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [post]
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}

		genID, _ := snowflake.GenID()
		createData := &user.CreateUserData{
			Address:       "",
			UserID:        int64(genID),
			TwitterID:     req.TwitterID,
			TwitterName:   req.TwitterName,
			TwitterAvatar: req.TwitterAvatar,
		}

		id, err := h.userService.Create(c, createData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminCreateError,
				code.Text(code.AdminCreateError)).WithError(err),
			)
			return
		}

		hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(id)})
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsEncodeError,
				code.Text(code.HashIdsEncodeError)).WithError(err),
			)
			return
		}
		res.UserID = hashId
		c.Payload(res)
	}
}
