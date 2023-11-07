package secret

import (
	"FuguBackend/app/pkg/core"
)

type listRequest struct{}

type listResponse struct{}

// List 管理员列表
// @Summary 管理员列表
// @Description 管理员列表
// @Tags API.admin
// @Accept application/json
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [get]
func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
