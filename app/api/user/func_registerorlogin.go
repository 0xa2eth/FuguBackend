package user

import (
	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/validation"
	"FuguBackend/app/services/user"
	"FuguBackend/pkg/snowflake"
	"errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

type registerOrLoginRequest struct {
	//Address        string `json:"address,omitempty"`
	TwitterID      string `json:"twitterID,omitempty" binding:"required"`
	TwitterAvatar  string `json:"twitterAvatar,omitempty" binding:"required"`
	TwitterName    string `json:"twitterName,omitempty" binding:"required"`
	InvitationCode string `json:"invitationCode,omitempty"`
}

type registerOrLoginResponse struct {
	UserID string `json:"UserID"`
}

// RegisterOrLogin 注册或登陆
// @Summary 注册或登陆
// @Description 注册或登陆
// @Tags API.user
// @Accept application/json
// @Produce json
// @Param Request body registerOrLoginRequest true "请求信息"
// @Success 200 {object} registerOrLoginResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/twitterlogin [post]
func (h *handler) RegisterOrLogin() core.HandlerFunc {
	return func(c core.Context) {
		req := new(registerOrLoginRequest)
		res := new(registerOrLoginResponse)
		if err := c.ShouldBindJSON(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}
		//todo  如果注册过 走login 逻辑 颁发一个令牌 如果没注册过 走注册create逻辑  颁发随机头像
		search := &user.SearchOneData{TwitterID: req.TwitterID}
		users, err2 := h.userService.Detail(c, search)
		if err2 != nil && errors.Is(err2, gorm.ErrRecordNotFound) {
			// 有错误 ,没找到  -> 注册

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
				h.logger.Info("err:", zap.Error(err))
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
		if err2 != nil && !errors.Is(err2, gorm.ErrRecordNotFound) {
			// 有错误 internal error
			c.Payload(err2)
		}
		// 没错误 找到
		hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(users.Id)})
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
