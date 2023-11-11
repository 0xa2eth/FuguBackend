package cave

import (
	"FuguBackend/app/pkg/core"
)

type secretListRequest struct{}

type secretListResponse struct{}

// SecretList 洞穴内秘密列表
// @Summary 洞穴内秘密列表
// @Description 洞穴内秘密列表
// @Tags API.cave
// @Accept application/json
// @Produce json
// @Param Request body secretListRequest true "请求信息"
// @Success 200 {object} secretListResponse
// @Failure 400 {object} code.Failure
// @Router /api/cave/:CaveID [get]
func (h *handler) SecretList() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
