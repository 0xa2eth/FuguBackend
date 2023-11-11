package secret

import (
	"FuguBackend/app/pkg/core"
)

type listRequest struct{}

type listResponse struct{}

// List 可见的秘密列表
// @Summary 可见的秘密列表
// @Description 可见的秘密列表
// @Tags API.secret
// @Accept application/json
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/secret/viewable [get]
func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
