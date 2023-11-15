package user

import (
	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/services/user"
	"errors"
	"net/http"
)

type modifyInfoRequest struct {
	Avatar   string `json:"avatar,omitempty"`
	NickName string `json:"nickName,omitempty"`
	Bio      string `json:"bio,omitempty"`
}

type modifyInfoResponse struct {
	Success bool `json:"success"`
}

// ModifyInfo 创建洞穴 修改个人（洞穴）信息
// @Summary 创建洞穴 修改个人（洞穴）信息
// @Description 创建洞穴 修改个人（洞穴）信息
// @Tags API.user
// @Accept application/json
// @Produce json
// @Param Request body modifyInfoRequest true "请求信息"
// @Success 200 {object} modifyInfoResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/:UserID [put]
func (h *handler) ModifyInfo() core.HandlerFunc {
	return func(c core.Context) {
		//uid := c.Param("UserID")
		//if uid == "" {
		//	c.AbortWithError(core.Error(
		//		http.StatusUnauthorized,
		//		code.AuthorizationError,
		//		code.Text(code.AuthorizationError)).WithError(errors.New("invalid token")),
		//	)
		//	return
		//}
		value, exists := c.Get("UserID")
		if !exists {
			c.AbortWithError(core.Error(
				http.StatusUnauthorized,
				code.AuthorizationError,
				code.Text(code.AuthorizationError)).WithError(errors.New("invalid token")),
			)
			return
		}
		hashID := value.(string)
		req := new(modifyInfoRequest)
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		modifyData := new(user.ModifyData)
		modifyData.NickName = req.NickName
		modifyData.Bio = req.Bio
		modifyData.Avatar = req.Avatar
		InnerID, err := h.hashids.HashidsDecode(hashID)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsEncodeError,
				code.Text(code.HashIdsEncodeError)).WithError(err),
			)
			return
		}
		if err := h.userService.Modify(c, int64(InnerID[0]), modifyData); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminModifyPersonalInfoError,
				code.Text(code.AdminModifyPersonalInfoError)).WithError(err),
			)
			return
		}
		m := new(modifyInfoResponse)
		m.Success = true
		c.Payload(m)
	}
}
