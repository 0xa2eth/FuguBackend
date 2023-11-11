package user

import (
	"errors"
	"net/http"

	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/services/user"
)

type modifyPersonalInfoRequest struct {
	Avatar   string `json:"avatar,omitempty"`
	NickName string `json:"nickName,omitempty"`
	Bio      string `json:"bio,omitempty"`
}

type modifyPersonalInfoResponse struct{}

// Modify 修改个人信息
// @Summary 修改个人信息
// @Description 修改个人信息
// @Tags API.admin
// @Accept application/json
// @Produce json
// @Param Request body modifyPersonalInfoRequest true "请求信息"
// @Success 200 {object} modifyPersonalInfoResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/modify_personal_info [patch]
func (h *handler) Modify() core.HandlerFunc {
	return func(c core.Context) {
		uid := c.Param("UserID")
		if uid == "" {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(errors.New("path param not found")),
			)
			return
		}
		req := new(modifyPersonalInfoRequest)
		//res := new(modifyPersonalInfoResponse)
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
		InnerID, err := h.hashids.HashidsDecode(uid)
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

		c.Payload(nil)
	}
}
