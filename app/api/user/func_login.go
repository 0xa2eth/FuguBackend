package user

import (
	"FuguBackend/app/pkg/core"
)

type loginRequest struct{}

type loginResponse struct{}

// Login 管理员登录
// @Summary 管理员登录
// @Description 管理员登录
// @Tags API.admin
// @Accept application/json
// @Produce json
// @Param Request body loginRequest true "请求信息"
// @Success 200 {object} loginResponse
// @Failure 400 {object} code.Failure
// @Router /api/login [post]
func (h *handler) Login() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
