package user

import (
	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/services/user"
	"github.com/pkg/errors"
	"net/http"
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
		res := new(modifyPersonalInfoResponse)
		if err := c.ShouldBindJSON(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		modifyData := new(user.ModifyData)
		if err := h.userService.Modify(ctx, uid, modifyData); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminModifyPersonalInfoError,
				code.Text(code.AdminModifyPersonalInfoError)).WithError(err),
			)
			return
		}

		c.Payload("success")
	}
}
