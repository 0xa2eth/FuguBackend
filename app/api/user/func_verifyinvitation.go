package user

import (
	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/pkg/errors"
	"net/http"
)

type verifyInviteCodeRequest struct {
	Code string `json:"invitationCode,omitempty"`
}

type verifyInviteCodeResponse struct{}

// VerifyInviteCode  验证邀请码有效性
// @Summary 验证邀请码有效性
// @Description 验证邀请码有效性
// @Tags API.user
// @Accept application/json
// @Produce json
// @Param Request body verifyInviteCodeRequest true "请求信息"
// @Success 200 {object} verifyInviteCodeResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/verifyinvitation [post]
func (h *handler) VerifyInviteCode() core.HandlerFunc {
	return func(c core.Context) {
		req := new(verifyInviteCodeRequest)

		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		exists := h.cache.Exists(req.Code)
		if !exists {
			c.AbortWithError(core.Error(
				http.StatusOK,
				code.CacheNotExist,
				code.Text(code.CacheNotExist)).WithError(errors.New("invitationCode not exist")),
			)
			return
		}

		c.Payload(exists)
	}
}
