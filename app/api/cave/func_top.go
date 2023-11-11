package cave

import (
	"FuguBackend/app/pkg/core"
)

type topRequest struct{}

type topResponse struct{}

// Top top5洞穴
// @Summary top5洞穴
// @Description top5洞穴
// @Tags API.cave
// @Accept application/json
// @Produce json
// @Param Request body topRequest true "请求信息"
// @Success 200 {object} topResponse
// @Failure 400 {object} code.Failure
// @Router /api/cave/top [get]
func (h *handler) Top() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
