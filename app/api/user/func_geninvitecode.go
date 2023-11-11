package user

import (
	"FuguBackend/app/pkg/core"
)

type genInviteCodeRequest struct{}

type genInviteCodeResponse struct{}

// GenInviteCode 生成邀请码
// @Summary 生成邀请码
// @Description 生成邀请码
// @Tags API.user
// @Accept application/json
// @Produce json
// @Param Request body genInviteCodeRequest true "请求信息"
// @Success 200 {object} genInviteCodeResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/invitecode [get]
func (h *handler) GenInviteCode() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
